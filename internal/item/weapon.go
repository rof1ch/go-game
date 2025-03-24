package item

type Weapon struct {
	DefaultItem
	Damage int
}

func (w *Weapon) GetType() string {
	return w.Type
}

func (w *Weapon) Use(locationName string) error {
	// Для оружия логика использования будет изменена (может, например, атаковать монстра, но не здесь)
	// Это будет специфичная логика, но интерфейс нужно соблюсти.
	return nil
}

func (w *Weapon) GetName() string {
	return w.Name
}
