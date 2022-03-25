package api

import (
	"fmt"
	"path/filepath"

	"github.com/operate-first/opfcli/constants"
	"github.com/operate-first/opfcli/models"
	"github.com/operate-first/opfcli/utils"
)

func (api *API) AddUsersToGroup(groupName string, users []string) error {

	groupPath := filepath.Join(api.RepoDirectory, api.AppName, constants.GroupPath, groupName, "group.yaml")
	exists, err := utils.PathExists(filepath.Dir(groupPath))
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("group %s does not exist", groupName)
	}

	err = models.AddUsersToGroup(filepath.Dir(groupPath), users)
	if err != nil {
		return err
	}

	commonOverlayPath := filepath.Join(api.RepoDirectory, api.AppName, constants.CommonOverlayPath)
	commonOverlayKustomizationPath := filepath.Join(commonOverlayPath, "kustomization.yaml")
	groupPath = fmt.Sprintf("../../../base/user.openshift.io/groups/%s", groupName)

	exist, err := utils.PathExists(commonOverlayKustomizationPath)
	if err != nil {
		return fmt.Errorf("encountered error when trying to verify kustomization in common overlay path: %s", commonOverlayKustomizationPath)
	} else if !exist {
		return fmt.Errorf("kustomization does not exist in common overlay path: %s", commonOverlayKustomizationPath)
	}

	kustom, err := models.KustomizeFromYAMLPath(commonOverlayKustomizationPath)
	if err != nil {
		return err
	}

	groupFound := false
	for _, resource := range kustom.Resources {
		if resource == groupPath {
			groupFound = true
			break
		}
	}

	if !groupFound {
		err = utils.AddKustomizeResources(commonOverlayPath, []string{groupPath})
		if err != nil {
			return err
		}
	}

	return nil

}
