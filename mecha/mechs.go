package mecha

type Wasp struct {
	standardCore
	jetpack
	lazer
}

func NewWasp(name string, serial string) Wasp {
	return Wasp{
		standardCore: NewStandardCore("Wasp", name, serial, 2),
	}
}

type Atlas struct {
	standardCore
	legs
	howitzer
}

func NewAtlas(name string, serial string) Atlas {
	return Atlas{
		standardCore: NewStandardCore("Atlas", name, serial, 4),
	}
}

type Marauder struct {
	massProducedCore
	legs
	machineGun
}

func NewMarauder(serial string) Marauder {
	return Marauder{
		massProducedCore: NewMassProducedCore("Marauder", serial, 3),
	}
}

type Ghost struct {
	standardCore
	jetpack
	cloak
}

func NewGhost(name string, serial string) Ghost {
	return Ghost{
		standardCore: NewStandardCore("Ghost", name, serial, 3),
	}
}
