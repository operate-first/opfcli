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

func (api *API) CreateNamespace(projectName, projectOwner, projectDescription string) error {
	path := filepath.Join(
		api.RepoDirectory, api.AppName,
		constants.NamespacePath, projectName, "namespace.yaml")

	exists, err := utils.PathExists(filepath.Dir(path))
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("namespace %s already exists", projectName)
	}

	ns := models.NewNamespace(projectName, projectOwner, projectDescription)
	nsOut := models.ToYAML(ns)

	log.Printf("writing namespace definition to %s", filepath.Dir(path))
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("failed to create namespace directory: %w", err)
	}

	err = ioutil.WriteFile(path, nsOut, 0644)
	if err != nil {
		return fmt.Errorf("failed to write namespace file: %w", err)
	}

	kustom := models.NewKustomization(
		[]string{"namespace.yaml"},
		[]string{
			filepath.Join(
				constants.ComponentRelPath,
				"project-admin-rolebindings",
				projectOwner,
			),
		},
	)
	err = kustom.Write(filepath.Dir(path))
	if err != nil {
		return err
	}

	return nil
}
