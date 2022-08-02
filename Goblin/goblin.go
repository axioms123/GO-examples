//The rules for the game are as follows:
//
//    Each new hero starts with random max health points between 20 to 30.
//    Each new hero will start with a random attack power between 1 to 3.
//    Each new hero will start with a random defense power between 1-5.
//    Each new hero will start with an array of potions (int) with a max of 5 slots. Each slot set to 2.
//    Each new hero will start with 0 gold.
//
//    The hero will take a step, and possibly encounter a 'goblin.'
//        A goblin starts with 5-10 Health points, 2-3 attack power, and 1-2 defense power
//        If a goblin is found, the hero fights the goblin until either the hero or goblin's health points reach 0.
//        If the hero wins, reward the hero with 2 gold.
//
//    Each potion is worth their value in the array: 1 potion = 2 health points.
//    Drinking a potion will increase the hero's health points up to the max random health points assigned at the start of the game.
//
//    After every 10 steps:
//        The hero advances one level in the game.
//        The hero may visit the "potion shop" to buy more potions for 4 gold each.
//    When the hero dies, ask the protagonist if they would like to play again.
//        If the protagonist says yes, create a new character with the same gold they had when they died.
//        If the protagonist says no, print out the level the current character reached with the number of goblins slain!
//
package main

import (
	"math/rand"
	"fmt"
)

 type hero struct{
	health, attack, defense, gold, NextStep int
	potions []int
}

type goblin struct {
	health, attack, defense int
}
type game struct {
	protagonist hero
	antagonist goblin
}


func (g *game) createHero() {
	g.protagonist = hero{
		health:  rand.Intn(11) + 20,
		attack:  rand.Intn(3) + 1,
		defense: rand.Intn(5) + 1,
		potions: []int{2, 2, 2, 2, 2},
		gold:    0,
	}
}

func createGoblin() goblin {
	var newGoblin = goblin{
		health:  rand.Intn(6) + 5,
		attack:  rand.Intn(2) + 2,
		defense: rand.Intn(2) + 1,
	}
	return newGoblin
}


func fight(protagonist *hero, goblin *goblin) bool {
	for protagonist.health > 0 && goblin.health > 0 {
		protagonistAttack := protagonist.attack
		protagonistDefense := protagonist.defense
		goblinAttack := goblin.attack
		goblinDefense := goblin.defense


		for protagonistDefense > 0 && goblinAttack > 0 {
			protagonistDefense--
			goblinAttack--
		}
		if goblinAttack > 0 {
			for protagonist.health > 0 && goblinAttack > 0 {
				protagonist.health--
				goblinAttack--
			}
		}
		for goblinDefense > 0 && protagonistAttack > 0 {
			goblinDefense--
			protagonistAttack--
		}
		if protagonistAttack > 0 {
			for goblin.health > 0 && protagonistAttack > 0 {
				goblin.health--
				protagonistAttack--
			}
		}
	}
	//Notice if a protagonist won or lost 
	if protagonist.health > 0 {
		protagonist.gold += 2
		return true 
	} else {
		return false 
	}
}


func (h *hero) Step() {
	h.NextStep++
	var goblin goblin
	goblinChance := rand.Intn(2)
	if goblinChance == 1 {
		goblin = createGoblin()
		outcome := fight(h, &goblin)
		return outcome;
	}
}

func main() {
	fmt.Println("Goblin Tower Game")
}
