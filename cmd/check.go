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
			username, _ := cmd.Flags().GetString("username")
			password, _ := cmd.Flags().GetString("password")
			database, _ := cmd.Flags().GetString("database")

			migr := migrate.NewMigrator(url, username, password, database)
			err := migr.Connect()
			if err != nil {
				fmt.Println(err.Error())
			}

			fmt.Println("Reading collections from database...")
			colls, err := migr.Collections()
			if err != nil {
				fmt.Println(err.Error())
			}
			for _, coll := range colls {
				fmt.Println("Found collection " + coll)
			}

			err = migr.Disconnect()
			if err != nil {
				fmt.Println(err.Error())
			}
		},
	}
)

func init() {
	viper.AutomaticEnv()
	checkCmd.Flags().String("dir", "./migrate", "the directory path to migrate")
	checkCmd.Flags().Bool("delete", false, "delete files not part of this package")
	checkCmd.Flags().String("host", "localhost", "the database host")
	checkCmd.Flags().Int("port", 27017, "the database port")
	checkCmd.Flags().StringP("database", "d", "database", "the database name")
	checkCmd.Flags().StringP("username", "u", "", "the database username")
	checkCmd.Flags().StringP("password", "p", "", "the database password")
	rootCmd.AddCommand(checkCmd)
}
