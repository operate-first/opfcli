// Package cmd implements the commands supported by opfcli.
package cmd

import (
	"path/filepath"
	"strings"

	"github.com/operate-first/opfcli/api"
	"github.com/operate-first/opfcli/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func initConfig(cmd *cobra.Command, opfapi *api.API, config *viper.Viper) error {
	cfgFile, err := cmd.Flags().GetString("config-file")
	if err != nil {
		return err
	}

	repoDirectory, err := cmd.Flags().GetString("repo-dir")
	if err != nil {
		return err
	}

	if repoDirectory == "" {
		repoDirectory, err = utils.FindRepoDir()
		if err != nil {
			repoDirectory, err = filepath.Abs(".")
			if err != nil {
				log.Fatalf("failed to determine repository directory: %v", err)
			}
		}
	}
	config.SetDefault("repo-dir", repoDirectory)
	log.Debugf("using %s as repository directory", repoDirectory)

	if cfgFile != "" {
		log.Debugf("using explicit configuration file %s", cfgFile)
		config.SetConfigFile(cfgFile)
	} else {
		config.AddConfigPath(repoDirectory)
	}
	if err := config.ReadInConfig(); err != nil {
		log.Debugf("failed to read configuration: %v", err)
	}
	log.Debugf("read configuration from %s", config.ConfigFileUsed())

	opfapi.UpdateFromConfig(config)

	return nil
}

func NewConfig() *viper.Viper {
	config := viper.New()

	config.SetEnvPrefix("opf")
	config.AutomaticEnv()
	config.SetConfigName(".opfcli")

	replacer := strings.NewReplacer("-", "_")
	config.SetEnvKeyReplacer(replacer)

	log.Debugf("returning new config")

	return config
}

func NewCmdRoot() *cobra.Command {
	opfapi := &api.API{}
	config := NewConfig()

	cmd := &cobra.Command{
		Use:   "opfcli",
		Short: "A command line tool for Operate First GitOps",
		Long: `A command line tool for Operate First GitOps.

Use opfcli to interact with an Operate First style Kubernetes
configuration repository.`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return initConfig(cmd, opfapi, config)
		},
	}

	cmd.PersistentFlags().StringP(
		"config-file", "f", "", "configuration file")
	cmd.PersistentFlags().StringP(
		"app-name", "a", "cluster-scope", "application name")
	cmd.PersistentFlags().StringP(
		"repo-dir", "r", "", "path to opf repository")

	if err := config.BindPFlag("app-name", cmd.PersistentFlags().Lookup("app-name")); err != nil {
		log.Fatalf("failed to bind flag app-name: %v", err)
	}

	if err := config.BindPFlag("repo-dir", cmd.PersistentFlags().Lookup("repo-dir")); err != nil {
		log.Fatalf("failed to bind flag repo-dir: %v", err)
	}

	cmd.AddCommand(
		NewCmdVersion(opfapi),
		NewCmdCreateGroup(opfapi),
		NewCmdCreateProject(opfapi),
		NewCmdGrantAccess(opfapi),
		NewCmdEnableMonitoring(opfapi),
		NewCmdCompletion(),
	)

	return cmd
}
