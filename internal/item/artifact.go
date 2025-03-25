package item

type Artifact struct {
	DefaultItem
}

func (a *Artifact) Use(locationName string, player player) error {
	return nil
}

func (a *Artifact) GetName() string {
	return a.Name
}

func (a *Artifact) GetType() string {
	return a.Type
}

func (a *Artifact) GetInfo() string {
	return "Просто редкий предмет"
}
