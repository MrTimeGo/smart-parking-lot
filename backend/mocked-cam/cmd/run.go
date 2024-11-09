package cmd

import "github.com/spf13/cobra"

var (
	runServiceCmd = &cobra.Command{
		Use:   "run",
		Short: "Run the mocked camera service",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: add logic here
		},
	}
)
