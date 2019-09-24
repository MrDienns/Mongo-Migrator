package cmd

import (
	"github.com/spf13/cobra"
)

var (
	host     string
	port     int
	username string
	password string
	debug    bool
	path     string
	rootCmd  = &cobra.Command{
		Use:   "mongomigrator",
		Short: "Migrates local data into MongoDB",
		Long: `A simple Go-based application which will migrate locally
package data (through a Docker container) into the specific MongoDB
instance. Can be used in combination with Kubernetes Jobs for easier
scheduling of migrations and automated updates.`,
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}
)

// Execute is the main entrypoint for the root command. This function
// will invoke the Cobra implementation of the command, which, in this
// case, will output the usage guides.
func Execute() {
	rootCmd.Execute()
}
