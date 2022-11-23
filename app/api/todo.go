package api

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	Id        uuid.UUID `json:"id"`
	Title     TodoTitle `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

type TodoTitle string

func (s TodoTitle) Validate() {
}

type CreateTodoRequest struct {
	Title TodoTitle `json:"title"`
}

type UpdateTodoRequest struct {
	Title     TodoTitle `json:"title"`
	Completed bool      `json:"completed"`
}

type TodoServer interface {
	CreateTodo(context.Context, CreateTodoRequest) (*Todo, error)
	GetTodo(context.Context, uuid.UUID) (*Todo, error)
	ListTodo(context.Context) (*[]Todo, error)
	UpdateTodo(context.Context, UpdateTodoRequest) (*Todo, error)
	DeleteTodo(context.Context, uuid.UUID) error
	mustEmbedUnimplementedTodoServer()
}

type UnimplementedTodoServer struct {
}

func (UnimplementedTodoServer) CreateTodo(context.Context, CreateTodoRequest) (*Todo, error) {
	return nil, &HTTPError{Code: http.StatusNotImplemented, Message: "Create request is implemented"}
}

func (UnimplementedTodoServer) GetTodo(context.Context, uuid.UUID) (*Todo, error) {
	return nil, &HTTPError{Code: http.StatusNotImplemented, Message: "Get request is implemented"}
}

func (UnimplementedTodoServer) ListTodo(context.Context) (*[]Todo, error) {
	return nil, &HTTPError{Code: http.StatusNotImplemented}
}

func (UnimplementedTodoServer) UpdateTodo(context.Context, UpdateTodoRequest) (*Todo, error) {
	return nil, &HTTPError{Code: http.StatusNotImplemented}
}

func (UnimplementedTodoServer) DeleteTodo(context.Context, uuid.UUID) error {
	return &HTTPError{Code: http.StatusNotImplemented}
}

func (UnimplementedTodoServer) mustEmbedUnimplementedTodoServer() {}
