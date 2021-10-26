package app

import (
	"fmt"
	"l-iam/internal/api_server/config/options"
	"os"
	"strings"

	"github.com/spf13/viper"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type App struct {
	name    string
	desc    string
	version string
	runFunc RunFunc
	command *cobra.Command
	options *options.Options
}

type (
	Option  func(app *App)
	RunFunc func(name string) error
)

func WithDesc(desc string) Option {
	return func(a *App) {
		a.desc = desc
	}
}

func WithVersion(version string) Option {
	return func(app *App) {
		app.version = version
	}
}

func WithOptions(opt *options.Options) Option {
	return func(a *App) {
		a.options = opt
	}
}

func WithRunFunc(run RunFunc) Option {
	return func(a *App) {
		a.runFunc = run
	}
}

// NewApp
// 函数式选项模式，适用于参数比较复杂，有清晰默认值的场景，同时让参数可选更方便后续拓展
func NewApp(name string, opts ...Option) *App {
	app := &App{name: name}
	for _, opt := range opts {
		opt(app)
	}

	cmd := cobra.Command{
		// 命名使用方式
		Use: app.name,
		// 输入help后的短描述
		Short: app.name,
		// 输入help后的长描述
		Long:    app.desc,
		Version: app.version,
	}
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	// 针对flags排序
	cmd.Flags().SortFlags = true
	// 忽略. - _造成的区别
	cmd.Flags().SetNormalizeFunc(wordSepNormalizeFunc)

	if app.options != nil {
		for _, flag := range app.options.Flags() {
			cmd.Flags().AddFlagSet(flag)
		}
	}
	_ = viper.BindPFlags(cmd.PersistentFlags())
	addConfigFlags(app.name)

	if app.runFunc != nil {
		cmd.RunE = app.runCommand
	}

	app.command = &cmd
	return app
}

func (a *App) runCommand(cmd *cobra.Command, args []string) error {
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: failed to bind flags: %v\n", err)
		return err
	}
	if err := viper.Unmarshal(a.options); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: failed to unmarshal configuration file: %v\n", err)
		return err
	}
	return a.runFunc(a.name)
}

func (a *App) Run() {
	if err := a.command.Execute(); err != nil {
		fmt.Printf("%v %v \n", color.RedString("Error:"), err)
		os.Exit(1)
	}
}

func wordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.Replace(name, sep, to, -1)
	}
	return pflag.NormalizedName(name)
}
