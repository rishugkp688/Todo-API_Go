package main

import (
	"fmt"
	"net/http"
)

var todos = make([]string, 0, 5)

func addTodo(w http.ResponseWriter, r *http.Request) {
	todos = append(todos, r.FormValue("addpoint"))
}

func showTodo(w http.ResponseWriter, r *http.Request) {
	for i, todo := range todos {
		fmt.Fprintf(w, "%d: %s", i+1, todo)
	}
}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("POST /todoadd", addTodo)
	http.HandleFunc("GET /showtodo", showTodo)

	http.ListenAndServe(":8080", nil)
}
