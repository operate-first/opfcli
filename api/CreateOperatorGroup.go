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

func (api *API) CreateOperatorGroup(namespace string, singleNamespace bool) error {

	path := filepath.Join(
		api.RepoDirectory, api.AppName,
		constants.OperatorGroupPath, namespace, "operatorgroup.yaml",
	)

	exists, err := utils.PathExists(filepath.Dir(path))
	if err != nil {
		return err
	}

	if exists {
		log.Warnf("operator group already exists in the target namespace (continuing)")
		return nil
	}

	og := models.NewOperatorGroup(
		namespace, singleNamespace,
	)

	ogOut := models.ToYAML(og)

	log.Printf("writing operator group definition to %s", filepath.Dir(path))
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("failed to create operator group in directory: %w", err)
	}

	err = ioutil.WriteFile(path, ogOut, 0644)
	if err != nil {
		return fmt.Errorf("failed to write operator group: %w", err)
	}

	kustomize := models.NewKustomization(
		[]string{"operatorgroup.yaml"}, []string{},
	)
	kustomize.Namespace = namespace

	err = kustomize.Write(filepath.Dir(path))
	if err != nil {
		return err
	}

	return nil
}
