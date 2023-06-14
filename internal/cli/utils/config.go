package utils

import (
	"fmt"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func InitConfig(envprefix string, cfgFilename string, subcommands func() []*cobra.Command) {

	// listen.port -> <envprefix>_LISTEN_PORT
	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvPrefix(envprefix)
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	LoadConfig(cfgFilename, viper.GetString("configs"))
	for _, cmd := range subcommands() {
		MergeConfig(cfgFilename, cmd.Use)
	}
}

func LoadConfig(cfgFile string, cfgDir string) {
	// Find home directory if no directory defined
	if len(cfgDir) == 0 {
		home, err := homedir.Dir()
		cfgDir = home
		if err != nil {
			logrus.Fatal("Could not open homedir")
		}
	}

	viper.SetConfigName(fmt.Sprintf("%s.yaml", cfgFile))

	viper.AddConfigPath(cfgDir)
	viper.SetConfigType("yaml")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		logrus.Infof("Using config file: %s", viper.ConfigFileUsed())
	}

}

func MergeConfig(cfgFile string, suffix string) {
	// Load either <cfgFile>.<suffix>.yaml or <cfgFile>.yaml
	if len(suffix) > 0 {
		cm := fmt.Sprintf("%s.%s.yaml", cfgFile, suffix)
		viper.SetConfigName(cm)
		if err := viper.MergeInConfig(); err == nil {
			logrus.Infof("Using config file: %s", viper.ConfigFileUsed())
		}
	}
}
