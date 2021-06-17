package cmd

import (
	"github.com/operate-first/opfcli/api"
	"github.com/spf13/cobra"
)

func NewCmdCreateProject(opfapi *api.API) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-project projectName projectOwner",
		Short: "Onboard a new project into Operate First",
		Long: `Onboard a new project into Operate First.

- Register a new group
- Register a new namespace with appropriate role bindings for your group
`,
		Args:          cobra.ExactArgs(2),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			projectName := args[0]
			projectOwner := args[1]
			projectDisplayName, err := cmd.Flags().GetString("display-name")
			if err != nil {
				return err
			}
			disableLimitrange, err := cmd.Flags().GetBool("no-limitrange")
			if err != nil {
				return err
			}
			projectQuota, err := cmd.Flags().GetString("quota")
			if err != nil {
				return err
			}

			return opfapi.CreateProject(
				projectName, projectOwner, projectDisplayName,
				projectQuota,
				disableLimitrange,
			)
		},
	}

	cmd.Flags().StringP("display-name", "d", "", "Short team description for easy identification of project")
	cmd.Flags().StringP("quota", "q", "", "Set a quota on this project")
	cmd.Flags().BoolP("no-limitrange", "n", false, "Do not set a limitrange on this project")
	return cmd
}
