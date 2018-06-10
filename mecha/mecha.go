package mecha

type Class string

type Mecha interface {
	GetClass() Class
	GetDesignation() string
	Damage(dmg int) int
	GetHitpoints() int
}

type Moveable interface {
	Move(location string)
}

type Weaponized interface {
	Attack()
}

type Infiltrator interface {
	Hide()
}

type CombatMech interface {
	Mecha
	Moveable
	Weaponized
}

type ReconMech interface {
	Mecha
	Moveable
	Infiltrator
}
