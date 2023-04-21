package utils

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func GetDataFromFlagOrConf(cmd *cobra.Command) string {
	return GetFlagOrConfString(cmd, "data")
}

func GetFlagOrConfString(cmd *cobra.Command, field string) string {
	if fromFlag, err := cmd.Flags().GetString(field); err == nil {
		if len(fromFlag) != 0 {
			viper.Set(field, fromFlag)
		}
	}
	s := viper.GetString(field)
	return s
}

func GetFlagOrConfUint(cmd *cobra.Command, field string) uint64 {
	if fromFlag, err := cmd.Flags().GetUint64(field); err == nil {
		viper.Set(field, fromFlag)
	}
	s := viper.GetUint64(field)
	return s
}
