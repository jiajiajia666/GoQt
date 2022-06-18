//source: https://doc.qt.io/qt-5.11/qtcanvas3d-threejs-oneqt-example.html

package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/jiajiajia666/GoQt/core"
	"github.com/jiajiajia666/GoQt/gui"
	"github.com/jiajiajia666/GoQt/quick"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	var app = gui.NewQGuiApplication(len(os.Args), os.Args)

	var viewer = quick.NewQQuickView(nil)

	var extraImportPath string
	if runtime.GOOS == "windows" {
		extraImportPath = "%v/../../../../%v"
	} else {
		extraImportPath = "%v/../../../%v"
	}

	viewer.Engine().AddImportPath(fmt.Sprintf(extraImportPath, app.ApplicationDirPath(), "qml"))
	viewer.SetSource(core.NewQUrl3("qrc:/oneqt.qml", 0))

	viewer.SetTitle("One Qt")
	viewer.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	viewer.SetColor(gui.NewQColor6("#FCFCFC"))
	viewer.Show()

	app.Exec()
}