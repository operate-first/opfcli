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
			projectDescription := cmd.Flag("description").Value.String()

			return opfapi.CreateProject(projectName, projectOwner, projectDescription)
		},
	}

	cmd.Flags().StringP("description", "d", "", "Team description")
	return cmd
}
