package api

import (
	"fmt"
	"path/filepath"

	"github.com/operate-first/opfcli/constants"
	"github.com/operate-first/opfcli/utils"
	log "github.com/sirupsen/logrus"
)

func (api *API) AddComponent(projectName, componentName string) error {
	nsPath := filepath.Join(
		api.RepoDirectory, api.AppName, constants.NamespacePath, projectName,
	)

	exists, err := utils.PathExists(nsPath)
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("namespace %s does not exist", projectName)
	}

	exists, err = utils.PathExists(filepath.Join(
		api.RepoDirectory, api.AppName, constants.ComponentPath, componentName,
	))
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("component %s does not exist", componentName)
	}

	log.Printf("adding component %s to project %s", componentName, projectName)
	err = utils.AddKustomizeComponent(
		nsPath,
		filepath.Join(constants.ComponentRelPath, componentName),
	)
	if err != nil {
		return err
	}

	return nil
}
