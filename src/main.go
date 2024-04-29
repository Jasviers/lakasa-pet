package main

import (
	"fmt"
	"time"
	"github.com/google/uuid"
)

type LakasaPet struct {
	id 			  string
	name      string
	hunger    int
	happiness int
	energy    int
}

func NewLakasaPet(name string) *LakasaPet {
	return &LakasaPet{
		id: 			 uuid.New().String(),
		name:      name,
		hunger:    0,
		happiness: 100,
		energy:    100,
	}
}

func (pet *LakasaPet) Feed() {
	pet.hunger -= 10
	if pet.hunger < 0 {
		pet.hunger = 0
	}
	pet.happiness += 5
	pet.energy -= 5
}

func (pet *LakasaPet) Play() {
	pet.hunger += 5
	pet.happiness += 10
	pet.energy -= 10
}

func (pet *LakasaPet) Sleep() {
	pet.hunger += 5
	pet.happiness -= 5
	pet.energy += 20
	if pet.energy > 100 {
		pet.energy = 100
	}
}

func (pet *LakasaPet) Update() {
	pet.hunger += 5
	pet.happiness -= 5
	pet.energy -= 10
}

func (pet *LakasaPet) IsDead() bool {
	return pet.hunger >= 100 || pet.happiness <= 0 || pet.energy <= 0
}

func main() {
	LakasaPet := NewLakasaPet("Mario")

	for !LakasaPet.IsDead() {
		fmt.Println("Name:", LakasaPet.name)
		fmt.Println("Hunger:", LakasaPet.hunger)
		fmt.Println("Happiness:", LakasaPet.happiness)
		fmt.Println("Energy:", LakasaPet.energy)
		fmt.Println("--------------------")

		LakasaPet.Update()

		time.Sleep(60 * time.Second)
	}

	fmt.Println("Pet died!")
}