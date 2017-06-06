package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/", home)
    router.HandleFunc("/todos", list)
    router.HandleFunc("/todos/{todoId}", one)
    log.Fatal(http.ListenAndServe(":9999", router))
}

func home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome")
}

func list(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Todo Index!")
}

func one(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}