package cmd

import (
	"fmt"
	"os"

	"github.com/jonnyowenpowell/opr/config"
	"github.com/jonnyowenpowell/opr/template"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "opr",
	Short: "Open a PR from a template",
	Run: func(cmd *cobra.Command, args []string) {
		hasConfig, err := config.Exists()
		if err != nil {
			panic(err)
		}

		if !hasConfig {
			
			fmt.Printf("running config!")
			return
		}

		template.Run()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
