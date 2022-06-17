package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

// InitLogger 初始化日志系统
func InitLogger(debug bool) *logrus.Logger {

	logger := logrus.New()

	// 设置日志格式为json格式
	logger.SetFormatter(&logrus.JSONFormatter{})

	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	logger.SetOutput(os.Stdout)

	// 正式发版设置日志级别为warn以上
	if !debug {
		logger.SetLevel(logrus.WarnLevel)
	} else {
		logger.SetLevel(logrus.TraceLevel)
		logrus.Println("\x1b[32mDEBUG MODE\x1b[0m")
	}
	return logger

}
