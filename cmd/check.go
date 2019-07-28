package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	checkCmd = &cobra.Command{
		Use:   "check",
		Short: "Outputs the tasks that the migrator will perform, without actually doing them.",
		Run: func(cmd *cobra.Command, _ []string) {
			fmt.Printf("Executing check command...\n")
		},
	}
)

func init() {
	checkCmd.Flags().Bool("delete", false, "delete files not part of this package")
	checkCmd.Flags().String("host", "localhost", "the database host")
	checkCmd.Flags().StringP("username", "u", "mongo", "the database username")
	checkCmd.Flags().StringP("password", "p", "mongo", "the database password")
	rootCmd.AddCommand(checkCmd)
}