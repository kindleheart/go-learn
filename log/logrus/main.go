package main

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	//log.Out = os.Stdout

	//可以设置像文件等任意`io.Writer`类型作为日志输出
	//file, err := os.OpenFile("./logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//if err == nil {
	//	log.Out = file
	//} else {
	//	log.Info("Failed to log to file, using default stderr")
	//}
	log.SetLevel(logrus.InfoLevel)
	log.SetReportCaller(true)
	log.WithFields(logrus.Fields{
		"animal": "dog",
	}).Info("一条舔狗出现了。")

	//log.Trace("Something very low level.")
	//log.Debug("Useful debugging information.")
	//log.Info("Something noteworthy happened!")
	//log.Warn("You should probably take a look at this.")
	//log.Error("Something failed but I'm not quitting.")
	//// 记完日志后会调用os.Exit(1)
	//log.Fatal("Bye.")
	//// 记完日志后会调用 panic()
	//log.Panic("I'm bailing.")
	requestLogger := log.WithFields(logrus.Fields{"request_id": 1111, "user_ip": 2222})
	requestLogger.Info("something happened on that request") // will log request_id and user_ip
	requestLogger.Warn("something not great happened")

}
