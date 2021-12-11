package log

import (
	"github.com/spf13/pflag"
	"go.uber.org/zap/zapcore"
)

const (
	flagName              = "log.name"
	flagLevel             = "log.level"
	flagDisableCaller     = "log.disable-caller"
	flagDisableStacktrace = "log.disable-stacktrace"
	flagFormat            = "log.format"
	flagEnableColor       = "log.enable-color"
	flagOutputPaths       = "log.output-paths"
	flagErrorOutputPaths  = "log.error-output-paths"
	flagDevelopment       = "log.development"
)

const (
	formatJson    = "json"
	formatConsole = "console"
)

type Options struct {
	Name             string   `json:"name"`
	Level            string   `json:"level"`
	Format           string   `json:"format"`
	OutputPaths      []string `json:"output_paths"`
	ErrorOutputPaths []string `json:"error_output_paths"`
	// EnableColor 是否允许颜色
	EnableColor bool `json:"enable_color"`
	Development bool `json:"development"`
	// DisableCaller是否禁用文件名+行号
	DisableCaller     bool `json:"disable_caller"`
	DisableStacktrace bool `json:"disable_stacktrace"`
}

func NewOptions() *Options {
	return &Options{
		Level:             zapcore.InfoLevel.String(),
		Format:            formatConsole,
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
		EnableColor:       true,
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
	}
}

func (o *Options) AddFlags() *pflag.FlagSet {
	fs := pflag.NewFlagSet("log", pflag.ExitOnError)

	fs.StringVar(&o.Level, flagLevel, o.Level, "Minimum log output `LEVEL`.")
	fs.BoolVar(&o.DisableCaller, flagDisableCaller, o.DisableCaller, "Disable output of caller information in the log.")
	fs.BoolVar(&o.DisableStacktrace, flagDisableStacktrace,
		o.DisableStacktrace, "Disable the log to record a stack trace for all messages at or above panic level.")
	fs.StringVar(&o.Format, flagFormat, o.Format, "Log output `FORMAT`, support plain or json format.")
	fs.BoolVar(&o.EnableColor, flagEnableColor, o.EnableColor, "Enable output ansi colors in plain format logs.")
	fs.StringSliceVar(&o.OutputPaths, flagOutputPaths, o.OutputPaths, "Output paths of log.")
	fs.StringSliceVar(&o.ErrorOutputPaths, flagErrorOutputPaths, o.ErrorOutputPaths, "Error output paths of log.")
	fs.BoolVar(
		&o.Development,
		flagDevelopment,
		o.Development,
		"Development puts the logger in development mode, which changes "+
			"the behavior of DPanicLevel and takes stacktraces more liberally.",
	)
	fs.StringVar(&o.Name, flagName, o.Name, "The name of the logger.")

	return fs
}
