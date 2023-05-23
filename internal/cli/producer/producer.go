package producer

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ProducerCmd = &cobra.Command{
	Use:     "producer",
	Aliases: []string{"p"},
	Short:   "commands related to indexes",
	Run: func(cmd *cobra.Command, args []string) {
		port := viper.GetInt("producer.target.port")
		host := viper.GetString("producer.target.host")

		// TODO
		_ = port
		_ = host

	},
}

func init() {

	ListenCmd.PersistentFlags().Int("target.port", 10161, "port to listen to")
	viper.BindPFlag("producer.target.port", ListenCmd.PersistentFlags().Lookup("target.por"))

	ListenCmd.PersistentFlags().String("target.host", "127.0.0.1", "host to listen")
	viper.BindPFlag("producer.target.host", ListenCmd.PersistentFlags().Lookup("target.host"))
}
