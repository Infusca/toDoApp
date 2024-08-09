package main

import (
	"encoding/json"
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
	log.Println("getToDos called")
	json.NewEncoder(w).Encode(toDos)
}

func createToDo(w http.ResponseWriter, r *http.Request) {
	var toDo ToDo
	err := json.NewDecoder(r.Body).Decode(&toDo)
	if err != nil {
		log.Printf("Error decoding toDo: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("Adding toDo: %+v", toDo)
	toDos = append(toDos, toDo)
	json.NewEncoder(w).Encode(toDo)
}

func updateToDo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Printf("updateToDo called with ID: %s", params["id"])
	for index, item := range toDos {
		if item.ID == params["id"] {
			toDos = append(toDos[:index], toDos[index+1:]...)
			var toDo ToDo
			err := json.NewDecoder(r.Body).Decode(&toDo)
			if err != nil {
				log.Printf("Error decoding toDo: %v", err)
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			toDo.ID = params["id"]
			toDos = append(toDos, toDo)
			log.Printf("Updated toDo: %+v", toDo)
			json.NewEncoder(w).Encode(toDo)
			return
		}
	}
	log.Printf("toDo with ID %s not found", params["id"])
	http.Error(w, "toDo not found", http.StatusNotFound)
}

func deleteToDo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	log.Printf("deleteToDo called with ID: %s", params["id"])
	for index, item := range toDos {
		if item.ID == params["id"] {
			toDos = append(toDos[:index], toDos[index+1:]...)
			log.Printf("Deleted toDo with ID %s", params["id"])
			json.NewEncoder(w).Encode(toDos)
			return
		}
	}
	log.Printf("toDo with ID %s not found", params["id"])
	http.Error(w, "toDo not found", http.StatusNotFound)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/todos", getToDos).Methods("GET")
	router.HandleFunc("/todos", createToDo).Methods("POST")
	router.HandleFunc("/todos/{id}", updateToDo).Methods("PUT")
	router.HandleFunc("/todos/{id}", deleteToDo).Methods("DELETE")

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // Your React app's URL
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	}).Handler(router)

	log.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
