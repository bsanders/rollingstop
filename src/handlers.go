package main

import (
    "log"
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

    count, err := strconv.Atoi(r.FormValue("count"))
    //fmt.Fprintln(w, err)
    if err != nil { count = 1 }
    if count <= 0 { return }

    param_value := r.FormValue("modifier")
    modifier := 0
    // note that if the param_value started with '+', this was converted to a ' '
    if param_value == "" {
        
    } else {
        modifier, err = strconv.Atoi(param_value)
        if err != nil { return }
    }

    var my_rolls Rolls
    for i := 0; i < count; i++ {
        this_roll := Roll {
            Result: rand.Intn(dieNumber) + 1,
            DieNumber: dieNumber,
            Count: count,
            Modifier: modifier,
            Date: time.Now(),
        }
        my_rolls = append(my_rolls, this_roll)
    }

    for i, roll := range my_rolls {
        total := roll.Result + roll.Modifier
        log.Printf(
            "die_index=%d\tsum=%d\troll=%d\tdie=d%d\tcount=%d\tmodfier=%d\t%s",
            i,
            total,
            roll.Result,
            roll.DieNumber,
            roll.Count,
            roll.Modifier,
            roll.Date,
        )
    }

    json.NewEncoder(w).Encode(my_rolls)
}