package main

import (
	"fmt"
	"math/rand/v2"
)

// Create a struct for an RPG creature. Give the struct data members for:
// Name
// Dexterity
// Strength
// Intelligence

// Create a NewRPGCreature(name string) function that returns a creature with a custom name and the value of 1 for all attributes.

// Create a RandomRPGCreature(name string) function that returns a creature with a custom name, and a random value between 1-9 for all attributes.

type Creature struct {
	Name         string
	Dexterity    int
	Strength     int
	Intelligence int
}

func main() {
	var customName string
	fmt.Print("Enter a name for a creature with all stats set to 1: ")
	fmt.Scanln(&customName)
	fmt.Println(NewRPGCreature(customName))

	var anotherName string
	fmt.Print("Enter a name for a creature with random stats: ")
	fmt.Scanln(&anotherName)
	fmt.Print(RandomRPGCreature(anotherName))
}

func NewRPGCreature(name string) Creature {
	return Creature{Name: name, Dexterity: 1, Strength: 1, Intelligence: 1}
}

func RandomRPGCreature(name string) Creature {
	return Creature{
		Name:         name,
		Dexterity:    rand.IntN(9) + 1,
		Strength:     rand.IntN(9) + 1,
		Intelligence: rand.IntN(9) + 1}

}
