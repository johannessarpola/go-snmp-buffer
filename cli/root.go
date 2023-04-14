package admin

import (
	"fmt"
	"os"

	admin "github.com/johannessarpola/go-network-buffer/cli/admin/index"
	"github.com/spf13/cobra"
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
	rootCmd.AddCommand(admin.IndexCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
