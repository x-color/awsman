package cmd

import (
	"github.com/x-color/awsman/awsman"

	"github.com/spf13/cobra"
)

func runRemoveCmd(cmd *cobra.Command, args []string) error {
	return awsman.Remove(args[0])
}

func newRemoveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "remove <account-id>",
		Short:   "Remove account",
		Example: "  awsman remove xxx-prod",
		Args:    cobra.ExactArgs(1),
		RunE:    runRemoveCmd,
	}

	return cmd
}
