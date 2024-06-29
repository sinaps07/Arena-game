package utils

import "math/rand"

func RollDice() int {
    return rand.Intn(6) + 1
}