package cmd

import (
	"apus-sample/internal/appconf"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use: "apus-sample",
	Short: "Apus Sample Project",
	Long: "Apus Sample Project ",
	PersistentPreRunE: func(cmd *cobra.Command, args []string)error {
		return appconf.LoadConfig()
	},
}

func Execute() error {
	return root.Execute()
}

func init()  {
	root.AddCommand(start)
}