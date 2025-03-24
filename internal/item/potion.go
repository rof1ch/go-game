package item

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

func (p *HealthPotion) Use(locationName string) error {
	// Логика использования зелья для восстановления здоровья.
	// Предположим, что мы используем зелье только для игрока
	return nil
}

func (p *HealthPotion) GetName() string {
	return p.Name
}

func (p *HealthPotion) GetType() string {
	return p.Type
}

func (p *HealthPotion) UseForPlayer(player *player) {
	player.Health += p.Health
}

// DamagePotion теперь реализует интерфейс Item
type DamagePotion struct {
	DefaultItem
	Damage int
}

func (p *DamagePotion) Use(locationName string) error {
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

func (p *DamagePotion) GetType() string {
	return p.Type
}
