package command

import (
	"fmt"
	"game/internal/colors"
	"game/internal/game"
	"strings"
)

var (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"
)

type Command struct {
	game *game.Game
}

func InitCommand(game *game.Game) *Command {
	return &Command{game: game}
}

func (c *Command) RunCommand(command string) bool {

	switch command {
	case "help":
		output := fmt.Sprintln(Magenta + `
        +----------------------------------------------------------------+` + Reset + `
        | ` + getCommand("help") + getText("Вызов этого меню") + `                                        |
        | ` + getCommand("quit / exit") + getText("Закрыть меню") + `                                     |
        | ` + getCommand("go <Название локации либо зоны>") + getText("Перейти в какую-либо локацию") + ` |
        | ` + getCommand("atack <Навазние монстра>") + getText("Атаковать монстра ") + `                  |
        | ` + getCommand("take <Навзание предмета>") + getText("Подобрать предмет") + `                   |
        | ` + getCommand("inventory") + getText("Открыть инвентарь") + `                                  |
        | ` + getCommand("talk <Название NPC>") + getText("Поговорить с NPC") + `                         |
        | ` + getCommand("location") + getText("Выведет информацию о текущей локации") + `                |
        | ` + getCommand("me") + getText("Выведет информацию о игроке") + `                               |
        ` + Magenta + `+----------------------------------------------------------------+` + Reset + `
        `)
		output = strings.ReplaceAll(output, "|", Magenta+"|"+Reset)
		fmt.Println(output)
	case "quit":
		fmt.Println("Игра закрыта")
		return true
	case "exit":
		fmt.Println("Игра закрыта")
		return true
	case "location":
		c.game.GetCurrentLocation()
	case "inventory":
		c.game.GetInventory()
	case "me":
		c.game.GetPlayerInfo()
	case "go":
		var locationName string
		fmt.Print(colors.GetCyanText("Введите название локации: "))
		fmt.Scan(&locationName)
		c.game.GoToLocation(locationName)
	default:
		fmt.Println("Комманда неизвестна")
	}

	return false
}

func getCommand(str string) string {
	return Blue + str + Reset
}

func getText(str string) string {
	return " - " + Green + str + Reset
}
