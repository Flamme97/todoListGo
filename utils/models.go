package utils

import (
	"time"
	"github.com/google/uuid"
	"github.com/Flamme97/todoListGo/internal/database"
)


type ToDoList struct {

	ID        uuid.UUID `json:"id"`
	List      string  	`json:"list"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Complete 	bool 			`json:"complete"`
}

func DatabaseListToList(dbList database.Todolist) ToDoList {
	return ToDoList{
		ID: 				dbList.ID,
		List: 		 dbList.List,
		CreatedAt: dbList.CreatedAt,
		UpdatedAt: dbList.UpdatedAt,
		Complete: dbList.Complete,
	}
}


func DatabaselistsTolists(dbToDoLists []database.Todolist) []ToDoList {
	listOfToDos := []ToDoList{}
	for _, list := range dbToDoLists {
		listOfToDos = append(listOfToDos, DatabaseListToList(list))
	}
	return listOfToDos
}