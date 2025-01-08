package main

import (
    "log"
    "net/http"
    "time"

    "github.com/gorilla/mux"
    "go-rest-api/src/handlers"
    "go-rest-api/src/utils"
)

func main() {
    utils.RotateKey()

    go rotateKey()

    r := mux.NewRouter()
    r.HandleFunc("/key", handlers.GetKeyHandler).Methods("GET")
    r.HandleFunc("/key", handlers.ResetKeyHandler).Methods("POST")

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal(err)
    }
}

func rotateKey() {
    for {
        time.Sleep(24 * time.Hour)
        utils.RotateKey()
    }
}