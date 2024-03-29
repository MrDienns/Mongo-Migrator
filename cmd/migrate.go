package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Executes the database migration",
		Run: func(cmd *cobra.Command, _ []string) {
			fmt.Printf("Executing migrate command...\n")
		},
	}
)

func init() {
	viper.AutomaticEnv()
	checkCmd.Flags().String("path", "./migrate", "the directory path to migrate")
	migrateCmd.Flags().Bool("delete", false, "delete files not part of this package")
	migrateCmd.Flags().BoolP("debug", "d", false, "outputs detailed activity information")
	migrateCmd.Flags().String("host", "localhost", "the database host")
	migrateCmd.Flags().Int("port", 27017, "the database port")
	migrateCmd.Flags().StringP("username", "u", "mongo", "the database username")
	migrateCmd.Flags().StringP("password", "p", "mongo", "the database password")
	rootCmd.AddCommand(migrateCmd)
}
