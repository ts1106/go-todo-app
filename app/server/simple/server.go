package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"github.com/ts1106/go-todo-app/app/api"
	todo "github.com/ts1106/go-todo-app/app/internal/context"
	ts "github.com/ts1106/go-todo-app/app/internal/server"
)

func main() {
	r := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}

type Router struct {
	srv *ts.TodoServer
}

func NewRouter() *Router {
	return &Router{srv: ts.NewTodoServer()}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := todo.Context{Writer: w, Request: req}
	path := req.URL.RawPath
	if path == "" {
		path = req.URL.Path
	}
	id := strings.TrimPrefix(path, "/todos/")
	if id != "" {
		id, err := uuid.Parse(id)
		if err != nil {
			err := &api.HTTPError{Code: http.StatusBadRequest, Message: err.Error()}
			ctx.JSON(err.Code, api.Response{Error: &api.Error{Code: err.Code, Message: err.Error()}})
			return
		}
		switch req.Method {
		case http.MethodGet:
			ts.GetTodoHandler(ctx, r.srv, id)
		case http.MethodPatch:
			ts.UpdateTodoHandler(ctx, r.srv, id)
		case http.MethodDelete:
			ts.DeleteTodoHandler(ctx, r.srv, id)
		default:
			err := &api.HTTPError{Code: http.StatusMethodNotAllowed}
			ctx.JSON(err.Code, api.Response{Error: &api.Error{Code: err.Code, Message: err.Error()}})
		}
	} else {
		switch req.Method {
		case http.MethodGet:
			ts.ListTodoHandler(ctx, r.srv)
		case http.MethodPost:
			ts.CreateTodoHandler(ctx, r.srv)
		default:
			err := &api.HTTPError{Code: http.StatusMethodNotAllowed}
			ctx.JSON(err.Code, api.Response{Error: &api.Error{Code: err.Code, Message: err.Error()}})
		}
	}
}
