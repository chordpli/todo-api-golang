package handlers

import (
	"context"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
	"time"
	"todo-api-golang/edge/database"
	"todo-api-golang/ent"
	"todo-api-golang/ent/todo"
	"todo-api-golang/internal/dto"
	response "todo-api-golang/middleware"
)

var client *ent.Client

func init() {
	client = database.InitDB()
	if client == nil {
		log.Fatal("Failed to initialize database client")
	}
}

type TodoForm struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Status      string `json:"status" validate:"required,oneof=PENDING COMPLETED PROGRESS"`
}

type UpdateStatusForm struct {
	Status string `json:"status" validate:"required,oneof=PENDING COMPLETED PROGRESS"`
}

// CreateTodo godoc
// @Summary Create a new Todo
// @Description Create a new Todo with the given details
// @Tags todos
// @Accept  json
// @Produce  json
// @Param todo body TodoForm true "Todo form"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/todos [post]
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var form TodoForm

	httpCode, errCode := response.BindAndValid(r, &form)

	if httpCode != http.StatusOK {
		response.ResponseJSON(w, httpCode, errCode, "Invalid form data", nil)
		return
	}

	status := todo.Status(form.Status)

	todoItem, err := client.Todo.
		Create().
		SetTitle(form.Title).
		SetDescription(form.Description).
		SetStatus(status).
		Save(context.Background())

	if err != nil {
		log.Printf("Failed to create todo: %v", err)
		response.ResponseJSON(w, http.StatusInternalServerError, 500, "Failed to create todo", nil)
		return
	}

	todoDTO := dto.ConvertTodoToDTO(todoItem)
	response.ResponseJSON(w, http.StatusOK, 200, "Todo created successfully", todoDTO)
}

// ListTodos godoc
// @Summary List all Todos
// @Description Get a list of all Todos
// @Tags todos
// @Produce  json
// @Success 200 {array} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/todos [get]
func ListTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := client.Todo.Query().
		Where(todo.DeletedAtIsNil()). // deleted_at이 NULL인 항목만 조회
		All(context.Background())

	if err != nil {
		response.ResponseJSON(w, http.StatusInternalServerError, 500, "Failed to list todos", nil)
		return
	}

	todoDTOs := make([]dto.TodoDTO, len(todos))
	for i, t := range todos {
		todoDTOs[i] = dto.ConvertTodoToDTO(t)
	}

	response.ResponseJSON(w, http.StatusOK, 200, "Todos fetched successfully", todoDTOs)
}

// GetTodo godoc
// @Summary Get a Todo by ID
// @Description Get details of a Todo by its ID
// @Tags todos
// @Produce  json
// @Param id path int true "Todo ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/todos/{id} [get]
func GetTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.ResponseJSON(w, http.StatusBadRequest, 400, "Invalid todo ID", nil)
		return
	}

	todoItem, err := client.Todo.Get(context.Background(), id)
	if err != nil {
		if ent.IsNotFound(err) {
			response.ResponseJSON(w, http.StatusNotFound, 404, "Todo not found", nil)
		} else {
			response.ResponseJSON(w, http.StatusInternalServerError, 500, "Failed to fetch todo", nil)
		}
		return
	}

	todoDTO := dto.ConvertTodoToDTO(todoItem)
	response.ResponseJSON(w, http.StatusOK, 200, "Todo fetched successfully", todoDTO)
}

// UpdateTodo godoc
// @Summary Update an existing Todo
// @Description Update the details of a Todo by its ID
// @Tags todos
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Param todo body TodoForm true "Todo form"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/todos/{id} [put]
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.ResponseJSON(w, http.StatusBadRequest, 400, "Invalid todo ID", nil)
		return
	}

	var form TodoForm
	httpCode, errCode := response.BindAndValid(r, &form)
	if httpCode != http.StatusOK {
		response.ResponseJSON(w, httpCode, errCode, "Invalid form data", nil)
		return
	}

	todoItem, err := client.Todo.Get(context.Background(), id)
	if err != nil {
		if ent.IsNotFound(err) {
			response.ResponseJSON(w, http.StatusNotFound, 404, "Todo not found", nil)
		} else {
			response.ResponseJSON(w, http.StatusInternalServerError, 500, "Failed to fetch todo", nil)
		}
		return
	}

	status := todo.Status(form.Status)

	_, err = todoItem.Update().
		SetTitle(form.Title).
		SetDescription(form.Description).
		SetStatus(status).
		Save(context.Background())

	if err != nil {
		response.ResponseJSON(w, http.StatusInternalServerError, 500, "Failed to update todo", nil)
		return
	}

	todoDTO := dto.ConvertTodoToDTO(todoItem)
	response.ResponseJSON(w, http.StatusOK, 200, "Todo updated successfully", todoDTO)
}

// DeleteTodo godoc
// @Summary Soft delete a Todo by ID
// @Description Soft delete a Todo by setting the deleted_at field
// @Tags todos
// @Param id path int true "Todo ID"
// @Success 204 {object} nil
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/todos/{id} [delete]
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		response.ResponseJSON(w, http.StatusBadRequest, 400, "Invalid todo ID", nil)
		return
	}

	// 소프트 삭제를 위해 deleted_at 필드에 현재 시간을 설정
	_, err = client.Todo.UpdateOneID(id).
		SetDeletedAt(time.Now()).
		Save(context.Background())

	if err != nil {
		if ent.IsNotFound(err) {
			response.ResponseJSON(w, http.StatusNotFound, 404, "Todo not found", nil)
		} else {
			response.ResponseJSON(w, http.StatusInternalServerError, 500, "Failed to delete todo", nil)
		}
		return
	}

	response.ResponseJSON(w, http.StatusNoContent, 204, "Todo deleted successfully", nil)
}

// UpdateTodoStatus godoc
// @Summary Update the status of a Todo
// @Description Update only the status field of a Todo by its ID
// @Tags todos
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Param status body UpdateStatusForm true "New status"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/todos/{id}/status [put]
func UpdateTodoStatus(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.ResponseJSON(w, http.StatusBadRequest, 400, "Invalid todo ID", nil)
		return
	}

	var form UpdateStatusForm
	httpCode, errCode := response.BindAndValid(r, &form)
	if httpCode != http.StatusOK {
		response.ResponseJSON(w, httpCode, errCode, "Invalid form data", nil)
		return
	}

	todoItem, err := client.Todo.Get(context.Background(), id)
	if err != nil {
		if ent.IsNotFound(err) {
			response.ResponseJSON(w, http.StatusNotFound, 404, "Todo not found", nil)
		} else {
			response.ResponseJSON(w, http.StatusInternalServerError, 500, "Failed to fetch todo", nil)
		}
		return
	}

	// Update the status field
	_, err = todoItem.Update().
		SetStatus(todo.Status(form.Status)).
		Save(context.Background())

	if err != nil {
		response.ResponseJSON(w, http.StatusInternalServerError, 500, "Failed to update status", nil)
		return
	}

	response.ResponseJSON(w, http.StatusOK, 200, "Status updated successfully", nil)
}
