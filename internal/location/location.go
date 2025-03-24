package location

import (
	"encoding/json"
	"fmt"
	"game/internal/item"
	"game/internal/npc"
	"strings"
)

type Location struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	IsOpen      bool        `json:"is_open"`
	Zones       []*Location `json:"zones"`
	Locations   []*Location `json:"-"`
	Monster     npc.Monster
	Items       []item.Item `json:"items"`
}

func (l *Location) UnmarshalJSON(data []byte) error {
	var temp struct {
		Name        string      `json:"name"`
		Description string      `json:"description"`
		IsOpen      bool        `json:"is_open"`
		Zones       []*Location `json:"zones"`
		Locations   []*Location `json:"-"`
		Monster     npc.Monster
		Items       json.RawMessage `json:"items"` // Используем RawMessage для отложенной десериализации
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	l.Name = temp.Name
	l.Description = temp.Description
	l.IsOpen = temp.IsOpen
	l.Zones = temp.Zones
	l.Locations = temp.Locations
	l.Monster = temp.Monster

	// Если items пуст, сразу назначаем пустой слайс
	if len(temp.Items) == 0 {
		l.Items = []item.Item{}
		return nil
	}

	// Десериализация Items
	var items []item.Item
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
			items = append(items, i)
		}
	}

	l.Items = items
	return nil
}

func NewLocation(name, description string, isOpen bool, zones []*Location, locations []*Location) *Location {
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
	zones := fmt.Sprintf("| Зоны: %-*s|", 30, l.getZonesName())           // 30 - максимальная длина
	locations := fmt.Sprintf("| Локации: %-*s|", 27, l.getLocationName()) // 30 - максимальная длина
	locationName := fmt.Sprintf("| Название: %-*s|", 26, l.Name)          // 30 - максимальная длина
	items := fmt.Sprintf("| Предметы: %-*s|", 26, l.getLocationItems())

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
	result += horizontalLine

	return result
}

func (l *Location) getZonesName() string {
	fmt.Println(l)
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
		items = append(items, item.GetName())
	}

	return strings.Join(items, ", ")
}
