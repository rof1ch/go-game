package item

type Key struct {
	DefaultItem
	LocationName string `json:"location_name"`
}

func (k *Key) Use(locationName string) error {
	if locationName != k.LocationName {
		return ErrOpenLocation
	}
	return nil
}

func (k *Key) GetName() string {
	return k.Name
}

func (k *Key) GetType() string {
	return k.Type
}
