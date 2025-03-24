package game

import (
	"fmt"
	"game/internal/location"
	"game/internal/player"
)

type Game struct {
	myPlayer *player.Player
}

func InitGame() *Game {
	var name string
	fmt.Print("Введите имя игрока: ")
	fmt.Scan(&name)
	currentLocation := location.Location{
		Name:        "Дом",
		Description: "Мой дом",
		IsOpen:      true,
	}
	myPlayer := player.NewPlayer(name, currentLocation)
	return &Game{myPlayer: myPlayer}
}

func (g *Game) GetCurrentLocation() {
	fmt.Println(g.myPlayer.CurrentLocation.GetLocationDraw())
}

func (g *Game) GetInventory() {
	fmt.Println(g.myPlayer.Inventory.GetItems())
}
