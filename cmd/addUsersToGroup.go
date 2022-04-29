package cmd

import (
	"strings"

	"github.com/operate-first/opfcli/api"
	"github.com/spf13/cobra"
)

func NewCmdAddUsersToGroup(opfapi *api.API) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "add-users group users cluster",
		Short:         "Add users to a group",
		Long:          `Add a comma seperated list of github handles to a group.`,
		Args:          cobra.ExactArgs(2),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			users := strings.Split(args[1], ",")
			return opfapi.AddUsersToGroup(args[0], users)
		},
	}
	return cmd
}
