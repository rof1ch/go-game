package main

import (
	"fmt"
	"game/internal/location"
	"game/internal/player"
)

func main() {
	chest := location.NewLocation("Сундук", "", true, nil, nil)
	location2 := location.NewLocation("Зал славы", "Зал славы", true, []location.Location{
		*chest,
	}, nil)
	location1 := location.NewLocation("Зал славы", "Зал славы", true, []location.Location{
		*chest,
	}, []location.Location{
		*location2,
	})

	player1 := player.Player{
		CurrentLocation: *location1,
	}

	fmt.Println(player1.CurrentLocation.GetLocationDraw())
}
