package cmd

import (
	"apus-sample/container"
	"apus-sample/internal/appctx"
	"context"
	"github.com/spf13/cobra"
)
var appContainer *container.Container
var start = &cobra.Command{
	Use: "start",
	Short: "start apus sample",
	Long: "start apus sample",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		err := appctx.InitContext()
		if err != nil{
			return err
		}
		appContainer,err = container.New()
		if err != nil{
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string)error {
		return appContainer.Start()
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		appContainer.Stop(context.Background())
	},
}