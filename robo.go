package main

import (
	"fmt"
	"math/rand"
	"mecha-wars/mecha"
	"mecha-wars/wars"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func fight(mech mecha.CombatMech, battlefield chan string) {
	defer wg.Done()

	for {
		action, ok := <-battlefield
		if !ok {
			wars.Celebrate(mech)
			return
		}

		if action == "attack" {
			hp := wars.Defend(mech)
			if hp == 0 {
				close(battlefield)
				return
			}
		}
		battlefield <- wars.Act(mech)
	}
}

func introductions(mechs []mecha.CombatMech) {
	for _, mech := range mechs {
		wars.Introduce(mech)
	}
}

func getReady(mechs []mecha.CombatMech, battlefield chan string) {
	for _, mech := range mechs {
		go fight(mech, battlefield)
	}
}

func main() {
	wg.Add(2)
	fmt.Printf("Assembling Mecha combatants...\n")

	battlefield := make(chan string)
	mechs := wars.DraftMechs()

	introductions(mechs)
	getReady(mechs, battlefield)

	battlefield <- "start"

	wg.Wait()
	fmt.Printf("Battle over\n")
}
