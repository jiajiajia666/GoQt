package main

import (
	"github.com/jiajiajia666/GoQt/core"
)

type Person struct {
	core.QObject

	_ string `property:"name"`
	_ int    `property:"shoeSize"`
}
