package cmd

import (
	"github.com/operate-first/opfcli/api"
	"github.com/operate-first/opfcli/constants"
	"github.com/spf13/cobra"
)

func NewCmdInstallOperator(opfapi *api.API) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install-operator name catalog",
		Short: "Install an operator",
		Long: `Install an operator via Operator Hub.

- Register a new subscription
- Register a new operator group and namespace if requested
`,
		Args:          cobra.ExactArgs(2),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			source := args[1]
			namespace, err := cmd.Flags().GetString("namespace")
			if err != nil {
				return err
			}
			singleNamespace, err := cmd.Flags().GetBool("single-namespace")
			if err != nil {
				return err
			}
			channel, err := cmd.Flags().GetString("channel")
			if err != nil {
				return err
			}
			manual, err := cmd.Flags().GetBool("manual")
			if err != nil {
				return err
			}
			owner, err := cmd.Flags().GetString("owner")
			if err != nil {
				return err
			}

			if err := opfapi.CreateNamespace(namespace, owner, "", "", true, true, "", ""); err != nil {
				return err
			}

			if err := opfapi.CreateOperatorGroup(namespace, singleNamespace); err != nil {
				return err
			}

			return opfapi.CreateSubscription(
				name, source, namespace, channel, manual,
			)
		},
	}

	cmd.Flags().StringP("namespace", "n", constants.DefaultOperatorNamespace, "Namespace where the operator will be deployed to")
	cmd.Flags().StringP("owner", "o", constants.DefaultOwner, "If a new namespace is created, owner must be set as well")
	cmd.Flags().StringP("channel", "c", "", "Subscription channel")
	cmd.Flags().BoolP("manual", "m", false, "Require manual approval install strategy")
	cmd.Flags().BoolP("single-namespace", "s", false, "Sets operator to be installed as 'SingleNamespace'")
	return cmd
}
