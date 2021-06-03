package cmd

import (
	"github.com/operate-first/opfcli/api"
	"github.com/spf13/cobra"
)

func NewCmdCreateGroup(opfapi *api.API) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-group group",
		Short: "Create a group",
		Long: `Create a group.

Create the group resource and associated kustomization file`,
		Args:          cobra.ExactArgs(1),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return opfapi.CreateGroup(args[0], false)
		},
	}

	return cmd
}
