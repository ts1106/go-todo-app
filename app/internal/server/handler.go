package server

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/ts1106/go-todo-app/app/api"
	"github.com/ts1106/go-todo-app/app/internal/context"
)

func CreateTodoHandler(c context.Context, srv api.TodoServer) {
	var code int
	var param api.CreateTodoRequest
	err := json.NewDecoder(c.Request.Body).Decode(&param)
	if err != nil {
		code = http.StatusBadRequest
		c.JSON(code, api.Error{Code: code})
	}
	d, err := srv.CreateTodo(c.Request.Context(), param)
	if err != nil {
		err, ok := err.(*api.HTTPError)
		var code int
		if ok {
			code = err.Code
		} else {
			code = http.StatusInternalServerError
			err = &api.HTTPError{Code: code}
		}
		res := api.Response{Error: &api.Error{Code: code, Message: err.Error()}}
		c.JSON(code, res)
		return
	}
	c.JSON(http.StatusOK, api.Response{Data: &api.Data{Items: []api.Todo{*d}}})
}

func GetTodoHandler(c context.Context, srv api.TodoServer, id uuid.UUID) {
	d, err := srv.GetTodo(c.Request.Context(), id)
	if err != nil {
		err, ok := err.(*api.HTTPError)
		var code int
		if ok {
			code = err.Code
		} else {
			code = http.StatusInternalServerError
			err = &api.HTTPError{Code: code}
		}
		res := api.Response{Error: &api.Error{Code: code, Message: err.Error()}}
		c.JSON(code, res)
		return
	}
	c.JSON(http.StatusOK, api.Response{Data: &api.Data{Items: []api.Todo{*d}}})
}

func ListTodoHandler(c context.Context, srv api.TodoServer) {
	d, err := srv.ListTodo(c.Request.Context())
	if err != nil {
		err, ok := err.(*api.HTTPError)
		var code int
		if ok {
			code = err.Code
		} else {
			code = http.StatusInternalServerError
			err = &api.HTTPError{Code: code}
		}
		res := api.Response{Error: &api.Error{Code: code, Message: err.Error()}}
		c.JSON(code, res)
		return
	}
	c.JSON(http.StatusOK, api.Response{Data: &api.Data{Items: d}})
}

func UpdateTodoHandler(c context.Context, srv api.TodoServer, id uuid.UUID) {
	var code int
	var param api.UpdateTodoRequest
	err := json.NewDecoder(c.Request.Body).Decode(&param)
	if err != nil {
		code = http.StatusBadRequest
		c.JSON(code, api.Response{Error: &api.Error{Code: code, Message: err.Error()}})
		return
	}
	param.Id = id
	d, err := srv.UpdateTodo(c.Request.Context(), param)
	if err != nil {
		err, ok := err.(*api.HTTPError)
		if ok {
			code = err.Code
		} else {
			code = http.StatusInternalServerError
			err = &api.HTTPError{Code: code}
		}
		res := api.Response{Error: &api.Error{Code: code, Message: err.Error()}}
		c.JSON(code, res)
		return
	}
	c.JSON(http.StatusOK, api.Response{Data: &api.Data{Items: []api.Todo{*d}}})
}

func DeleteTodoHandler(c context.Context, srv api.TodoServer, id uuid.UUID) {
	err := srv.DeleteTodo(c.Request.Context(), id)
	if err != nil {
		err, ok := err.(*api.HTTPError)
		var code int
		if ok {
			code = err.Code
		} else {
			code = http.StatusInternalServerError
			err = &api.HTTPError{Code: code}
		}
		res := api.Response{Error: &api.Error{Code: code, Message: err.Error()}}
		c.JSON(code, res)
		return
	}
	c.JSON(http.StatusOK, api.Response{})
}
