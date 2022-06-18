package controller

import (
	"github.com/jiajiajia666/GoQt/core"

	"github.com/jiajiajia666/GoQt/internal/examples/showcases/wallet/view/controller"
)

type buttonController struct {
	core.QObject

	_ func(code string) `signal:"clicked,auto"`
}

//lazy binding to the view/stack controller
func (c *buttonController) clicked(code string) { controller.StackController.Clicked(code) }
