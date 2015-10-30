package main

import "time"

type Roll struct {
    Result      int       `json:"result"`
    DieNumber   int       `json:"dienumber"`
    Date        time.Time `json:"date"`
}

type Rolls []Roll
