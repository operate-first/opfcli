package cmd

import (
	"github.com/operate-first/opfcli/api"
	"github.com/spf13/cobra"
)

func NewCmdOnboard(opfapi *api.API) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "onboard file-path",
		Long:          "Onboard a project to a specific cluster in Operate-First using a provided onboard config file. Creates necessary groups and rolebindings, adds users to the project namespace(s), and deploys these changes to the specified cluster.",
		Short:         `Onboard a new project, group, namespace(s), and rolebindings into Operate First.`,
		Args:          cobra.ExactArgs(1),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			ConfigPath := args[0]
			return opfapi.Onboard(ConfigPath)
		},
	}
	return cmd
}
