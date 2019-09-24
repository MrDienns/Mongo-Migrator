package cmd

import (
	"fmt"

	"github.com/Dyescape/Mongo-Migrator/pkg/migrate"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	checkCmd = &cobra.Command{
		Use:   "check",
		Short: "Outputs the tasks that the migrator will perform, without actually doing them.",
		Run: func(cmd *cobra.Command, _ []string) {
			host, _ := cmd.Flags().GetString("host")
			port, _ := cmd.Flags().GetInt("port")
			url := fmt.Sprintf("%s:%v", host, port)
			username, _ = cmd.Flags().GetString("username")
			password, _ = cmd.Flags().GetString("password")

			migr := migrate.NewMigrator(url, username, password)
			err := migr.Connect()
			if err != nil {
				fmt.Println(err.Error())
			}
			migr.Disconnect()
		},
	}
)

func init() {
	viper.AutomaticEnv()
	checkCmd.Flags().String("path", "./migrate", "the directory path to migrate")
	checkCmd.Flags().Bool("delete", false, "delete files not part of this package")
	checkCmd.Flags().String("host", "localhost", "the database host")
	checkCmd.Flags().Int("port", 27017, "the database Wport")
	checkCmd.Flags().StringP("username", "u", "", "the database username")
	checkCmd.Flags().StringP("password", "p", "", "the database password")
	rootCmd.AddCommand(checkCmd)
}
