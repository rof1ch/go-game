package item

import (
	"errors"
	"game/internal/location"
)

var (
	ErrOpenLocation = errors.New("Данный ключ не от этой локации")

	ErrUserDied = errors.New("Вы не смогли сразиться с монстром и погибли")
)

type Item interface {
	GetName() string
	GetType() string
	Use(locat location.Location) error
}

type DefaultItem struct {
	Name        string
	Description string
	Type        string
}

// Key теперь реализует интерфейс Item
type Key struct {
	DefaultItem
	LocationName string
}

func (k *Key) Use(locat location.Location) error {
	if locat.Name != k.LocationName {
		return ErrOpenLocation
	}
	return nil
}

func (k *Key) GetName() string {
	return k.LocationName
}

func (k *Key) GetType() string {
	return k.Type
}

// Weapon теперь реализует интерфейс Item
type Weapon struct {
	DefaultItem
	Damage int
}

func (w *Weapon) Use(locat location.Location) error {
	// Для оружия логика использования будет изменена (может, например, атаковать монстра, но не здесь)
	// Это будет специфичная логика, но интерфейс нужно соблюсти.
	return nil
}

func (w *Weapon) GetName() string {
	return w.Name
}

type player struct {
	Health int
	Damage int
}

// Potion теперь реализует интерфейс Item
type Potion interface {
	Item // Встраиваем интерфейс Item
	UseForPlayer(player *player)
}

// HealthPotion теперь реализует интерфейс Item
type HealthPotion struct {
	DefaultItem
	Health int
}

func (p *HealthPotion) Use(locat location.Location) error {
	// Логика использования зелья для восстановления здоровья.
	// Предположим, что мы используем зелье только для игрока
	return nil
}

func (p *HealthPotion) GetName() string {
	return p.Name
}

func (p *HealthPotion) UseForPlayer(player *player) {
	player.Health += p.Health
}

// DamagePotion теперь реализует интерфейс Item
type DamagePotion struct {
	DefaultItem
	Damage int
}

func (p *DamagePotion) Use(locat location.Location) error {
	// Логика использования зелья для увеличения урона.
	// Также можно добавить логику для различных типов использования
	return nil
}

func (p *DamagePotion) UseForPlayer(player *player) {
	player.Damage += p.Damage
}

func (p *DamagePotion) GetName() string {
	return p.Name
}
