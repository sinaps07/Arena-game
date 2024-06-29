package player

type Player struct {
    Name     string
    Health   int
    Strength int
    Attack   int
}

func (p *Player) IsAlive() bool {
    return p.Health > 0
}

func (p *Player) TakeDamage(damage int) {
    p.Health -= damage
    if p.Health < 0 {
        p.Health = 0
    }
}