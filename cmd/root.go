package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	alias     string
	role      string
	oneline   bool
	webView   bool
	htmlMode  bool
	outputURL bool
)

func rootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "awsman",
		Long: "awsman is aws account management tool.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(newAddCmd())
	cmd.AddCommand(newRemoveCmd())
	cmd.AddCommand(newListCmd())
	cmd.AddCommand(newSigninCmd())

	return cmd
}

func Execute() {
	cmd := rootCmd()
	cmd.SetOutput(os.Stdout)
	if err := cmd.Execute(); err != nil {
		cmd.SetOutput(os.Stderr)
		cmd.Println(err)
		os.Exit(1)
	}
}
