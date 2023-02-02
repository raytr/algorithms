package main

import "fmt"

//import "fmt"

type Reptile interface {
	Lay() ReptileEgg
}

type ReptileCreator func() Reptile

type ReptileEgg struct {
	CreateReptile ReptileCreator
	hatchCount    int
}

func (egg *ReptileEgg) Hatch() Reptile {
	if egg.hatchCount > 0 {
		return nil
	}
	egg.hatchCount++
	return egg.CreateReptile()
}

type FireDragon struct {
}

func (d FireDragon) Lay() ReptileEgg {
	egg := ReptileEgg{
		CreateReptile: func() Reptile {
			return FireDragon{}
		},
	}

	return egg
}

func main() {
	var fireDragon FireDragon
	var egg ReptileEgg = fireDragon.Lay()
	var childDragon Reptile = egg.Hatch()
	fmt.Println(childDragon)
}
