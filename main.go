package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"encoding/json"
)

type task struct {
	ID        int    `json:"ID"`
	Name      string `json:"Name"`
	Content   string `json:"Content"`
	Completed bool   `json:"Completed"`
}

type allTasks []task

var tasks = allTasks{
	{
		ID:        1,
		Name:      "Task One",
		Content:   "Some Content",
		Completed: false,
	},
	{
		ID:        2,
		Name:      "Task Two",
		Content:   "Some Content",
		Completed: true,
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the task name and content only in order to update")
	}

	json.Unmarshal(reqBody, &newTask)

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)
	w.WriteHeader(http.StatusCreated)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newTask)
}

func main() {
	// Init Router and configure it
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/create", createTask).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", router))

}
