package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Flamme97/todoListGo/handlers"
	"github.com/Flamme97/todoListGo/internal/database"
	"github.com/google/uuid"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	_ "github.com/danielgtaylor/huma/v2/formats/cbor"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// var learning = "learn the go lang even better"
// var runtime = "Go for a run"
// var dinnertime= "make dinner"
// var alltask = []string{learning, runtime, dinnertime}
// var newtask = "go for a walk"

// type GreetingOutput struct {
// 	Body struct {
// 		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
// 	}
// }

func main() {
	router := chi.NewMux()
	godotenv.Load(".env")

	dbURL := os.Getenv("DB_URL")
	if dbURL == ""{
		log.Fatal("DB_URL is not found in the environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil{
		log.Fatal("cannot connect to server", err)
	}
	db := database.New(conn)

	apiCfg := handlers.ApiConfig{
		DB: db,
	}


	router.Use(
		cors.Handler(cors.Options{}),
	)
	humaConfig := humachi.New(router, huma.DefaultConfig("ToDoList App", "1.0.0"))

	// huma.Register(humaConfig, huma.Operation{
	// 	OperationID: "get-greeting",
	// 	Method:      http.MethodGet,
	// 	Path:        "/greeting/{name}",
	// 	Summary:     "Get a greeting",
	// 	Description: "Get a greeting for a person by name.",
	// 	Tags:        []string{"Greetings"},
	// }, func(ctx context.Context, input *struct{
	// 	Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
	// }) (*GreetingOutput, error) {
	// 	resp := &GreetingOutput{}
	// 	resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
	// 	return resp, nil
	// })

	huma.Register(humaConfig, huma.Operation{
		OperationID: "get-list",
		Method:      http.MethodGet,
		Path:        "/v1/todolist",
		Summary:     "Getting all list",
		Description: "Get the total overviews of todos",
		Tags:        []string{"ToDoList"},
		Errors: 		 	 []int{400},
		DefaultStatus: 201,
	}, func(ctx context.Context, input *struct{}) (*handlers.ToDoListOutput, error) {
		
		return apiCfg.HandlerGetTodoList()
	})

	huma.Register(humaConfig, huma.Operation{
    OperationID: 	 "create-todo",
    Method:      	 http.MethodPost,
    Path:        	 "/v1/createtodolist",
    Summary:     	 "Create a Todo Task",
    Description: 	 "Creating a new ToDo Task",
    Tags:        	 []string{"CreateToDoInput"},
		Errors: 		 	 []int{400},
		DefaultStatus: 201,
	}, func(ctx context.Context, input *handlers.CreateToDoInput) (*handlers.CreateToDoResponse, error) {
    return apiCfg.HandlerCreateToDo(ctx, input)
})

huma.Register(humaConfig, huma.Operation{
	OperationID: 	 "complete-todo",
	Method:      	 http.MethodPut,
	Path:        	 "/v1/todolist/{id}",
	Summary:     	 "complete Todo List",
	Description: 	 "Mark it as complete in the DB",
	Tags:        	 []string{"CompleteToDo"},
	Errors: 		 	 []int{400},
	DefaultStatus: 201,
}, func(ctx context.Context, input *struct {
	ID string `path:"id" doc:"The ID of the todo item to mark as complete"`
}) (*struct{}, error) { 
	todoID, err := uuid.Parse(input.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID format: %w", err)
	}
	return apiCfg.HandlerCompleteTodo(ctx, todoID)
})

http.ListenAndServe("localhost:8080", router)


}
		





