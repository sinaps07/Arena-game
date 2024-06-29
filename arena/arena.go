package arena

import (
	"fmt"
	"math/rand"
	"swiggy_game/player"
	"swiggy_game/utils"
	"time"
)

type Arena struct {
    PlayerA *player.Player
    PlayerB *player.Player
}

func NewArena(playerA, playerB *player.Player) *Arena {
    return &Arena{
        PlayerA: playerA,
        PlayerB: playerB,
    }
}

func (a *Arena) StartBattle() {
    rand.Seed(time.Now().UnixNano())
    attacker, defender := a.determineFirstAttacker()
    for attacker.IsAlive() && defender.IsAlive() {
        a.fightRound(attacker, defender)
        attacker, defender = defender, attacker
    }
    a.printWinner(attacker, defender)
}

func (a *Arena) determineFirstAttacker() (*player.Player, *player.Player) {
    if a.PlayerA.Health < a.PlayerB.Health {
        return a.PlayerA, a.PlayerB
    }
    return a.PlayerB, a.PlayerA
}

func (a *Arena) fightRound(attacker, defender *player.Player) {
    attackRoll := utils.RollDice()
    defendRoll := utils.RollDice()

    attackDamage := attacker.Attack * attackRoll
    defense := defender.Strength * defendRoll
    damage := attackDamage - defense

    if damage > 0 {
        defender.TakeDamage(damage)
    }

    fmt.Printf("%s attacks %s\n", attacker.Name, defender.Name)
    fmt.Printf("%s rolls %d, %s rolls %d\n", attacker.Name, attackRoll, defender.Name, defendRoll)
    fmt.Printf("Damage dealt: %d, %s's health: %d\n", damage, defender.Name, defender.Health)
}

func (a *Arena) printWinner(attacker, defender *player.Player) {
    if attacker.IsAlive() {
        fmt.Printf("%s wins!\n", attacker.Name)
    } else {
        fmt.Printf("%s wins!\n", defender.Name)
    }
}