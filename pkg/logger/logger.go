/*
  @Author : Mustang Kong
*/

package logger

import (
	"fmt"
	"golang-base-flamego/pkg/settings"
	"os"
	"time"

	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// error logger
var log *zap.SugaredLogger

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

// Filename: 日志文件的位置
// MaxSize：在进行切割之前，日志文件的最大大小（以MB为单位）
// MaxBackups：保留旧文件的最大个数
// MaxAges：保留旧文件的最大天数
// Compress：是否压缩/归档旧文件
func Initial() {
	logPath := fmt.Sprintf("%s%s", settings.ObjectPath(), "/log")
	_, err := os.Stat(logPath)
	if err != nil {
		err = os.Mkdir(logPath, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		}
		//logPath = viper.GetString(`log.path`)
	}
	//fmt.Println(logPath)

	lv := viper.GetString(`log.level`)
	level := getLoggerLevel(lv)
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.AddCallerSkip(1)
	if lv == "debug" {
		// 写入到console
		consoleDebugging := zapcore.Lock(os.Stdout)
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core := zapcore.NewCore(consoleEncoder, consoleDebugging, zap.NewAtomicLevelAt(level))

		logger := zap.New(core, caller, development)
		log = logger.Sugar()
	} else {
		path := fmt.Sprintf("%s/%s", logPath, viper.GetString(`log.fileName`))
		fmt.Println(path)

		syncWriter := zapcore.AddSync(&lumberjack.Logger{
			Filename:   path,                           // 日志文件名
			MaxSize:    viper.GetInt(`log.maxsize`),    // 日志文件大小
			MaxAge:     viper.GetInt(`log.maxage`),     // 最长保存天数
			MaxBackups: viper.GetInt(`log.maxbackups`), // 最多备份几个
			LocalTime:  viper.GetBool(`log.localtime`), // 日志时间戳 是否使用本地时间，默认使用UTC时间
			Compress:   viper.GetBool(`log.compress`),  // 是否压缩文件，使用gzip
		})
		encoder := zap.NewProductionEncoderConfig()
		//encoder.EncodeTime = zapcore.ISO8601TimeEncoder
		encoder.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05.000000"))
		}
		//encoder.EncodeLevel = zapcore.CapitalLevelEncoder

		core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(level))
		logger := zap.New(core, caller, development)
		log = logger.Sugar()
	}
	Info(logPath)
	Info("init logs success")
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func DPanic(args ...interface{}) {
	log.DPanic(args...)
}

func DPanicf(format string, args ...interface{}) {
	log.DPanicf(format, args...)
}

func Panic(args ...interface{}) {
	log.Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func GetLogger() *zap.Logger {
	return log.Desugar()
}

// ==============================================
