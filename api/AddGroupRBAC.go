package api

import (
	"fmt"
	"path/filepath"

	"github.com/operate-first/opfcli/constants"
	"github.com/operate-first/opfcli/utils"
	log "github.com/sirupsen/logrus"
)

var validRoles = map[string]bool{
	"admin": true,
	"view":  true,
	"edit":  true,
}

func (api *API) AddGroupRBAC(projectName, groupName, roleName string) error {
	if !validRoles[roleName] {
		return fmt.Errorf("no such role named %q", roleName)
	}

	bindingName := fmt.Sprintf("project-%s-rolebindings", roleName)

	nsPath := filepath.Join(
		api.RepoDirectory, api.AppName, constants.NamespacePath, projectName,
	)

	groupPath := filepath.Join(
		api.RepoDirectory, api.AppName, constants.GroupPath, groupName,
	)

	exists, err := utils.PathExists(nsPath)
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("namespace %s does not exist", projectName)
	}

	exists, err = utils.PathExists(groupPath)
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("group %s does not exist", groupName)
	}

	if err = api.CreateRoleBinding(groupName, roleName); err != nil {
		return err
	}

	log.Printf("granting %s role %s on %s", groupName, roleName, projectName)
	err = utils.AddKustomizeComponent(
		nsPath,
		filepath.Join(constants.ComponentRelPath, bindingName, groupName),
	)
	if err != nil {
		return err
	}

	return nil
}
