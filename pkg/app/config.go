package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var configFile string

const FlagConfigFile = "apiserver.config"

func init() {
	pflag.StringVarP(&configFile, FlagConfigFile, "c", configFile, "Read config from specified file")
}

func addConfigFlags(name string) {
	viper.AutomaticEnv()
	viper.SetEnvPrefix(strings.Replace(strings.ToUpper(name), "-", "_", -1))
	// 奇数参数为旧字符串，偶数参数为新字符串
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	cobra.OnInitialize(func() {
		if configFile != "" {
			// 使用标识指定的config file
			viper.SetConfigFile(configFile)
		} else {
			// 依次使用默认的config file
			viper.AddConfigPath(".")
			viper.AddConfigPath(filepath.Join("./configs", name))
			home, _ := os.UserHomeDir()
			viper.AddConfigPath(filepath.Join(home, name))

			viper.SetConfigName(name)
		}

		if err := viper.ReadInConfig(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error: failed to read configuration file(%s): %v\n", configFile, err)
			os.Exit(1)
		}
	})
}
