//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct {
	health, maxHealth uint // uint can only be positive
	energy, maxEnergy uint
	name              string
}

func (ply *Player) addHealth(hp uint) {
	if ply.health+hp >= ply.maxHealth {
		ply.health = ply.maxHealth
	} else {
		ply.health += hp
	}
}

func (ply *Player) damageHealth(dmg uint) {
	if ply.health < dmg {
		ply.health = 0
	} else {
		ply.health -= dmg
	}
}

func (ply Player) addEnergy(ep uint) Player {
	if ply.energy+ep >= ply.maxEnergy {
		return Player{ply.health, ply.maxHealth, ply.maxEnergy, ply.maxEnergy, ply.name}
	} else {
		return Player{ply.health, ply.maxHealth, ply.energy + ep, ply.maxEnergy, ply.name}
	}
}

func (ply *Player) useEnergy(usg uint) {
	// if ply.energy < usg {
	// 	ply.energy = 0
	// } else {
	// 	ply.energy -= usg
	// }
	ply.energy -= usg // artificial error

}

func main() {
	player1 := Player{60, 100, 10, 100, "Akko"}
	fmt.Println(player1)
	player1.addHealth(90)
	fmt.Println(player1)
	player1.damageHealth(10)
	fmt.Println(player1)
	player1.damageHealth(120)
	fmt.Println(player1)
	player1 = player1.addEnergy(20)
	fmt.Println(player1)
	player1.useEnergy(99999)
	fmt.Println(player1)
}
