package npc

import "fmt"

type Monster struct {
	Name   string
	Health int
	Damage int
}

type Npc struct {
	Name string
	Text string
}

func (n *Npc) Talk() {
	fmt.Printf("%s говорит: %s\n", n.Name, n.Text)
}
