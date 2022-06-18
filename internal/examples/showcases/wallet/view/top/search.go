package top

import (
	"github.com/jiajiajia666/GoQt/quick"

	_ "github.com/jiajiajia666/GoQt/internal/examples/showcases/wallet/view/top/controller"
)

func init() { searchTemplate_QmlRegisterType2("TopTemplate", 1, 0, "SearchTemplate") }

type searchTemplate struct {
	quick.QQuickItem

	_ func(string) `signal:"search,->(controller.NewSearchController(nil))"`
}
