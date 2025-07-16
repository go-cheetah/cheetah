package log

import (
	"fmt"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Client struct {
	Logger *zap.Logger
}

func NewClient(infoFilePath, errorFilePath, level string,
	maxSize, maxBackups, maxAge int,
) *Client {
	writeSyncer := getLogWriter(infoFilePath,
		maxSize,
		maxBackups,
		maxAge)
	// error单独使用一个日志文件，平时好排查
	writeSyncer2 := getLogWriter(infoFilePath,
		maxSize,
		maxBackups,
		maxAge)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	var l2 = new(zapcore.Level)
	err := l.UnmarshalText([]byte(level))
	if err != nil {
		fmt.Println(err)
	}
	err = l2.UnmarshalText([]byte("ERROR"))
	if err != nil {
		fmt.Println(err)
	}
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, writeSyncer, l),
		zapcore.NewCore(encoder, writeSyncer2, l2),
	)

	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()和zap.S()调用即可
	return &Client{
		Logger: logger,
	}
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}
