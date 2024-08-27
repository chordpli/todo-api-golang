package handlers

import (
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"todo-api-golang/ent"
	"todo-api-golang/internal/dto"
	"todo-api-golang/internal/service"
	response "todo-api-golang/middleware"
)

type TodoHandlerInterface interface {
	CreateTodo(w http.ResponseWriter, r *http.Request)
	ListTodos(w http.ResponseWriter, r *http.Request)
	GetTodo(w http.ResponseWriter, r *http.Request)
	UpdateTodo(w http.ResponseWriter, r *http.Request)
	DeleteTodo(w http.ResponseWriter, r *http.Request)
	UpdateTodoStatus(w http.ResponseWriter, r *http.Request)
}

type TodoHandler struct {
	service service.TodoService
}

// NewTodoHandler creates a new TodoHandler.
func NewTodoHandler(service service.TodoService) TodoHandlerInterface {
	return &TodoHandler{service: service}
}

// CreateTodo godoc
// @Summary Create a new Todo
// @Description Create a new Todo with the given details
// @Tags todos
// @Accept  json
// @Produce  json
// @Param todo body dto.TodoForm true "Todo form"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/todos [post]
func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var form dto.TodoForm
	httpCode, errCode := response.BindAndValid(r, &form)
	if httpCode != http.StatusOK {
		response.ResponseJSON(w, httpCode, errCode, "Invalid form data", nil)
		return
	}

	todoDTO, err := h.service.CreateTodo(context.Background(), form)
	if err != nil {
		response.ResponseJSON(w, http.StatusInternalServerError, 500, "Failed to create todo", nil)
		return
	}

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
func (h *TodoHandler) ListTodos(w http.ResponseWriter, r *http.Request) {
	todoDTOs, err := h.service.ListTodos(context.Background())
	if err != nil {
		response.ResponseJSON(w, http.StatusInternalServerError, 500, "Failed to list todos", nil)
		return
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
func (h *TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.ResponseJSON(w, http.StatusBadRequest, 400, "Invalid todo ID", nil)
		return
	}

	todoDTO, err := h.service.GetTodoByID(context.Background(), id)
	if err != nil {
		if ent.IsNotFound(err) {
			response.ResponseJSON(w, http.StatusNotFound, 404, "Todo not found", nil)
		} else {
			response.ResponseJSON(w, http.StatusInternalServerError, 500, "Failed to fetch todo", nil)
		}
		return
	}

	response.ResponseJSON(w, http.StatusOK, 200, "Todo fetched successfully", todoDTO)
}

// UpdateTodo godoc
// @Summary Update an existing Todo
// @Description Update the details of a Todo by its ID
// @Tags todos
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Param todo body dto.TodoForm true "Todo form"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/todos/{id} [put]
func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.ResponseJSON(w, http.StatusBadRequest, 400, "Invalid todo ID", nil)
		return
	}

	var form dto.TodoForm
	httpCode, errCode := response.BindAndValid(r, &form)
	if httpCode != http.StatusOK {
		response.ResponseJSON(w, httpCode, errCode, "Invalid form data", nil)
		return
	}

	todoDTO, err := h.service.UpdateTodo(context.Background(), id, form)
	if err != nil {
		if ent.IsNotFound(err) {
			response.ResponseJSON(w, http.StatusNotFound, 404, "Todo not found", nil)
		} else {
			response.ResponseJSON(w, http.StatusInternalServerError, 500, "Failed to update todo", nil)
		}
		return
	}

	response.ResponseJSON(w, http.StatusOK, 200, "Todo updated successfully", todoDTO)
}

// DeleteTodo godoc
// @Summary Soft delete a Todo by ID
// @Description Soft delete a Todo by setting the deleted_at field
// @Tags todos
// @Param id path int true "Todo ID"
// @Success 204 {object} nil "No content"
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/todos/{id} [delete]
func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.ResponseJSON(w, http.StatusBadRequest, 400, "Invalid todo ID", nil)
		return
	}

	err = h.service.DeleteTodo(context.Background(), id)
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
// @Param status body dto.UpdateStatusForm true "New status"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/todos/{id}/status [put]
func (h *TodoHandler) UpdateTodoStatus(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.ResponseJSON(w, http.StatusBadRequest, 400, "Invalid todo ID", nil)
		return
	}

	var form dto.UpdateStatusForm
	httpCode, errCode := response.BindAndValid(r, &form)
	if httpCode != http.StatusOK {
		response.ResponseJSON(w, httpCode, errCode, "Invalid form data", nil)
		return
	}

	todoDTO, err := h.service.UpdateTodoStatus(context.Background(), id, form.Status)
	if err != nil {
		if ent.IsNotFound(err) {
			response.ResponseJSON(w, http.StatusNotFound, 404, "Todo not found", nil)
		} else {
			response.ResponseJSON(w, http.StatusInternalServerError, 500, "Failed to update status", nil)
		}
		return
	}

	response.ResponseJSON(w, http.StatusOK, 200, "Status updated successfully", todoDTO)
}
