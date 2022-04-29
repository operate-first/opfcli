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

func (api *API) CreateNamespace(
	projectName string,
	projectOwner string,
	projectDisplayName string,
	projectQuota string,
	disableLimitrange bool,
	existsOk bool,
	onboardingIssue string,
	docs string,
) error {
	path := filepath.Join(
		api.RepoDirectory, api.AppName,
		constants.NamespacePath, projectName, "namespace.yaml")

	exists, err := utils.PathExists(filepath.Dir(path))
	if err != nil {
		return err
	}

	if exists {
		if existsOk {
			log.Warnf("namespace %s already exists (continuing)", projectName)
			// check if onboardingIssue and docs are set
			return nil
		}
		return fmt.Errorf("namespace %s already exists", projectName)
	}

	components := []string{
		filepath.Join(
			constants.ComponentRelPath,
			"project-admin-rolebindings",
			projectOwner,
		),
	}

	if len(projectQuota) > 0 {
		components = append(
			components,
			filepath.Join(
				constants.ComponentRelPath,
				"resourcequotas",
				projectQuota,
			))
	}

	if !disableLimitrange {
		components = append(
			components,
			filepath.Join(
				constants.ComponentRelPath,
				"limitranges",
				"default",
			))
	}

	ns := models.NewNamespace(projectName, projectOwner, projectDisplayName, onboardingIssue, docs)
	nsOut, err := models.ToYAML(ns)
	if err != nil {
		return err
	}

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
		components,
		projectName,
	)
	err = utils.WriteKustomization(filepath.Dir(path), kustom)
	if err != nil {
		return err
	}

	return nil
}
