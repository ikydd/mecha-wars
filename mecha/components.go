package mecha

import "fmt"

type howitzer struct {
}

func (h howitzer) Attack() {
	fmt.Println("BOOM!")
}

type machineGun struct {
}

func (m machineGun) Attack() {
	fmt.Println("Dakka dakka dakka!")
}

type lazer struct {
}

func (l lazer) Attack() {
	fmt.Println("Zaaaap!")
}

type jetpack struct {
}

func (j jetpack) Move(location string) {
	fmt.Printf("Flying to %s\n", location)
}

type legs struct {
}

func (l legs) Move(location string) {
	fmt.Printf("Walking to %s\n", location)
}

type cloak struct {
}

func (c cloak) Hide() {
	fmt.Println("Going invis...")
}

type camoflage struct {
}

func (c camoflage) Hide() {
	fmt.Println("Camoflaging now...")
}
