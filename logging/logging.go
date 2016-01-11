package logging

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/evalphobia/logrus_fluent"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"time"
)

var (
	debugFlag     bool
	fluentdServer string
)

func Connect() bool {

	success := false

	logrus.SetFormatter(&logrus.JSONFormatter{})

	if !debugFlag {
		logrus.SetOutput(ioutil.Discard)
	} else {
		logrus.SetOutput(os.Stdout)
	}
	if fluentdServer != "" {
		host, port, err := net.SplitHostPort(fluentdServer)
		if err != nil {
			LogError(fmt.Sprintf("Unable to split to fluentd server [%s]!\n", fluentdServer), err.Error())
		}
		portInt, err := strconv.Atoi(port)

		if err != nil {
			LogError(fmt.Sprintf("Unable to convert port format [%s]!\n", port), err.Error())

		}
		hook := logrus_fluent.NewHook(host, portInt)
		LogStd(fmt.Sprintf("Received hook to fluentd server [%s]!\n", fluentdServer), false)
		logrus.AddHook(hook)
		success = true
	}

	return success
}

func SetupLogging(fluentdSvr string, debug bool) {
	debugFlag = debug
	fluentdServer = fluentdSvr
}

func LogStd(message string, force bool) {
	Log(message, force, false, nil)
}

func LogError(message string, errMsg interface{}) {

	Log(message, false, true, errMsg)
}

func Log(message string, force bool, isError bool, err interface{}) {

	if debugFlag || force || isError {

		writer := os.Stdout
		var formattedMessage string

		if isError {
			writer = os.Stderr
			formattedMessage = fmt.Sprintf("[%s] Exception occurred! Message: %s Details: %v", time.Now().String(), message, err)
		} else {
			formattedMessage = fmt.Sprintf("[%s] %s", time.Now().String(), message)
		}

		fmt.Fprintln(writer, formattedMessage)
	}
}
