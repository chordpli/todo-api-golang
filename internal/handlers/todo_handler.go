package handlers

import (
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"todo-api-golang/edge/database"
	"todo-api-golang/ent"
	"todo-api-golang/ent/todo"
	response "todo-api-golang/middleware"
)

var client *ent.Client = database.InitDB()

type TodoForm struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Status      string `json:"status" validate:"required,oneof='PENDING COMPLETED PROGRESS'"`
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var form TodoForm

	// Bind and validate the form
	httpCode, errCode := response.BindAndValid(r, &form)
	if httpCode != http.StatusOK {
		response.ResponseJSON(w, httpCode, errCode, "Invalid form data", nil)
		return
	}

	// Create the todo using the validated data
	status := todo.Status(form.Status)

	todos, err := client.Todo.
		Create().
		SetTitle(form.Title).
		SetDescription(form.Description).
		SetStatus(status).
		Save(context.Background())
	if err != nil {
		response.ResponseJSON(w, http.StatusInternalServerError, 500, "Failed to create todo", nil)
		return
	}

	response.ResponseJSON(w, http.StatusOK, 200, "Todo created successfully", todos)
}

func ListTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := client.Todo.Query().All(context.Background())
	if err != nil {
		response.ResponseJSON(w, http.StatusInternalServerError, 500, "Failed to list todos", nil)
		return
	}

	response.ResponseJSON(w, http.StatusOK, 200, "Todos fetched successfully", todos)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.ResponseJSON(w, http.StatusBadRequest, 400, "Invalid todo ID", nil)
		return
	}

	todo, err := client.Todo.Get(context.Background(), id)
	if err != nil {
		if ent.IsNotFound(err) {
			response.ResponseJSON(w, http.StatusNotFound, 404, "Todo not found", nil)
		} else {
			response.ResponseJSON(w, http.StatusInternalServerError, 500, "Failed to fetch todo", nil)
		}
		return
	}

	response.ResponseJSON(w, http.StatusOK, 200, "Todo fetched successfully", todo)
}

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

	todos, err := client.Todo.Get(context.Background(), id)
	if err != nil {
		if ent.IsNotFound(err) {
			response.ResponseJSON(w, http.StatusNotFound, 404, "Todo not found", nil)
		} else {
			response.ResponseJSON(w, http.StatusInternalServerError, 500, "Failed to fetch todo", nil)
		}
		return
	}

	status := todo.Status(form.Status)

	_, err = todos.Update().
		SetTitle(form.Title).
		SetDescription(form.Description).
		SetStatus(status).
		Save(context.Background())
	if err != nil {
		response.ResponseJSON(w, http.StatusInternalServerError, 500, "Failed to update todo", nil)
		return
	}

	response.ResponseJSON(w, http.StatusOK, 200, "Todo updated successfully", todos)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		response.ResponseJSON(w, http.StatusBadRequest, 400, "Invalid todo ID", nil)
		return
	}

	err = client.Todo.DeleteOneID(id).Exec(context.Background())
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
