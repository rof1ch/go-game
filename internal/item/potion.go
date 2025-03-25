package item

import (
	"fmt"
	"strings"
)

// Potion теперь реализует интерфейс Item
type Potion interface {
	Item // Встраиваем интерфейс Item
}

// HealthPotion теперь реализует интерфейс Item
type HealthPotion struct {
	DefaultItem
	Health int
}

func (p *HealthPotion) Use(locationName string, player player) error {
	player.UpdateHealth(p.Health)
	player.DeleteInventoryItem(strings.TrimSpace(p.GetName()))
	fmt.Printf("Ваше здоровье повышенно на %d\n", p.Health)
	return nil
}

func (p *HealthPotion) GetName() string {
	return p.Name
}

func (p *HealthPotion) GetType() string {
	return p.Type
}

func (p *HealthPotion) GetInfo() string {
	return fmt.Sprintf("Добавлет игроку %d здоровья\n", p.Health)
}

// DamagePotion теперь реализует интерфейс Item
type DamagePotion struct {
	DefaultItem
	Damage int
}

func (p *DamagePotion) Use(locationName string, player player) error {
	player.UpdateDamage(p.Damage)
	player.DeleteInventoryItem(strings.TrimSpace(p.GetName()))
	fmt.Printf("Ваш урон повышен на %d\n", p.Damage)
	return nil
}

func (p *DamagePotion) GetName() string {
	return p.Name
}

func (p *DamagePotion) GetType() string {
	return p.Type
}

func (p *DamagePotion) GetInfo() string {
	return fmt.Sprintf("Добавляет игроку силы на %d\n", p.Damage)
}
