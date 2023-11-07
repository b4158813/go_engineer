package db

import (
	"my_test/cn"
	"os"

	"github.com/sirupsen/logrus"
)

var L *logrus.Entry = nil

func InitLog() {
	tempL := logrus.New()
	tempL.SetOutput(os.Stdout)
	tempL.SetFormatter(&logrus.JSONFormatter{})
	L = tempL.WithFields(logrus.Fields{
		"prdline": cn.PrdLine,
		"app":     cn.Schema + cn.Name,
	})
	L.Infof("%s-%s log start!", cn.Schema, cn.Name)
}
