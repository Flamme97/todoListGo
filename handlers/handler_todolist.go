package handlers

import (
	"context"
	"fmt"

	"github.com/Flamme97/todoListGo/internal/database"
	"github.com/google/uuid"
)
type ApiConfig struct {
	DB *database.Queries
}

type ToDoListOutput struct {
	Body []ToDoItem `json:"body"`
}

type ToDoItem struct {
	ID   uuid.UUID  `json:"id" `
	Task string 		`json:"task" example:"Buy groceries"`
}





func (apiCfg *ApiConfig) HandlerGetTodoList(ctx context.Context, input *struct{}) (*ToDoListOutput, error) {
	list, err := apiCfg.DB.GetTodoList(ctx)
	if err != nil {
		return nil, fmt.Errorf("couldn't get feeds: %w", err)
	}

	var ToDoLists []ToDoItem
	for _, dbItem := range list {
		item := ToDoItem{
			ID:   dbItem.ID,        
			Task: dbItem.List,      
		}
		ToDoLists = append(ToDoLists, item)
	}

	return &ToDoListOutput{Body: ToDoLists}, nil
}







// type parameters struct {
	// List string `json:"list"`
	// Complete bool `json:"complete"`
	
// }
