package item

type Key struct {
	DefaultItem
	LocationName string
}

func (k *Key) Use(locationName string) error {
	if locationName != k.LocationName {
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
