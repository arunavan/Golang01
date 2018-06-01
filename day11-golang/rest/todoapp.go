package main

import (
	"encoding/json"
	"fmt"
	"log"
     "time"
	"net/http"

	"github.com/gorilla/mux"
)



type Todo struct {
	Name      string
	Completed bool
	Due       time.Time
}

type Todos []Todo

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	log.Fatal(http.ListenAndServe(":8088", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation", Completed: true},
		Todo{Name: "Host meetup",Completed:true},
	}

	json.NewEncoder(w).Encode(todos)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintf(w, "Todo show: %s\n", todoId)
}
