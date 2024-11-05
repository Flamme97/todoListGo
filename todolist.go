package main

import (
	"fmt"
	"time"
)

type TodoList struct {
	List      string `json:"list" example:"1. learn go! 2. Create Todo !" doc:"list"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


func HandlerGetTodoList() *TodoList{
	fmt.Println("hello world!")
	example := TodoList{
		List: "Hello world",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	return example
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

// func showTask(w []string, task []string) {
// 	for _, task := range alltask {
// 		fmt.Fprintln(w, task)
// 	}
// }