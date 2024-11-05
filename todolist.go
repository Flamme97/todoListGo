package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Flamme97/todoListGo/utils"
)

type TodoList struct {
	List      string `json:"list" example:"1. learn go! 2. Create Todo !" doc:"list"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


func (apiCfg *ApiConfig) HandlerGetTodoList(w http.ResponseWriter, r *http.Request) {
	list, err := apiCfg.DB.GetTodoList(r.Context())
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprint("couldn't get feeds", err))
		return
	}

	utils.RespondWithJSON(w, 201, utils.DatabaselistsTolists(list))

}

func printTask(taskItems []string){
	fmt.Println("List of my todos")
	for index, task := range taskItems {
		fmt.Printf("%d: %s \n", index+1, task)
	}
}

func addTask(taskItems []string, newTask string) []string{
	updatedlist :=append(taskItems, newTask)
	return updatedlist
}
