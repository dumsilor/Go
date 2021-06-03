/*
problem:
1. Program will have 3 animal hardcoded
2. Program will promot user wiht ">" sign
3. User will enter a animal name and a activity in two strings
4. Program will print our the name of the animal along with the requested activity


Solution:
1. Make a type called animal which is a struct
2. Animal will contian 3 string fields
	a. Food
	b. Locomotion
	c. noise
3. Make 3 methods:
	a. Eat()
	b. Move()
	c. Speak()
4. Receiver type of all methods will be animal type

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type animal struct {
	food  string
	move  string
	noise string
}

func (species *animal) Eat() {
	fmt.Print(species.food)
}

func (species *animal) Move() {
	fmt.Print(species.move)
}

func (species *animal) Speak() {
	fmt.Print(species.noise)
}

func store_data(species_desc []string) {
	var new_object animal
	switch species_desc[0] {
	case "cow":
		new_object.food = "grass"
		new_object.move = "walk"
		new_object.noise = "moo"
	case "bird":
		new_object.food = "worms"
		new_object.move = "fly"
		new_object.noise = "peep"
	case "snake":
		new_object.food = "mice"
		new_object.move = "sither"
		new_object.noise = "hss"

	}

	if species_desc[1] == "eat" {
		new_object.Eat()
	} else if species_desc[1] == "move" {
		new_object.Move()
	} else if species_desc[1] == "noise" {
		new_object.Speak()
	}

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\n> ")
		a, _ := reader.ReadString('\n')
		s := strings.Split(strings.TrimSpace(a), " ")

		store_data(s)
	}

}
