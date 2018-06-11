package main

import (
	"fmt"
	"math/rand"
	"mecha-wars/mecha"
	"mecha-wars/mecha/actions"
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

func getRandomMech(roster []mecha.CombatMech) (mecha.CombatMech, []mecha.CombatMech) {
	num := rand.Intn(100)
	index := num % len(roster)
	mech := roster[index]
	roster = append(roster[:index], roster[index+1:]...)

	return mech, roster
}

func loadMechs() []mecha.CombatMech {

	var firstMech mecha.CombatMech
	var secondMech mecha.CombatMech

	roster := []mecha.CombatMech{
		mecha.NewWasp("Viper", "V8"),
		mecha.NewWasp("Mercury", "9GG"),
		mecha.NewAtlas("Brunhilda", "WG1"),
		mecha.NewAtlas("Suzy", "309"),
		mecha.NewMarauder("W355"),
	}

	firstMech, roster = getRandomMech(roster)
	secondMech, roster = getRandomMech(roster)

	return []mecha.CombatMech{
		firstMech,
		secondMech,
	}
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
