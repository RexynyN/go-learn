package main

import "fmt"

type Player struct {
	health int
}

func (player *Player) takeDamageFromExplosion(dmg int) {
	fmt.Println("player is taking damage from explosion")
	player.health -= dmg
}

func takeDamageFromExplosion(player *Player, dmg int) {
	fmt.Println("player is taking damage from explosion")
	player.health -= dmg
}

func main() {
	player := &Player{
		health: 100,
	}

	fmt.Printf("before explosion damage: %+v\n", player)
	// player = nil // if you dereference the pointer bad things happen to good people

	takeDamageFromExplosion(*&*&*&*&*&*&*&*&*&*&*&*&*&*&*&*&player, 10)
	fmt.Printf("after explosion damage: %+v\n", player)

	player.takeDamageFromExplosion(10)
	fmt.Printf("after explosion damage: %+v\n", player)

}

type BigData struct {
	// 500mb
}
