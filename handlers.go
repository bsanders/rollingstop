package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "math/rand"
    "strconv"
    "time"

    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func RollDie(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    dieNumber, err := strconv.Atoi(vars["dieNumber"])
    if err != nil { return }
    if dieNumber <= 0 { return }

    result := Roll{
        Result: rand.Intn(dieNumber) + 1,
        DieNumber: dieNumber,
        Date: time.Now(),
    }
    
    json.NewEncoder(w).Encode(result)
}