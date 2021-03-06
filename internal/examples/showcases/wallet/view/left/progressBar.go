package left

import (
	"github.com/jiajiajia666/GoQt/quick"

	"github.com/jiajiajia666/GoQt/internal/examples/showcases/wallet/view/left/controller"
)

func init() { progressBarTemplate_QmlRegisterType2("LeftTemplate", 1, 0, "ProgressBarTemplate") }

type progressBarTemplate struct {
	quick.QQuickItem

	_ func() `constructor:"init"`

	_ string  `property:"text,<-(this.c)"`
	_ float64 `property:"value,<-(this.c)"`

	_ func() `signal:"clicked,->(this.c)"`

	c *controller.ProgressBarController
}

func (t *progressBarTemplate) init() {
	t.c = controller.NewProgressBarController(nil)
}
