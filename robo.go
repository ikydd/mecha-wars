package main

import (
	"fmt"
	"mecha-wars/mecha"
	"mecha-wars/mecha/actions"
	"sync"
)

var wg sync.WaitGroup

func fight(mech mecha.CombatMech, battlefield chan string) {
	defer wg.Done()

	for {
		action, ok := <-battlefield
		if !ok {
			actions.Celebrate(mech)
			return
		}

		if action == "attack" {
			hp := actions.Defend(mech)
			if hp == 0 {
				close(battlefield)
				return
			}
		}
		battlefield <- actions.Act(mech)
	}
}

func introductions(mechs []mecha.CombatMech) {
	for _, mech := range mechs {
		actions.Introduce(mech)
	}
}

func getReady(mechs []mecha.CombatMech, battlefield chan string) {
	for _, mech := range mechs {
		go fight(mech, battlefield)
	}
}

func loadMechs() []mecha.CombatMech {
	var mechA mecha.CombatMech
	var mechB mecha.CombatMech
	var mechs []mecha.CombatMech

	mechA = mecha.NewWasp("Viper", "V8")
	mechB = mecha.NewAtlas("Brunhilda", "WG1")
	mechs = append(mechs, mechA, mechB)

	return mechs
}

func main() {
	wg.Add(2)
	fmt.Printf("Assembling Mecha combatants...\n")

	battlefield := make(chan string)
	mechs := loadMechs()

	introductions(mechs)
	getReady(mechs, battlefield)

	battlefield <- "start"

	wg.Wait()
	fmt.Printf("Battle over\n")
}
