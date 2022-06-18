package dialog

import (
	"github.com/jiajiajia666/GoQt/core"

	fcontroller "github.com/jiajiajia666/GoQt/internal/examples/showcases/wallet/files/dialog/controller"
	wcontroller "github.com/jiajiajia666/GoQt/internal/examples/showcases/wallet/wallet/dialog/controller"
)

type filesDialogTemplate struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(cident string) `signal:"show,<-(fcontroller.Controller)"`
	_ func(bool)          `signal:"blur,->(fcontroller.Controller)"`
}

func (t *filesDialogTemplate) init() {
	if fcontroller.Controller == nil {
		fcontroller.NewDialogController(nil)
	}
}

func (t *filesDialogTemplate) show(cident string) {
	if fcontroller.Controller.IsLocked() {
		wcontroller.Controller.Show("unlock")
	} else {
		t.Show(cident)
	}
}
