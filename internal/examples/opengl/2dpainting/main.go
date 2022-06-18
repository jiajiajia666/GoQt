//source: https://doc.qt.io/qt-5/qtopengl-2dpainting-example.html

package main

import (
	"os"

	"github.com/jiajiajia666/GoQt/gui"
	"github.com/jiajiajia666/GoQt/widgets"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	fmt := gui.NewQSurfaceFormat()
	fmt.SetSamples(4)
	gui.QSurfaceFormat_SetDefaultFormat(fmt)

	window := NewWindow(nil, 0)
	window.Show()

	widgets.QApplication_Exec()
}
