package player

import (
	"errors"
	"game/internal/item"
	"game/internal/location"
	"game/internal/npc"
)

type Player struct {
	Name            string
	CurrentLocation location.Location
	Inventory       Inventory
	Health          int
	Damage          int
	Weapon          item.Item
}

type Inventory struct {
	Items map[string]item.Item
}

func (p *Player) TakeItem(item item.Item) {
	p.Inventory.Items[item.GetName()] = item
}

func (p *Player) Attack(monster *npc.Monster) {
	monster.Health -= p.Damage
	p.Health -= monster.Damage
}

func (p *Player) UseWeapon(weapon item.Item) {
	p.Weapon = weapon
}

func (p *Player) GoToLocation(locat *location.Location) error {
	if !locat.IsOpen {
		for _, item := range p.Inventory.Items {
			if item.GetType() == "key" && item.GetName() == locat.Name {
				err := item.Use(*locat) // вызываем Use
				if err != nil {
					return err // если ошибка — возвращаем её сразу
				}
				locat.IsOpen = true // Открываем локацию
				break               // выходим из цикла, так как ключ найден и использован
			}
		}

		// Если локация осталась закрытой после попытки использовать ключ
		if !locat.IsOpen {
			return errors.New("локация закрыта, ключ не найден")
		}
	}

	// Переход в новую локацию
	p.CurrentLocation = *locat
	return nil
}
