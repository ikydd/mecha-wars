package actions

import (
	"fmt"
	"math/rand"
	"mecha-wars/mecha"
	"time"
)

var locations []string

func init() {
	locations = []string{"the left", "the right"}
	rand.Seed(time.Now().UnixNano())
}

func Introduce(m mecha.Mecha) {
	fmt.Printf("Designate %s, %s Class Mech\n", m.GetDesignation(), m.GetClass())
	fmt.Printf("Ready for duty\n")
}

func Defend(mech mecha.CombatMech) int {
	n := rand.Intn(100)
	var hp int
	name := mech.GetDesignation()
	if n < 20 {
		fmt.Println("Attack misses!")
		hp = mech.GetHitpoints()
	} else {
		hp = mech.Damage(1)
		fmt.Printf("Attack hits! %d HP remaining on %s\n", hp, name)
		if hp == 0 {
			fmt.Printf("Mech %s destroyed\n", name)
		}
	}
	return hp
}

func attack(mech mecha.CombatMech) {
	fmt.Printf("%s attacks: ", mech.GetDesignation())
	mech.Attack()
}

func move(mech mecha.CombatMech) {
	l := rand.Intn(100)
	fmt.Printf("%s moves: ", mech.GetDesignation())
	mech.Move(locations[l%len(locations)])
}

func Act(mech mecha.CombatMech) string {
	a := rand.Intn(100)
	var action string
	if a > 20 {
		attack(mech)
		action = "attack"
	} else {
		move(mech)
		action = "move"
	}
	return action
}

func Celebrate(m mecha.Mecha) {
	fmt.Printf("Mech %s has won!\n", m.GetDesignation())
}
