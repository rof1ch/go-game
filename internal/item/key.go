package item

import "fmt"

type Key struct {
	DefaultItem
	LocationName string `json:"location_name"`
}

func (k *Key) Use(locationName string, _ player) error {
	if locationName != k.LocationName {
		return ErrOpenLocation
	}
	fmt.Printf("Вы успешно открыли %s\n", locationName)
	return nil
}

func (k *Key) GetName() string {
	return k.Name
}

func (k *Key) GetType() string {
	return k.Type
}

func (k *Key) GetInfo() string {
	return fmt.Sprintf("Открывает %s\n", k.LocationName)
}
