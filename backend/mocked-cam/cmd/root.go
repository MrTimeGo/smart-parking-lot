package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{Use: "mocked-cam"}
)

func init() {
	rootCmd.AddCommand(runServiceCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
