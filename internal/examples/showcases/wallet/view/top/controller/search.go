package controller

import (
	"github.com/jiajiajia666/GoQt/core"

	"github.com/jiajiajia666/GoQt/internal/examples/showcases/wallet/files/controller"
	lcontroller "github.com/jiajiajia666/GoQt/internal/examples/showcases/wallet/view/left/controller"
)

type searchController struct {
	core.QObject

	_ func(string) `signal:"search,auto"`
}

func (c *searchController) search(name string) {
	lcontroller.LeftController.Clicked("files")
	controller.FilesController.Model().Filter.SetFilterFixedString(name)
}
