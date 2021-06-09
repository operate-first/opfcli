package cmd

import (
	"fmt"

	"github.com/operate-first/opfcli/api"
	"github.com/spf13/cobra"
)

func NewCmdKafkaCaCrt(opfapi *api.API) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "kafka-ca-crt",
		Short:         "Retrieve kafka CA certificate to stdout",
		Args:          cobra.ExactArgs(0),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			secret, err := opfapi.GetSecretDataString(
				"opf-kafka",
				"odh-message-bus-cluster-ca-cert",
				"ca.crt",
			)
			if err != nil {
				return err
			}

			fmt.Printf("%s\n", secret)
			return nil
		},
	}

	return cmd
}
