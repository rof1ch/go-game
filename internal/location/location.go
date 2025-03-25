package location

import (
	"encoding/json"
	"fmt"
	"game/internal/item"
	"game/internal/npc"
	"strings"
)

type Location struct {
	Name        string               `json:"name"`
	Description string               `json:"description"`
	IsOpen      bool                 `json:"is_open"`
	Zones       map[string]*Location `json:"zones"`
	Locations   map[string]*Location `json:"-"`
	Monster     *npc.Monster
	Items       map[string]item.Item `json:"items"`
	Npc         *npc.Npc             `json:"npc"`
}

func (l *Location) UnmarshalJSON(data []byte) error {
	var temp struct {
		Name        string      `json:"name"`
		Description string      `json:"description"`
		IsOpen      bool        `json:"is_open"`
		Zones       []*Location `json:"zones"`
		Locations   []*Location `json:"-"`
		Monster     *npc.Monster
		Items       json.RawMessage `json:"items"` // Используем RawMessage для отложенной десериализации
		Npc         *npc.Npc        `json:"npc"`
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	l.Name = temp.Name
	l.Description = temp.Description
	l.IsOpen = temp.IsOpen
	for _, zone := range temp.Zones {
		l.Zones[zone.Name] = zone

	}
	for _, location := range temp.Locations {
		l.Locations[location.Name] = location
	}
	l.Monster = temp.Monster
	l.Npc = temp.Npc

	// Если items пуст, сразу назначаем пустой слайс
	if len(temp.Items) == 0 {
		l.Items = map[string]item.Item{}
		return nil
	}

	// Десериализация Items
	items := make(map[string]item.Item)
	// Парсим каждый элемент в Items через кастомную функцию
	if err := json.Unmarshal(temp.Items, &items); err != nil {
		// Пробуем по одному элементу в массиве
		var tempItems []json.RawMessage
		if err := json.Unmarshal(temp.Items, &tempItems); err != nil {
			return fmt.Errorf("ошибка при десериализации items: %w", err)
		}

		for _, itemData := range tempItems {
			i, err := item.UnmarshalItem(itemData) // Используем кастомный Unmarshal
			if err != nil {
				return fmt.Errorf("ошибка при десериализации предмета: %w", err)
			}
			if i != nil {
				items[strings.TrimSpace(i.GetName())] = i
			}
		}
	}

	l.Items = items

	return nil
}

func NewLocation(name, description string, isOpen bool, zones map[string]*Location, locations map[string]*Location) *Location {
	return &Location{
		Name:        name,
		Description: description,
		IsOpen:      isOpen,
		Zones:       zones,
		Locations:   locations,
	}
}

func (l *Location) GetLocationDraw() string {
	// Формируем строки
	zones := fmt.Sprintf("| Зоны: %-*s|", 40, l.getZonesName())           // 30 - максимальная длина
	locations := fmt.Sprintf("| Локации: %-*s|", 37, l.getLocationName()) // 30 - максимальная длина
	locationName := fmt.Sprintf("| Название: %-*s|", 36, l.Name)          // 30 - максимальная длина
	items := fmt.Sprintf("| Предметы: %-*s|", 36, l.getLocationItems())
	montsters := fmt.Sprintf("| Монстры: %-*s|", 37, l.getMonsterInfo())

	// Находим максимальную длину для рамки
	maxLength := len([]rune(locationName)) // Одна из строк имеет длину 30

	// Создаем горизонтальную линию с учетом максимальной длины
	horizontalLine := "+" + strings.Repeat("-", maxLength-2) + "+\n"

	// Формируем результат
	result := horizontalLine
	result += locationName + "\n"
	result += zones + "\n"
	result += locations + "\n"
	result += items + "\n"
	result += montsters + "\n"
	result += horizontalLine

	return result
}

func (l *Location) getZonesName() string {
	var zones []string
	for _, zone := range l.Zones {
		zones = append(zones, zone.Name)
	}
	return strings.Join(zones, ",")
}

func (l *Location) getLocationName() string {
	var locations []string
	for _, location := range l.Locations {
		locations = append(locations, location.Name)
	}
	return strings.Join(locations, ", ")
}

func (l *Location) getLocationItems() string {
	var items []string
	for _, item := range l.Items {
		if item != nil {
			items = append(items, item.GetName())
		}
	}

	return strings.Join(items, ", ")
}

func (l *Location) getMonsterInfo() string {
	if l.Monster == nil {
		return "В данной локации нет монстров"
	}

	return fmt.Sprintf("Имя: %s; Здоровье: %d; Урон: %d", l.Monster.Name, l.Monster.Health, l.Monster.Damage)
}
