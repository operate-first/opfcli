package cmd

import (
	"github.com/spf13/cobra"

	"github.com/operate-first/opfcli/api"
)

func NewCmdEnableMonitoring(opfapi *api.API) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enable-monitoring namespace",
		Short: "Enable monitoring for a Kubernetes namespace",
		Long: `Enable monitoring fora Kubernetes namespace.

This will add a RoleBinding to the target namespace that permits
Prometheus to access certain metrics about pods, services, etc.`,
		Args:          cobra.ExactArgs(1),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return opfapi.AddComponent(args[0], "monitoring-rbac")
		},
	}

	return cmd
}
