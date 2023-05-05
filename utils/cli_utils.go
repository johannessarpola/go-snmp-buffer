package utils

import (
	"fmt"
	"time"

	g "github.com/gosnmp/gosnmp"
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

func GetFlagOrConfInt(cmd *cobra.Command, field string) int {
	if fromFlag, err := cmd.Flags().GetInt(field); err == nil {
		viper.Set(field, fromFlag)
	}
	s := viper.GetInt(field)
	return s
}

// TODO Rework to be smarter like translate variable types to string so remove fmt.Prints
func PrintVars(variable g.SnmpPDU) {
	fmt.Printf("oid: %s ", variable.Name)

	switch variable.Type {
	case g.OctetString:
		bytes := variable.Value.([]byte)
		fmt.Printf("string: %s\n", string(bytes))
	case g.TimeTicks:
		n := g.ToBigInt(variable.Value)
		tm := time.Unix(n.Int64(), 0)
		fmt.Printf("time: %s\n", tm.String())
	default:
		// ... or often you're just interested in numeric values.
		// ToBigInt() will return the Value as a BigInt, for plugging
		// into your calculations.
		fmt.Printf("number: %d\n", g.ToBigInt(variable.Value))
	}
}

func PrettyPrintSnmpKey(b []byte, prefix_len int) string {
	prefix := string(b[:prefix_len])
	idx_arr := b[prefix_len:]
	idx := ConvertToUint64(idx_arr)
	pretty_k := fmt.Sprintf("%s%d", prefix, idx)
	return pretty_k
}
