package item

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	ErrOpenLocation = errors.New("Данный ключ не от этой локации")

	ErrUserDied = errors.New("Вы не смогли сразиться с монстром и погибли")
)

type Item interface {
	GetName() string
	GetType() string
	Use(locationName string) error
}

type DefaultItem struct {
	Name        string
	Description string
	Type        string
}

type player struct {
	Health int
	Damage int
}

func UnmarshalItem(data json.RawMessage) (Item, error) {
	var temp struct {
		Type string `json:"type"`
	}

	// Декодируем только type для определения типа предмета
	if err := json.Unmarshal(data, &temp); err != nil {
		return nil, err
	}

	var item Item
	switch temp.Type {
	case "key":
		item = &Key{}
	case "weapon":
		item = &Weapon{}
	case "health_potion":
		item = &HealthPotion{}
	case "damage_potion":
		item = &DamagePotion{}
	case "artifact":
		item = &Artifact{}
	default:
		return nil, fmt.Errorf("неизвестный тип предмета: %s", temp.Type)
	}

	// Декодируем полные данные предмета
	if err := json.Unmarshal(data, &item); err != nil {
		return nil, err
	}

	return item, nil
}
