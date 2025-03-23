package location

import (
	"fmt"
	"strings"
)

type Location struct {
	Name          string
	Description   string
	IsOpen        bool
	Zones         []Location
	NextLocations []Location
	Chest         interface{}
}

//TODO Написать стрктуру chest которая будет хранить также items и иметь функционал открыть, взять один, взять все, закрыть 

func NewLocation(name, description string, isOpen bool, zones []Location, locations []Location) *Location {
	return &Location{
		Name:          name,
		Description:   description,
		IsOpen:        isOpen,
		Zones:         zones,
		NextLocations: locations,
	}
}

func (l *Location) GetLocationDraw() string {
	// Формируем строки
	zones := fmt.Sprintf("| Зоны: %-*s|", 30, l.getZonesName())           // 30 - максимальная длина
	locations := fmt.Sprintf("| Локации: %-*s|", 27, l.getLocationName()) // 30 - максимальная длина
	locationName := fmt.Sprintf("| Название: %-*s|", 26, l.Name)          // 30 - максимальная длина

	// Находим максимальную длину для рамки
	maxLength := len([]rune(locationName)) // Одна из строк имеет длину 30

	// Создаем горизонтальную линию с учетом максимальной длины
	horizontalLine := "+" + strings.Repeat("-", maxLength-2) + "+\n"

	// Формируем результат
	result := horizontalLine
	result += locationName + "\n"
	result += zones + "\n"
	result += locations + "\n"
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
	for _, location := range l.NextLocations {
		locations = append(locations, location.Name)
	}
	return strings.Join(locations, ",")
}
