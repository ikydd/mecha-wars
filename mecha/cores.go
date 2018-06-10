package mecha

import "fmt"

type core struct {
	className Class
	hitpoints int
}

func (c core) GetClass() Class {
	return c.className
}

func (c core) GetHitpoints() int {
	return c.hitpoints
}

func (c *core) Damage(dmg int) int {
	c.hitpoints--
	return c.GetHitpoints()
}

type standardCore struct {
	*core
	name   string
	serial string
}

func NewStandardCore(className Class, name string, serial string, hitpoints int) standardCore {
	return standardCore{
		core: &core{
			className: className,
			hitpoints: hitpoints,
		},
		name:   name,
		serial: serial,
	}
}

func (c standardCore) GetDesignation() string {
	return fmt.Sprintf("%s %s", c.name, c.serial)
}

type massProducedCore struct {
	*core
	serial string
}

func (c massProducedCore) GetDesignation() string {
	return fmt.Sprintf("%s %s", c.className, c.serial)
}

func NewMassProducedCore(className Class, serial string, hitpoints int) massProducedCore {
	return massProducedCore{
		core: &core{
			className: className,
			hitpoints: hitpoints,
		},
		serial: serial,
	}
}
