package log

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

//ArkLogger log Facade Object
type ArkLogger struct {
	*logrus.Logger
}

//Log 系统Log全局变量
var defaultLog *ArkLogger

func init() {
	defaultLog = createLogger()
}

//new create new log for global use
func createLogger() *ArkLogger {
	if defaultLog != nil {
		return defaultLog
	}

	Log := logrus.New()

	path := "logs/ark.log"
	writer, _ := rotatelogs.New(
		path+".%Y%m%d",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(604800)*time.Second),
	)

	Log.Hooks.Add(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  writer,
			logrus.ErrorLevel: writer,
		},
		&logrus.JSONFormatter{},
	))

	// log := logrus.New()
	// //设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	// logrus.SetFormatter(&logrus.TextFormatter{})

	// //设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
	// logFile, err := os.OpenFile("logs/ark.log", os.O_CREATE|os.O_WRONLY, 0666)
	// if err == nil {
	// 	mw := io.MultiWriter(os.Stdout, logFile)
	// 	logrus.SetOutput(mw)
	// } else {
	// 	fmt.Println("failed to create log file")
	// 	os.Exit(-1)
	// }
	// defer logFile.Close()

	// //设置最低loglevel
	// logrus.SetLevel(logrus.InfoLevel)

	return &ArkLogger{Log}
}

//Error output error log message
func Error(v ...interface{}) {
	defaultLog.Error(v)
}

//Info output Info log message
func Info(v ...interface{}) {
	defaultLog.Info(v)
}

//Debug output Debug log message
func Debug(v ...interface{}) {
	defaultLog.Debug(v)
}

//TODO: log level in config file
//TODO: console appender can be disable in config file
