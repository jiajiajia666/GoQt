package controller

import (
	"github.com/jiajiajia666/GoQt/core"

	"github.com/jiajiajia666/GoQt/internal/examples/showcases/wallet/theme/controller"
)

type colorController struct {
	core.QObject

	_ func() `signal:"change,auto"`
}

//lazy binding to the (qml singleton) theme controller
func (c *colorController) change() { controller.Controller.Change() }
