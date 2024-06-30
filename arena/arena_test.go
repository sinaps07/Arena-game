package arena

import (
    "testing"
    "swiggy_game/player"
)

func TestArenaDetermineFirstAttacker(t *testing.T) {
    playerA := &player.Player{Name: "Player A", Health: 50, Strength: 5, Attack: 10}
    playerB := &player.Player{Name: "Player B", Health: 100, Strength: 10, Attack: 5}
    gameArena := NewArena(playerA, playerB)

    attacker, _ := gameArena.determineFirstAttacker()
    if attacker != playerA {
        t.Errorf("Expected Player A to attack first, got %s", attacker.Name)
    }

    playerA.Health = 100
    playerB.Health = 50
    attacker, _ = gameArena.determineFirstAttacker()
    if attacker != playerB {
        t.Errorf("Expected Player B to attack first, got %s", attacker.Name)
    }
}

func TestArenaFightRound(t *testing.T) {
    playerA := &player.Player{Name: "Player A", Health: 50, Strength: 5, Attack: 10}
    playerB := &player.Player{Name: "Player B", Health: 100, Strength: 10, Attack: 5}
    gameArena := NewArena(playerA, playerB)

    gameArena.fightRound(playerA, playerB)

    if playerB.Health >= 100 {
        t.Errorf("Expected Player B's health to be less than 100 after attack, got %d", playerB.Health)
    }
}

func TestArenaStartBattle(t *testing.T) {
    playerA := &player.Player{Name: "Player A", Health: 50, Strength: 5, Attack: 10}
    playerB := &player.Player{Name: "Player B", Health: 100, Strength: 10, Attack: 5}
    gameArena := NewArena(playerA, playerB)

    gameArena.StartBattle()

    if playerA.IsAlive() && playerB.IsAlive() {
        t.Errorf("Expected one player to be dead")
    }

    if !playerA.IsAlive() && !playerB.IsAlive() {
        t.Errorf("Expected one player to be alive")
    }
}

func TestArenaNegativeHealth(t *testing.T) {
    playerA := &player.Player{Name: "Player A", Health: -50, Strength: 5, Attack: 10}
    playerB := &player.Player{Name: "Player B", Health: 100, Strength: 10, Attack: 5}
    gameArena := NewArena(playerA, playerB)

    gameArena.StartBattle()

    if playerA.IsAlive() {
        t.Errorf("Expected Player A to be dead due to negative health")
    }
}

func TestArenaHighDamage(t *testing.T) {
    playerA := &player.Player{Name: "Player A", Health: 50, Strength: 5, Attack: 1000}
    playerB := &player.Player{Name: "Player B", Health: 100, Strength: 10, Attack: 5}
    gameArena := NewArena(playerA, playerB)

    gameArena.StartBattle()

    if playerB.IsAlive() {
        t.Errorf("Expected Player B to be dead due to high damage from Player A")
    }
}
