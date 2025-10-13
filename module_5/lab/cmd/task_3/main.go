package main

import (
	"fmt"
	"lab/internal/prototype"
)

func main() {
	orc := &Monster{Type: "Orc", Health: 120, Damage: 25}
	orcClone := orc.Clone()
	orc2 := orc
	_ = orc2
	orc.GetInfo()
	orcClone.GetInfo()
}
