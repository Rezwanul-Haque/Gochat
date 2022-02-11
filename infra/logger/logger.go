package logger

import (
	"encoding/json"
	"fmt"
	"gochat/infra/config"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	LVLINFO  = "INFO"
	LVLERROR = "ERROR"
)

var client *zap.Logger

func NewLoggerClient() {
	if config.App().LoggerClientType == "zap" {
		initializeZap(config.App().LogLevel)
	}
}

func initializeZap(level string) {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		MessageKey:     "message",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     "\n",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}

	config := zap.Config{
		Level:            stringToLevel(level),
		Encoding:         "json",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
	}
	var err error

	client, err = config.Build()
	if err != nil {
		panic(err)
	}
	defer client.Sync()
}

func stringToLevel(str string) zap.AtomicLevel {
	str = strings.ToLower(str)
	switch str {
	case "debug":
		return zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case "error":
		return zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	case "fatal":
		return zap.NewAtomicLevelAt(zapcore.FatalLevel)
	case "warn":
		return zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case "info":
		return zap.NewAtomicLevel()
	default:
		return zap.NewAtomicLevelAt(zapcore.PanicLevel)
	}
}

func Debug(msg string, data interface{}, tags ...zap.Field) {
	tags = append(tags, zap.Any("data", data))
	client.Debug(msg, tags...)
	_ = client.Sync()
}

func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	client.Error(msg, tags...)
	_ = client.Sync()
}

func Info(msg string) {
	client.Info(msg)
	_ = client.Sync()
}

func Warn(msg string) {
	client.Warn(msg)
	_ = client.Sync()
}

func Fatal(msg string) {
	client.Fatal(msg)
	_ = client.Sync()
}

func Panic(msg string) {
	client.Panic(msg)
	_ = client.Sync()
}

func InfoAsJson(msg string, data ...interface{}) {
	prettyPrint(LVLINFO, msg, data)
}

func ErrorAsJson(msg string, data ...interface{}) {
	prettyPrint(LVLERROR, msg, data)
}

func prettyPrint(level string, msg string, data ...interface{}) {
	if r, err := json.MarshalIndent(&data, "", "  "); err == nil {
		fmt.Printf("[%v] %v %v: \n %v\n", level, time.Now(), msg, string(r))
	}
}
