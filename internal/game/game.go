package game

import (
	"fmt"
	"game/internal/colors"
	"game/internal/config"
	"game/internal/player"
	"log"
)

type Game struct {
	myPlayer *player.Player
}

func InitGame() *Game {
	var name string
	fmt.Print(colors.GetCyanText("Введите имя игрока: "))
	fmt.Scan(&name)
	currentLocation, err := config.LoadLocations("config/game.json", "Дом")
	if err != nil {
		log.Fatal(err)
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
func (g *Game) GetPlayerInfo() {
	output := fmt.Sprintf(`
        +----------------------------+
        | Имя: %s
        | Здоровье: %d
        | Дамаг: %d
        +----------------------------+
    `, g.myPlayer.Name, g.myPlayer.Health, g.myPlayer.Damage)

	fmt.Println(output)
}

func (g *Game) GoToLocation(locationName string) {

	for _, location := range g.myPlayer.CurrentLocation.Locations {
		fmt.Println(location.IsOpen)
		if location.Name == locationName {
			err := g.myPlayer.GoToLocation(location)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
			return
		}
	}
	for _, zone := range g.myPlayer.CurrentLocation.Zones {
		if zone.Name == locationName {
			g.myPlayer.GoToLocation(zone)
			return
		}
	}
	fmt.Println("Нет такой локации")
}
