package command

import (
	"bufio"
	"fmt"
	"game/internal/colors"
	"game/internal/game"
	"os"
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
        | ` + getCommand("go") + getText("Перейти в какую-либо локацию") + `                              |
        | ` + getCommand("atack") + getText("Атаковать монстра ") + `                                     |
        | ` + getCommand("take") + getText("Подобрать предмет") + `                                       |
        | ` + getCommand("inventory") + getText("Открыть инвентарь") + `                                  |
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
		fmt.Print(colors.GetCyanText("Введите название локации/зоны: "))
		fmt.Scan(&locationName)
		c.game.GoToLocation(locationName)
	case "take":
		fmt.Printf(colors.GetCyanText("Введите название предмета: "))
		itemName, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		c.game.TakeItem(itemName)
	case "use":
		fmt.Printf(colors.GetCyanText("Введите название предмета: "))
		itemName, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		c.game.UseItem(itemName)
	case "atack":
		var monsterName string
		fmt.Print(colors.GetCyanText("Введите название монстра: "))
		fmt.Scan(&monsterName)
		if err := c.game.Atack(monsterName); err != nil {
			fmt.Println(colors.GetRedText(err.Error()))
			return true
		}
	default:
		fmt.Println(colors.GetRedText("Комманда неизвестна"))
	}

	return false
}

func getCommand(str string) string {
	return Blue + str + Reset
}

func getText(str string) string {
	return " - " + Green + str + Reset
}
