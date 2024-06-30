package arena

import (
	"swiggy_game/player"
	"testing"
)

func TestArena(t *testing.T) {
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
