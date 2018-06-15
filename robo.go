package main

import (
	"fmt"
	"math/rand"
	"mecha-wars/mecha"
	"mecha-wars/wars"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func fight(mech mecha.CombatMech, battlefield chan string, wg *sync.WaitGroup) {
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

func getReady(mechs []mecha.CombatMech, battlefield chan string, wg *sync.WaitGroup) {
	for _, mech := range mechs {
		go fight(mech, battlefield, wg)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Printf("Assembling Mecha combatants...\n")

	battlefield := make(chan string)
	mechs := wars.DraftMechs()

	wars.Introductions(mechs)
	getReady(mechs, battlefield, &wg)

	battlefield <- "start"

	wg.Wait()
	fmt.Printf("Battle over\n")
}
