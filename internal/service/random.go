package service

import (
    "math/rand"
    "time"
)

var limit = 100;

func GetRandom() int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(limit)
}

func SetLimit(newLimit int) {
    if newLimit > 0 {
        limit = newLimit
    }
}