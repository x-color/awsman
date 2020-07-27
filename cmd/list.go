package cmd

import (
	"github.com/x-color/awsman/awsman"

	"github.com/spf13/cobra"
)

func runListCmd(cmd *cobra.Command, args []string) error {
	var output string
	var err error

	switch {
	case webView:
		err = awsman.WebView()
	case oneline:
		output, err = awsman.Oneline()
	case htmlMode:
		output, err = awsman.HTMLView()
	default:
		output, err = awsman.TreeView()
	}

	if err != nil {
		return err
	}

	cmd.Print(output)

	return nil
}

func newListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List up accounts",
		Example: "  awsman list",
		Args:    cobra.NoArgs,
		RunE:    runListCmd,
	}

	cmd.Flags().BoolVarP(&oneline, "oneline", "l", false, "oneline output")
	cmd.Flags().BoolVar(&webView, "web", false, "open in browser")
	cmd.Flags().BoolVar(&htmlMode, "html", false, "html output")

	return cmd
}
