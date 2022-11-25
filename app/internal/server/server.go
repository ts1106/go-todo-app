package server

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/ts1106/go-todo-app/app/api"
)

const IdKey = "IdKey"

type TodoServer struct {
	api.UnimplementedTodoServer
	todos []*api.Todo
}

func NewTodoServer() *TodoServer {
	s := &TodoServer{todos: []*api.Todo{}}
	return s
}

func (s *TodoServer) CreateTodo(ctx context.Context, param api.CreateTodoRequest) (*api.Todo, error) {
	data := &api.Todo{Id: uuid.New(), Title: param.Title, Completed: false, CreatedAt: time.Now()}
	s.todos = append(s.todos, data)
	return data, nil
}

func (s *TodoServer) GetTodo(ctx context.Context, id uuid.UUID) (*api.Todo, error) {
	for _, v := range s.todos {
		if v.Id == id {
			return v, nil
		}
	}
	return nil, nil
}

func (s *TodoServer) DeleteTodo(ctx context.Context, id uuid.UUID) error {
	for i, v := range s.todos {
		if v.Id == id {
			s.todos = append(s.todos[:i], s.todos[i+1:]...)
			return nil
		}
	}
	return nil
}

func (s *TodoServer) UpdateTodo(ctx context.Context, param api.UpdateTodoRequest) (*api.Todo, error) {
	for _, v := range s.todos {
		if v.Id == param.Id {
			if param.Title != nil {
				v.Title = *param.Title
			}
			if param.Completed != nil {
				v.Completed = *param.Completed
			}
			return v, nil
		}
	}
	return nil, nil
}

func (s *TodoServer) ListTodo(ctx context.Context) ([]*api.Todo, error) {
	return s.todos, nil
}
