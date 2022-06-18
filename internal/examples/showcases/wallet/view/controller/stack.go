package controller

import "github.com/jiajiajia666/GoQt/core"

var StackController *stackController

type stackController struct {
	core.QObject

	_ func() `constructor:"init`

	_ func(string) `signal:"clicked"`
}

func (c *stackController) init() {
	StackController = c
}
