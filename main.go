package main

import (
	"fmt"
	"os"
)

type Villains struct {
	Type string
}

type Player struct {
	name string
}

type Inventory struct {
	breakSword string
	pickaxe    string
	torch      int
	metal      string
	armour     string
}

func (p *Player) Mining(i *Inventory) (string, string) {
	return "Player " + p.name + " mined " + i.metal, ""
}

func (p *Player) Damage(v *Villains) (string, string) {
	return "Player " + p.name + " kill " + v.Type, ""
}

func (p Player) Craft(i *Inventory) (string, string) {
	return "Player " + p.name + " crafted " + i.armour, ""
}

func (p Player) Repair(i *Inventory) (string, string) {
	return "Player " + p.name + " repair " + i.breakSword, ""

}

func main() {
	p := &Player{name: "Steve"}

	boss := &Villains{Type: "Ender Dragon"}
	zombie := &Villains{Type: "Zombie"}
	skeleton := &Villains{Type: "Skeleton"}

	inventory := &Inventory{
		breakSword: "Iron sword",
		pickaxe:    "Rock pickaxe",
		metal:      "metal",
		armour:     "Metal armour",
	}

	fmt.Println("MINECRAFT: STORY GAME")

	fmt.Println("Steve woke up at the entrance to the cave. He only remembers his name. Next to him is a backpack in which he finds ", inventory, ". It is dark in the cave")

	choice1 := "1.Start mining"
	choice2 := "2.Look for a way out"

	var choicePlayer1 int
	fmt.Println(choice1, choice2)

	fmt.Scanln(&choicePlayer1)

	if choicePlayer1 == 1 {
		fmt.Println(p.Mining(inventory))
		fmt.Println(p.Craft(inventory))
		fmt.Println(p.Repair(inventory))
	} else if choicePlayer1 == 2 {
		fmt.Println("Steve start Looking for a way out, but right at the exit steaven see zombie. Steaven died from zombie because his sword broken.")
		fmt.Println("GAME OVER!")
		os.Exit(1)
	}

	fmt.Println("Steve is looking for a way out but gets on zombies")
	fmt.Println(p.Damage(zombie))

	fmt.Println("Steve left the cave and entered the Forest. In long wanderings, he found a traveler who was attacked by zombies and skeletons.")
	choice3 := "1.Help Wanderer"
	choice4 := "2.pass by"

	var choicePlayer2 int
	fmt.Println(choice3, choice4)

	fmt.Scanln(&choicePlayer2)

	if choicePlayer2 == 1 {
		fmt.Println(p.Damage(zombie))
		fmt.Println(p.Damage(skeleton))
	} else if choicePlayer2 == 2 {
		fmt.Println("Steve died from hunger")
		fmt.Println("GAME OVER!")
		os.Exit(1)
	}

	fmt.Println("The Stranger thanked Steve, fed him and took him to the village. In the village, " +
		"everyone told him about the secret world from which the dragon comes and attacks their village." +
		" Steve decided to help the residents and went to fight the dragon.")

	fmt.Println(p.Damage(boss))

	fmt.Println("Steve defeated the dragon and freed the village")
	fmt.Println("Happy end!")

}
