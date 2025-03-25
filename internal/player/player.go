package player

import (
	"errors"
	"fmt"
	"game/internal/item"
	"game/internal/location"
	"game/internal/npc"
	"strings"
)

type Player struct {
	Name            string
	CurrentLocation *location.Location
	Inventory       Inventory
	Health          int
	Damage          int
	Weapon          *item.Weapon
}

func NewPlayer(name string, location *location.Location) *Player {
	inventory := Inventory{
		Items: make(map[string]item.Item),
	}
	return &Player{
		Name:            name,
		Inventory:       inventory,
		CurrentLocation: location,
		Health:          100,
		Damage:          9,
	}
}

type Inventory struct {
	Items map[string]item.Item
}

func (i *Inventory) GetItems() string {
	if len(i.Items) == 0 {
		return "Ваш инвентарь пуст"
	}
	var output string
	for _, item := range i.Items {
		output += fmt.Sprintf("%s - %s\n", item.GetName(), item.GetInfo())
	}
	return output
}

func (p *Player) TakeItem(item item.Item) {
	p.Inventory.Items[item.GetName()] = item
}

func (p *Player) Attack(monster *npc.Monster) {
	if p.Weapon != nil {
		monster.Health -= p.Damage + p.Weapon.Damage
	} else {
		monster.Health -= p.Damage
	}
	p.UpdateHealth(-monster.Damage)
}

func (p *Player) UseWeapon(weapon *item.Weapon) {
	p.Weapon = weapon
}

func (p *Player) GoToLocation(locat *location.Location) error {
	if !locat.IsOpen {
		for _, item := range p.Inventory.Items {
			if item.GetType() == "key" {
				err := item.Use(locat.Name, nil) // вызываем Use
				if err != nil {
					return err // если ошибка — возвращаем её сразу
				}
				locat.IsOpen = true // Открываем локацию
				delete(p.Inventory.Items, strings.TrimSpace(item.GetName()))
				break // выходим из цикла, так как ключ найден и использован
			}
		}

		// Если локация осталась закрытой после попытки использовать ключ
		if !locat.IsOpen {
			return errors.New("локация закрыта, ключ не найден")
		}
	}

	// Переход в новую локацию
	p.CurrentLocation = locat
	fmt.Printf("Вы перешли в локацию %s\n", locat.Name)
	return nil
}

func (p *Player) UpdateHealth(health int) {
	p.Health += health
}

func (p *Player) UpdateDamage(damage int) {
	p.Damage += damage
}

func (p *Player) DeleteInventoryItem(itemName string) {
	delete(p.Inventory.Items, itemName)
}
