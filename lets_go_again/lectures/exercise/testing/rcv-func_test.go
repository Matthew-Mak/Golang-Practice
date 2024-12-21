//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

package main

import "testing"

func TestAddHealth(t *testing.T) {
	p := Player{10, 100, 10, 100, "Constanze"}
	p.addHealth(100)
	if p.health > p.maxHealth {
		//t.Errorf("health can't be above maximum value, current: %v want %v", p.health, p.maxHealth)
		t.Fatalf("health can't be above maximum value, current: %v want %v", p.health, p.maxHealth) // FatalF is better here as it will also terminate the program
	}
}

func TestDamageHealth(t *testing.T) {
	p := Player{10, 100, 10, 100, "Constanze"}
	p.damageHealth(999)
	if p.health != 0 {
		t.Fatalf("health can't be other than 0, current: %v want %v", p.health, 0)
	}
}

func TestAddEnergy(t *testing.T) {
	p := Player{10, 100, 10, 100, "Constanze"}
	p = p.addEnergy(100)
	if p.energy > p.maxEnergy {
		t.Fatalf("energy can't be above maximum value, current: %v want %v", p.energy, p.maxEnergy)
	}
}

func TestDamageEnergy(t *testing.T) {
	p := Player{10, 100, 10, 100, "Constanze"}
	p.useEnergy(999)
	if p.energy != 0 {
		t.Fatalf("energy can't be other than 0, current: %v want %v", p.energy, 0)
	}
}
