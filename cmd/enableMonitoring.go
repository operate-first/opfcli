package cmd

import (
	"fmt"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"

	"github.com/operate-first/opfcli/utils"
)

func addMonitoringRBAC(projectName string) error {
	appName := config.GetString("app-name")
	path := filepath.Join(repoDirectory, appName, namespacePath, projectName)

	exists, err := utils.PathExists(path)
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("namespace %s does not exist", projectName)
	}

	log.Printf("enabling monitoring on namespace %s", projectName)

	return utils.AddKustomizeComponent(
		path,
		filepath.Join(componentRelPath, "monitoring-rbac"),
	)
}

var enableMonitoringCmd = &cobra.Command{
	Use:   "enable-monitoring namespace",
	Short: "Enable monitoring for a Kubernetes namespace",
	Long: `Enable monitoring fora Kubernetes namespace.

This will add a RoleBinding to the target namespace that permits
Prometheus to access certain metrics about pods, services, etc.`,
	Args:          cobra.ExactArgs(1),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return addMonitoringRBAC(args[0])
	},
}

func init() {
	rootCmd.AddCommand(enableMonitoringCmd)
}
