package log

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	m sync.Mutex
	l = New(NewOptions())
)

const (
	KeyUsername   = "username"
	KeyXRequestID = "X-Request-ID"
)

func Init(opts *Options) {
	m.Lock()
	defer m.Unlock()
	l = New(opts)
}

func New(opts *Options) *zap.Logger {
	if opts == nil {
		opts = NewOptions()
	}
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}

	encodeLevel := zapcore.CapitalLevelEncoder
	// 如果是文本格式且允许颜色的话，则按级别显示颜色
	if opts.Format == formatConsole && opts.EnableColor {
		encodeLevel = zapcore.CapitalColorLevelEncoder
	}

	loggerConfig := &zap.Config{
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Development:       opts.Development,
		DisableCaller:     opts.DisableCaller,
		DisableStacktrace: opts.DisableStacktrace,
		// 为了避免因日志造成CPU和I/O负载，同样级别的日志在每秒内如果输出Initial条，则丢弃后续Thereafter条
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: opts.Format,
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:       "ts",
			LevelKey:      "level",
			NameKey:       "logger",
			CallerKey:     "caller",
			MessageKey:    "message",
			StacktraceKey: "stacktrace",
			EncodeLevel:   encodeLevel,
			LineEnding:    zapcore.DefaultLineEnding,
			// time的默认展示为时间戳，此处将其转换为对人友好的时间格式
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
			},
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
		OutputPaths:      opts.OutputPaths,
		ErrorOutputPaths: opts.ErrorOutputPaths,
	}

	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}

	zap.ReplaceGlobals(logger)

	return logger
}

func Debug(msg string, fields ...zapcore.Field) {
	l.Debug(msg, fields...)
}

func Info(msg string, fields ...zapcore.Field) {
	l.Info(msg, fields...)
}

func Warn(msg string, fields ...zapcore.Field) {
	l.Warn(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	l.Error(msg, fields...)
}

func Panic(msg string, fields ...zapcore.Field) {
	l.Panic(msg, fields...)
}

func Flush() {
	_ = l.Sync()
}

func L(ctx context.Context) *zap.Logger {
	requestId, _ := ctx.Value(KeyXRequestID).(string)

	return l.With(zap.String(KeyXRequestID, requestId))
}
