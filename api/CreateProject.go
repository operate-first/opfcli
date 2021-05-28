package api

import (
	"fmt"
	"path/filepath"

	"github.com/operate-first/opfcli/constants"
	"github.com/operate-first/opfcli/utils"
)

func (api *API) ValidateQuota(projectQuota string) error {
	quotaPath := filepath.Join(
		api.RepoDirectory, api.AppName,
		constants.ComponentPath,
		"resourcequotas",
		projectQuota,
	)

	exists, err := utils.PathExists(quotaPath)
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("quota %s does not exist", projectQuota)
	}

	return nil
}

func (api *API) CreateProject(
	projectName, projectOwner, projectDescription string,
	projectQuota string,
	disableLimitrange bool,
) error {
	if projectQuota != "" {
		if err := api.ValidateQuota(projectQuota); err != nil {
			return err
		}
	}

	if err := api.CreateGroup(
		projectOwner,
		false,
	); err != nil {
		return err
	}

	if err := api.CreateRoleBinding(projectOwner, "admin"); err != nil {
		return err
	}

	return api.CreateNamespace(
		projectName,
		projectOwner,
		projectDescription,
		projectQuota,
		disableLimitrange,
	)
}
