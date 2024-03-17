package main

import (
	"fmt"
	"log"
)

type Animals interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "Labrador",
	}

	gorilla := Gorilla{
		Name:  "Coco",
		Color: "Black",
	}

	PrintInfo(dog)
	PrintInfo(gorilla)
}

func (d Dog) Says() string {
	return "Woof"
}

func (d Dog) NumberOfLegs() int {
	return 4
}

func (g Gorilla) Says() string {
	return "Ugh"
}

func (g Gorilla) NumberOfLegs() int {
	return 2
}

func PrintInfo(a Animals) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
	log.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}
