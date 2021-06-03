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

func (api *API) CreateRoleBinding(groupName, roleName string) error {
	bindingName := fmt.Sprintf("project-%s-rolebindings", roleName)

	path := filepath.Join(
		api.RepoDirectory, api.AppName,
		constants.ComponentPath, bindingName, groupName, "rbac.yaml",
	)

	exists, err := utils.PathExists(filepath.Dir(path))
	if err != nil {
		return err
	}

	if exists {
		log.Printf("rolebinding already exists (continuing)")
		return nil
	}

	rbac := models.NewRoleBinding(
		fmt.Sprintf("namespace-%s-%s", roleName, groupName),
		roleName,
	)
	rbac.AddGroup(groupName)
	rbacOut := models.ToYAML(rbac)

	log.Printf("writing rbac definition to %s", filepath.Dir(path))
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("failed to create rolebinding directory: %w", err)
	}

	err = ioutil.WriteFile(path, rbacOut, 0644)
	if err != nil {
		return fmt.Errorf("failed to write rbac: %w", err)
	}

	komp := models.NewKomponent(
		[]string{"rbac.yaml"},
	)

	err = komp.Write(filepath.Dir(path))
	if err != nil {
		return err
	}

	return nil
}
