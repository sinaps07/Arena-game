package player

import (
    "testing"
)

func TestPlayer(t *testing.T) {
    player := &Player{Name: "Test Player", Health: 100, Strength: 10, Attack: 20}

    if !player.IsAlive() {
        t.Errorf("Expected player to be alive")
    }

    player.TakeDamage(50)
    if player.Health != 50 {
        t.Errorf("Expected health to be 50, got %d", player.Health)
    }

    player.TakeDamage(60)
    if player.Health != 0 {
        t.Errorf("Expected health to be 0, got %d", player.Health)
    }

    if player.IsAlive() {
        t.Errorf("Expected player to be dead")
    }
}
