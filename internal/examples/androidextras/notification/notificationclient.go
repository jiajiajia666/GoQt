package main

import (
	"github.com/jiajiajia666/GoQt/androidextras"
	"github.com/jiajiajia666/GoQt/core"
)

type NotificationClient struct {
	core.QObject

	_ string `property:"notification,auto,changed"`
}

func (c *NotificationClient) notificationChanged(n string) {

	var err = androidextras.QAndroidJniObject_CallStaticMethodVoid2Caught(
		"org/qtproject/example/notification/NotificationClient",
		"notify",
		"(Ljava/lang/String;)V",
		n,
	)

	if err != nil {
		println(err.Error())
	}
}
