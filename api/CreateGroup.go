package api

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/operate-first/opfcli/constants"
	"github.com/operate-first/opfcli/models"
	"github.com/operate-first/opfcli/utils"
	log "github.com/sirupsen/logrus"
)

func (api *API) CreateGroup(groupName string, users []string, existsOk bool) error {
	path := filepath.Join(
		api.RepoDirectory, api.AppName,
		constants.GroupPath, groupName, "group.yaml")

	exists, err := utils.PathExists(filepath.Dir(path))
	if err != nil {
		return err
	}

	if exists {
		if existsOk {
			log.Printf("group %s already exists (continuing)", groupName)
			return nil
		}
		return fmt.Errorf("group %s already exists", groupName)
	}

	group := models.NewGroup(groupName, users)
	groupOut, err := models.ToYAML(group)
	if err != nil {
		return err
	}

	log.Printf("writing group definition to %s", filepath.Dir(path))
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("failed to create group directory: %w", err)
	}

	err = ioutil.WriteFile(path, groupOut, 0644)
	if err != nil {
		return fmt.Errorf("failed to write group: %w", err)
	}

	kustom := models.NewKustomization(
		[]string{"group.yaml"},
		nil,
		"",
	)

	err = utils.WriteKustomization(filepath.Dir(path), kustom)
	if err != nil {
		return err
	}

	return nil
}
