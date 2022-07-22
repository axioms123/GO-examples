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
//    When the hero dies, ask the player if they would like to play again.
//        If the player says yes, create a new character with the same gold they had when they died.
//        If the player says no, print out the level the current character reached with the number of goblins slain!
//
package main

import (
	"fmt"
	"math/random"
	"time"
)

 type player struct{
	health string
	attack, defense, gold int16
	potions []
	go

//constructor function
func(std *Student) fill_defaults() {
	//setting default values
	//if no values present
	if std.name == "" {
		std.name = "character"
	}
}
func main() {
	import 
}
