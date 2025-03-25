package item

import "fmt"

type Weapon struct {
	DefaultItem
	Damage int
}

func (w *Weapon) GetType() string {
	return w.Type
}

func (w *Weapon) Use(locationName string, player player) error {
	player.UseWeapon(w)
	return nil
}

func (w *Weapon) GetName() string {
	return w.Name
}

func (w *Weapon) GetInfo() string {
	return fmt.Sprintf("Наносит %d урона\n", w.Damage)
}
