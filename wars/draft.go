package wars

import (
	"math/rand"
	"mecha-wars/mecha"
)

func getRandomMech(roster []mecha.CombatMech) (mecha.CombatMech, []mecha.CombatMech) {
	num := rand.Intn(100)
	index := num % len(roster)
	mech := roster[index]
	roster = append(roster[:index], roster[index+1:]...)

	return mech, roster
}

func getRandomMechs(roster []mecha.CombatMech, num int) []mecha.CombatMech {
	var mechs []mecha.CombatMech

	for i := 0; i < num; i++ {
		var mech mecha.CombatMech
		mech, roster = getRandomMech(roster)
		mechs = append(mechs, mech)
	}

	return mechs
}

func DraftMechs() []mecha.CombatMech {
	roster := []mecha.CombatMech{
		mecha.NewWasp("Viper", "V8"),
		mecha.NewWasp("Mercury", "9GG"),
		mecha.NewAtlas("Brunhilda", "WG1"),
		mecha.NewAtlas("Suzy", "309"),
		mecha.NewMarauder("W355"),
	}

	return getRandomMechs(roster, 2)
}
