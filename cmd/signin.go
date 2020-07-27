package cmd

import (
	"github.com/x-color/awsman/awsman"

	"github.com/spf13/cobra"
)

func runSigninCmd(cmd *cobra.Command, args []string) error {
	var output string
	var err error

	if outputURL {
		output, err = awsman.SignInURL(args[0])
	} else {
		err = awsman.SignIn(args[0])
	}
	if err != nil {
		return err
	}

	cmd.Println(output)

	return nil
}

func newSigninCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "signin <alias>",
		Short:   "Signin account",
		Example: "  awsman signin xxx-prod",
		Args:    cobra.ExactArgs(1),
		RunE:    runSigninCmd,
	}

	cmd.Flags().BoolVarP(&outputURL, "url", "u", false, "output url to sign-in")

	return cmd
}
