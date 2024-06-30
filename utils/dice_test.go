package utils

import (
    "testing"
)

func TestRollDice(t *testing.T) {
    for i := 0; i < 100; i++ {
        result := RollDice()
        if result < 1 || result > 6 {
            t.Errorf("Expected dice roll between 1 and 6, got %d", result)
        }
    }
}
