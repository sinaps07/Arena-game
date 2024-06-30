package player

import (
    "testing"
)

func TestPlayerInitialization(t *testing.T) {
    player := &Player{Name: "Test Player", Health: 100, Strength: 10, Attack: 20}
    
    if player.Name != "Test Player" {
        t.Errorf("Expected player name to be 'Test Player', got %s", player.Name)
    }
    
    if player.Health != 100 {
        t.Errorf("Expected health to be 100, got %d", player.Health)
    }

    if player.Strength != 10 {
        t.Errorf("Expected strength to be 10, got %d", player.Strength)
    }

    if player.Attack != 20 {
        t.Errorf("Expected attack to be 20, got %d", player.Attack)
    }
}

func TestPlayerTakeDamage(t *testing.T) {
    player := &Player{Name: "Test Player", Health: 100, Strength: 10, Attack: 20}
    
    player.TakeDamage(50)
    if player.Health != 50 {
        t.Errorf("Expected health to be 50, got %d", player.Health)
    }

    player.TakeDamage(60)
    if player.Health != 0 {
        t.Errorf("Expected health to be 0, got %d", player.Health)
    }
}

func TestPlayerIsAlive(t *testing.T) {
    player := &Player{Name: "Test Player", Health: 100, Strength: 10, Attack: 20}

    if !player.IsAlive() {
        t.Errorf("Expected player to be alive")
    }

    player.TakeDamage(100)
    if player.IsAlive() {
        t.Errorf("Expected player to be dead")
    }
}


func TestPlayerExcessiveDamage(t *testing.T) {
    player := &Player{Name: "Test Player", Health: 100, Strength: 10, Attack: 20}
    
    player.TakeDamage(200)
    if player.Health != 0 {
        t.Errorf("Expected health to be 0, got %d", player.Health)
    }
}
