package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/x-color/awsman/awsman"

	"github.com/spf13/cobra"
)

func runAddCmd(cmd *cobra.Command, args []string) error {
	id := args[0]
	if len(id) != 12 {
		return errors.New("invalid account id")
	}
	if _, err := strconv.Atoi(id); err != nil {
		return errors.New("invalid account id")
	}

	if alias == "" {
		if role == "" {
			alias = id
		} else {
			alias = fmt.Sprintf("%s - %s", id, role)
		}
	}
	return awsman.Add(id, role, alias)
}

func newAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add <account-id>",
		Short:   "Add account",
		Example: "  awsman add 123456789012 -r SampleUserRole -a xxx-prod",
		Args:    cobra.ExactArgs(1),
		RunE:    runAddCmd,
	}

	cmd.Flags().StringVarP(&alias, "alias", "a", "", "Alias of account. default alias is account id")
	cmd.Flags().StringVarP(&role, "role", "r", "", "IAM Role name")

	return cmd
}
