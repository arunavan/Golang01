package main

import (
	"encoding/json"
	"fmt"
	"log"
     	"net/http"
	"github.com/gorilla/mux"
)
type Student struct {
	Name      string
	age int
	marks int
}
type Students []Student

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Home)
	router.HandleFunc("/showall", ShowStudents)
	router.HandleFunc("/showstudent/{studentname}",ShowStudent)

	log.Fatal(http.ListenAndServe(":8090", router))
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to student app!\n")
}

func ShowStudents(w http.ResponseWriter, r *http.Request) {
	students := Students{
		Student{Name: "Srikar", age :20 ,marks:90},
		Student{Name: "Sai",age:21, marks:89},
	}

	json.NewEncoder(w).Encode(students)
}

func ShowStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sname := vars["studentname"]
	fmt.Fprintf(w, "Student details: %s\n", sname)
}
