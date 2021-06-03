package cmd

import (
	"github.com/operate-first/opfcli/api"
	"github.com/spf13/cobra"
)

func NewCmdGrantAccess(opfapi *api.API) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "grant-access namespace group role",
		Short: "Grant a group access to a namespace",
		Long: `Grant a group access to a namespace.

Grant a group access to a namespace with the specifed role
(admin, edit, or view).`,
		Args:          cobra.ExactArgs(3),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return opfapi.AddGroupRBAC(args[0], args[1], args[2])
		},
	}

	return cmd
}
