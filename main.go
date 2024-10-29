package main

import (
	"fmt"
	"net/http"
);

var learning = "learn the go lang even better"
var runtime = "Go for a run"
var dinnertime= "make dinner"
var alltask = []string{learning, runtime, dinnertime}
var newtask = "go for a walk"

func main() {



		fmt.Println()
		alltask = addTask(alltask, newtask)
		printTask(alltask)


		numb := 10
		numb2 := 15
		result := addtwonumberS(numb, numb2)
		fmt.Println(result)
		http.HandleFunc("/", helloUser)
		http.HandleFunc("/show-task", showTask)
		http.ListenAndServe(":8080", nil)

	

}

func helloUser(w http.ResponseWriter, request *http.Request) {
	greeting := "hello fellow developer, welcome to golang"
	fmt.Fprintln(w, greeting)
}

func showTask(w http.ResponseWriter, request *http.Request) {
	for _, task := range alltask {
		fmt.Fprintln(w, task)
	}
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

	func addtwonumberS(number int, numb2 int ) int{
		return number + numb2 
	}
