package main

import (
	"fmt"
	"math/rand"
	"mecha-wars/mecha"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func prefix(m mecha.Mecha) {
	fmt.Printf("%s: ", m.GetDesignation())
}

func announce(m mecha.Mecha) {
	prefix(m)
	fmt.Printf("My name is %s,` %s Class Mech\n", m.GetDesignation(), m.GetClass())
	prefix(m)
	fmt.Printf("Ready for duty\n")
}

func act(mech mecha.CombatMech) string {
	a := rand.Intn(100)
	var action string
	locations := [4]string{"forwards", "backwards", "left", "right"}
	if a > 20 {
		fmt.Printf("%s attacks: ", mech.GetDesignation())
		mech.Attack()
		action = "attack"
	} else {
		l := rand.Intn(100)
		fmt.Printf("%s moves: ", mech.GetDesignation())
		mech.Move(locations[l%4])
		action = "move"
	}
	return action
}

func defend(mech mecha.CombatMech) int {
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

func fight(mech mecha.CombatMech, battlefield chan string) {
	defer wg.Done()

	for {
		action, ok := <-battlefield
		if !ok {
			fmt.Printf("Mech %s has won!\n", mech.GetDesignation())
			return
		}

		if action == "attack" {
			hp := defend(mech)
			if hp == 0 {
				close(battlefield)
				return
			}
		}
		battlefield <- act(mech)
	}
}

func main() {
	wg.Add(2)
	fmt.Printf("Assembling Mecha combatants...\n")

	battlefield := make(chan string)

	var mechA mecha.CombatMech
	var mechB mecha.CombatMech
	var mechs []mecha.CombatMech

	mechA = mecha.NewWasp("Viper", "V8")
	mechB = mecha.NewAtlas("Brunhilda", "WG1")
	mechs = append(mechs, mechA, mechB)

	go fight(mechA, battlefield)
	go fight(mechB, battlefield)

	battlefield <- "start"

	wg.Wait()
	fmt.Printf("Battle over\n")
}
