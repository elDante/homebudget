package cmd

import (
	"github.com/spf13/cobra"
)

// ConfigPath path of config file
var ConfigPath string

// RootCmd root cobra.Command variable
var RootCmd = &cobra.Command{
	Use:   "homebudget",
	Short: "homebudget is backend part of home budget tool",
	Long:  ``,
}

// Execute run cobra
func Execute() error {
	return RootCmd.Execute()
}

func init() {
	RootCmd.PersistentFlags().StringVar(&ConfigPath, "config", "config.toml", "Path to TOML configuration file")
}
