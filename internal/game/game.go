package game

import (
	"fmt"
	"game/internal/colors"
	"game/internal/config"
	"game/internal/location"
	"game/internal/player"
	"log"
	"strings"
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
        | Оружие: %s
        +----------------------------+
    `, g.myPlayer.Name, g.myPlayer.Health, g.myPlayer.Damage, g.myPlayer.Weapon.Name)

	fmt.Println(output)
}

func (g *Game) GoToLocation(locationName string) {
	var location *location.Location

	location, ok := g.myPlayer.CurrentLocation.Locations[locationName]
	if !ok {
		location, ok = g.myPlayer.CurrentLocation.Zones[locationName]
		if !ok {
			fmt.Println("Нет такой локации")
			return
		}
	}

	err := g.myPlayer.GoToLocation(location)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}

func (g *Game) TakeItem(itemName string) {
	itemName = strings.TrimSpace(itemName)
	item, ok := g.myPlayer.CurrentLocation.Items[itemName]
	if !ok {
		fmt.Println("Такого предмета нет в данной локации")
		return
	}
	g.myPlayer.TakeItem(item)
	delete(g.myPlayer.CurrentLocation.Items, itemName)
}

func (g *Game) UseItem(itemName string) {
	itemName = strings.TrimSpace(itemName)
	item, ok := g.myPlayer.Inventory.Items[itemName]
	if !ok {
		fmt.Println("В инвентаре нет данного предмета")
		return
	}

	item.Use("", g.myPlayer)
}
