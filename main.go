package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Flamme97/todoListGo/internal/database"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
);

var learning = "learn the go lang even better"
var runtime = "Go for a run"
var dinnertime= "make dinner"
var alltask = []string{learning, runtime, dinnertime}
var newtask = "go for a walk"


type ApiConfig struct {
	DB *database.Queries
}

type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}


func main() {
	router := chi.NewMux()

	router.Use(
		cors.Handler(cors.Options{}),
	)
	humaConfig := humachi.New(router, huma.DefaultConfig("ToDoList App", "1.0.0"))

	huma.Register(humaConfig, huma.Operation{
		OperationID: "get-greeting",
		Method:      http.MethodGet,
		Path:        "/greeting/{name}",
		Summary:     "Get a greeting",
		Description: "Get a greeting for a person by name.",
		Tags:        []string{"Greetings"},
	}, func(ctx context.Context, input *struct{
		Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
	}) (*GreetingOutput, error) {
		resp := &GreetingOutput{}
		resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
		return resp, nil
	})


	huma.Register(humaConfig, huma.Operation{
		OperationID: "get-list",
		Method:      http.MethodGet,
		Path:        "/v1/getTodoList",
		Summary:     "get Todo List",
		Description: "Getting all the todolist",
		Tags:        []string{"TodoList"},
	}, func(ctx context.Context, input *struct{
		Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
	}) (*GreetingOutput, error) {
		resp := &GreetingOutput{}
		resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
		return resp, nil
	})

	alltask = addTask(alltask, newtask)
	printTask(alltask)
		
	http.ListenAndServe("localhost:8080", router)
	

}

