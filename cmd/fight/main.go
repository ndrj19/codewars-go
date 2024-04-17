package main

import (
	"fmt"
)

type Fighter struct {
	Name            string
	Health          int
	DamagePerAttack int
}

func main() {
	result := DeclareWinner(Fighter{"Lew", 10, 2}, Fighter{"Harry", 5, 4}, "Lew")
	fmt.Println(result)
}

func DeclareWinner(fighter1 Fighter, fighter2 Fighter, firstAttacker string) string {
	for fighter1.Health > 0 && fighter2.Health > 0 {
		if fighter1.Name == firstAttacker {
			fighter2.Health -= fighter1.DamagePerAttack
			if fighter2.Health <= 0 {
				return fighter1.Name
			}
			fighter1.Health -= fighter2.DamagePerAttack
			if fighter1.Health <= 0 {
				return fighter2.Name
			}
		} else {
			fighter1.Health -= fighter2.DamagePerAttack
			if fighter1.Health <= 0 {
				return fighter2.Name
			}
			fighter2.Health -= fighter1.DamagePerAttack
			if fighter2.Health <= 0 {
				return fighter1.Name
			}
		}
	}
	return ""
}
