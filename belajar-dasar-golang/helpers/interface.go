package helpers

import "fmt"

// definisi interface
type Geng interface {
	Strength() 	int 
	Punch() 	int
}

// implement interface
type Person struct {
	Name, Role 		string
	Blood, Attack	int
} 

func (p Person) Strength() int {
	if p.Role == "Assassin" {
		return p.Blood * 10
	} else if p.Role == "Tank" {
		return p.Blood * 30
	}else{
		return p.Blood * 20
	}
}

func (p Person) Punch() int {
	if p.Role == "Assassin" {
		return p.Attack * 30
	} else if p.Role == "Tank" {
		return p.Attack * 5
	}else{
		return p.Attack * 15
	}
}

// polimorfisme
func PrintStrengthAndPunch(p Person) {
	fmt.Println(p.Name,"strength is : ", p.Strength())
	fmt.Println(p.Name,"punch is : ", p.Punch())
}