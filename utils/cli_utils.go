package utils

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func GetDataFromFlagOrConf(cmd *cobra.Command) string {
	return GetFlagOrConf(cmd, "data")
}

func GetFlagOrConf(cmd *cobra.Command, field string) string {
	if fromFlag, err := cmd.Flags().GetString(field); err == nil {
		if len(fromFlag) != 0 {
			viper.Set(field, fromFlag)
		}
	}
	s := viper.GetString(field)
	return s
}
