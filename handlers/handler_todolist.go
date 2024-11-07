package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/Flamme97/todoListGo/internal/database"
	"github.com/google/uuid"
)
type ApiConfig struct {
	DB *database.Queries
}
type CreateTodoParams struct {
	List      string
	Complete  bool
}
type ToDoListOutput struct {
	Body []ToDoList `json:"body"`
}

type ToDoList struct {
	ID   uuid.UUID  `json:"id" `
	List string 		`json:"list" example:"Buy groceries"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Complete 	bool 			`json:"complete"`
}
type CreateToDoInput struct {
		Body  CreateTodoParams
}
type CreateToDoResponse struct {
	ID 				uuid.UUID 				`json:"id"`
	List 			string 	    			`json:"list"`
	CreatedAt time.Time 			 	`json:"created_at"`
	UpdatedAt time.Time	  			`json:"updated_at"`
	Complete bool  							`json:"complete"`
}

type CompleteToDo struct {
	ID 			 uuid.UUID 		`json:"id"`
	Complete bool 				`json:"complete"`
}


func (apiCfg *ApiConfig) HandlerGetTodoList() (*ToDoListOutput, error) {
	list, err := apiCfg.DB.GetTodoList(context.Background())
	if err != nil {
		return nil, fmt.Errorf("couldn't get feeds: %w", err)
	}

	var ToDoLists []ToDoList
	for _, dbItem := range list {
		item := ToDoList{
			ID:   dbItem.ID,        
			List: dbItem.List,      
		}
		ToDoLists = append(ToDoLists, item)
	}

	return &ToDoListOutput{Body: ToDoLists}, nil
}


func (apiCfg *ApiConfig) HandlerCreateToDo(ctx context.Context, input *CreateToDoInput) (*CreateToDoResponse, error) {
	todoParams := database.CreateTodoParams{
		ID: uuid.New(),
		List: input.Body.List,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Complete: input.Body.Complete,
	}
	todo, err := apiCfg.DB.CreateTodo(ctx, todoParams)
	if err != nil {
		fmt.Println("failed to create todo, error:", err)
		return nil, err
	}
	return &CreateToDoResponse{
		List: todo.List,
		Complete: false,
	}, nil
}

func (apiCfg *ApiConfig) HandlerCompleteTodo(ctx context.Context, input uuid.UUID) (*struct{}, error) {
	todo, err := apiCfg.DB.CompleteToDo(ctx, input)
	if err != nil {
		fmt.Println("failed to create todo, error:", err)
		return nil, err
	}

	fmt.Printf("task has been marked as complete : %+v\n", todo)
	return nil, nil
}

