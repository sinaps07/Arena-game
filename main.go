package main

import (
	"fmt"
	"swiggy_game/arena"
	"swiggy_game/player"
)

func main() {
    var nameA, nameB string
    var healthA, healthB, strengthA, strengthB, attackA, attackB int

    fmt.Println("Enter details for Player A:")
    fmt.Print("Name: ")
    fmt.Scan(&nameA)
    fmt.Print("Health: ")
    fmt.Scan(&healthA)
    fmt.Print("Strength: ")
    fmt.Scan(&strengthA)
    fmt.Print("Attack: ")
    fmt.Scan(&attackA)

    fmt.Println("Enter details for Player B:")
    fmt.Print("Name: ")
    fmt.Scan(&nameB)
    fmt.Print("Health: ")
    fmt.Scan(&healthB)
    fmt.Print("Strength: ")
    fmt.Scan(&strengthB)
    fmt.Print("Attack: ")
    fmt.Scan(&attackB)

    playerA := &player.Player{Name: nameA, Health: healthA, Strength: strengthA, Attack: attackA}
    playerB := &player.Player{Name: nameB, Health: healthB, Strength: strengthB, Attack: attackB}

    gameArena := arena.NewArena(playerA, playerB)
    gameArena.StartBattle()
}
