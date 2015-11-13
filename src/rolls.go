package main

import "time"

type Roll struct {
    Result      int       `json:"result"`
    DieNumber   int       `json:"dienumber"`
    Count       int       `json:"count"`
    Modifier    int       `json:"modifier"`
    Date        time.Time `json:"date"`
}

type Rolls []Roll
