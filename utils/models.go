package utils

import (
	"time"

	"github.com/Flamme97/todoListGo/internal/database"
)


type ToDoList struct {
	List      string  `json:"list"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Complete 	bool 			`json:"complete"`
}

func DatabaseListToList(dbList database.Todolist) ToDoList {
	return ToDoList{
		List: 		 dbList.List,
		CreatedAt: dbList.CreatedAt,
		UpdatedAt: dbList.UpdatedAt,
		Complete: dbList.Complete,
	}
}


func DatabaselistsTolists(dbToDoLists []database.Todolist) []ToDoList {
	listOfToDo := []ToDoList{}
	for _, list := range dbToDoLists {
		listOfToDo = append(listOfToDo, DatabaseListToList(list))
	}
	return listOfToDo
}