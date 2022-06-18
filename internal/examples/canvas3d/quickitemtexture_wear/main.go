//source: https://doc.qt.io/qt-5.11/qtcanvas3d-quickitemtexture-example.html

package main

import (
	"os"

	"github.com/jiajiajia666/GoQt/core"
	"github.com/jiajiajia666/GoQt/gui"
	"github.com/jiajiajia666/GoQt/qml"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	var app = gui.NewQGuiApplication(len(os.Args), os.Args)

	var engine = qml.NewQQmlApplicationEngine(nil)
	engine.Load(core.NewQUrl3("qrc:/main.qml", 0))

	app.Exec()
}
