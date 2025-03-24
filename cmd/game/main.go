package main

import (
	"fmt"
	"game/internal/command"
	"game/internal/game"
)

func main() {
    app := game.InitGame()
    commands := command.InitCommand(app)
	for {
		var inputCommand string
		fmt.Print("Введите команду: ")
		fmt.Scan(&inputCommand)

		if inputCommand != "" {
			isClose := commands.RunCommand(inputCommand)
			if isClose {
				break
			}
		}
	}

}
