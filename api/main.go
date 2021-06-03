package api

import (
	"github.com/operate-first/opfcli/constants"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type API struct {
	AppName       string
	RepoDirectory string
}

func New(appName, repoDirectory string) *API {
	if appName == "" {
		appName = constants.DefaultAppName
	}

	if repoDirectory == "" {
		repoDirectory = "."
	}

	api := API{
		AppName:       appName,
		RepoDirectory: repoDirectory,
	}

	return &api
}

func (api *API) UpdateFromConfig(config *viper.Viper) {
	api.AppName = config.GetString("app-name")
	log.Debugf("got appname: %s", api.AppName)

	api.RepoDirectory = config.GetString("repo-dir")
	log.Debugf("got repodirectory: %s", api.RepoDirectory)
}
