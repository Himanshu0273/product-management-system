package logging

import (
    "github.com/sirupsen/logrus"
)

var log = logrus.New()

func InitLogging() {
    log.SetFormatter(&logrus.JSONFormatter{})
    log.SetReportCaller(true)
}

func LogInfo(message string) {
    log.Info(message)
}

func LogError(message string, err error) {
    log.Errorf("%s: %v", message, err)
}
