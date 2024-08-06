package main

import (
	"encoding/json"
	// "fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

type ToDo struct {
	ID       string `json:"id"`
	Task     string `json:"task"`
	Complete bool   `json:"done"`
}

var toDos []ToDo

func getToDos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(toDos)
}

func createToDo(w http.ResponseWriter, r *http.Request) {
	var toDo ToDo
	_ = json.NewDecoder(r.Body).Decode(&toDo)
	toDos = append(toDos, toDo)
	json.NewEncoder(w).Encode(toDo)
}

func upDateToDo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range toDos {
		if item.ID == params["id"] {
			toDos = append(toDos[:index], toDos[index+1:]...)
			var toDo ToDo
			_ = json.NewDecoder(r.Body).Decode(&toDo)
			toDo.ID = params["id"]
			toDos = append(toDos, toDo)
			json.NewEncoder(w).Encode(toDo)
			return
		}
	}
}

func deleteToDo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range toDos {
		if item.ID == params["id"] {
			toDos = append(toDos[:index], toDos[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(toDos)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/todos", getToDos).Methods("GET")
	router.HandleFunc("/todos", createToDo).Methods("POST")
	router.HandleFunc("/todos", upDateToDo).Methods("PUT")
	router.HandleFunc("/todos", deleteToDo).Methods("DELETE")
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8001", handler))
}
