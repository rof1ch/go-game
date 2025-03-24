package config

import (
	"encoding/json"
	"fmt"
	"game/internal/location"
	"os"
)

// tempLocation — временная структура для парсинга JSON
type tempLocation struct {
	Name        string               `json:"name"`
	Description string               `json:"description"`
	IsOpen      bool                 `json:"is_open"`
	Zones       []*location.Location `json:"zones"`
	Locations   []*location.Location `json:"locations"` // Здесь пока просто имена локаций
}

// LoadLocations загружает локации из JSON и возвращает корневую локацию
func LoadLocations(filename, rootName string) (*location.Location, error) {
	// Открываем JSON-файл
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer file.Close()

	// Парсим JSON в промежуточную структуру
	var tempLocations []tempLocation
	err = json.NewDecoder(file).Decode(&tempLocations)
	if err != nil {
		return nil, fmt.Errorf("ошибка декодирования JSON: %w", err)
	}

	// Создаем карту {имя → объект локации}
	locationMap := make(map[string]*location.Location)

	// Создаем объекты Location без ссылок
	for _, tempLoc := range tempLocations {
		loc := &location.Location{
			Name:        tempLoc.Name,
			Description: tempLoc.Description,
			Zones:       tempLoc.Zones,
			IsOpen:      tempLoc.IsOpen,
		}
		locationMap[tempLoc.Name] = loc
	}

	// Проставляем ссылки на связанные локации
	for _, tempLoc := range tempLocations {
		loc := locationMap[tempLoc.Name]
		for _, linkedName := range tempLoc.Locations {
			if linkedLoc, exists := locationMap[linkedName.Name]; exists {
				loc.Locations = append(loc.Locations, linkedLoc)
			} else {
				fmt.Printf("⚠️ Локация \"%s\" ссылается на несуществующую \"%s\"\n", tempLoc.Name, linkedName.Name)
			}
		}
	}

	// Возвращаем корневую локацию
	rootLoc, exists := locationMap[rootName]
	if !exists {
		return nil, fmt.Errorf("корневая локация \"%s\" не найдена", rootName)
	}
	return rootLoc, nil
}
