package main

import (
	"os"

	"github.com/jiajiajia666/GoQt/widgets"

	"github.com/jiajiajia666/GoQt/internal/examples/uitools/calculator/ui"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	ui.NewCalculatorForm(nil).Show()

	widgets.QApplication_Exec()
}
