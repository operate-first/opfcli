package cmd

import (
	"fmt"

	"github.com/operate-first/opfcli/api"
	"github.com/operate-first/opfcli/version"
	"github.com/spf13/cobra"
)

func NewCmdVersion(opfapi *api.API) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "version",
		Short:         "Show opfcli version information",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Name: %s\n", version.Name)
			fmt.Printf("Version: %s\n", version.BuildVersion)
			fmt.Printf("Git Commit Hash: %s\n", version.BuildHash)
			fmt.Printf("Build Date: %s\n", version.BuildDate)
			return nil
		},
	}

	return cmd
}
