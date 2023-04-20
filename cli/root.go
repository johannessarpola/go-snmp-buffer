package cli

import (
	"fmt"
	"os"

	i "github.com/johannessarpola/go-network-buffer/cli/index"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "admin",
	Short: "admin - CLI to view the database",
	Long: `description
stuff
stuff`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("Started cli")
	// },
}

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // type of config file
	viper.AddConfigPath(".")      // path to look for the config file in the working directory
	// Set default values for the configuration parameters
	//	viper.SetDefault("server.host", "localhost")
	//	viper.SetDefault("server.port", 9999)

	// Read the configuration file (if present) and override any defaults with the file values
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Using default values")
	} else {
		fmt.Printf("Using config file of: %s\n", viper.GetViper().ConfigFileUsed())
		fmt.Println(viper.AllSettings())
	}

	rootCmd.AddCommand(i.IndexCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
