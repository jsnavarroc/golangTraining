package main

import "fmt"

type Animal interface {
	Says() string 
	NumberOfLegs() int
}

type Dog struct {
	Name string
	Breed string
}

type Gorilla struct {
	Name string
	Color string
	NumberOfTheeth int
}

func main(){
	dog := Dog {
		Name: "Samson",
		Breed: "German Shephered",
	}
	PrintAnimal(&dog)

	gorilla := Gorilla {
		Name: "Jock",
		Color: "grey",
		NumberOfTheeth: 38,
	}

	PrintAnimal(&gorilla)
}

func PrintAnimal (a Animal) {
	fmt.Println("This animal says", a.Says(),"and has", a.NumberOfLegs(), "legs")
}

func (d *Dog) Says() string	{
	return "Woof"
}
func (d *Dog) NumberOfLegs() int	{
	return 4
}
func (d *Gorilla) Says() string	{
	return "Ugh"
}
func (d *Gorilla) NumberOfLegs() int	{
	return 2
}
