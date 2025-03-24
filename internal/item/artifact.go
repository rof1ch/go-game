package item

type Artifact struct {
	DefaultItem
}

func (a *Artifact) Use(_ string) error {
	return nil
}

func (a *Artifact) GetName() string {
	return a.Name
}

func (a *Artifact) GetType() string {
	return a.Type
}
