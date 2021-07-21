package api

import (
	// "fmt"

	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/operate-first/opfcli/constants"
	"github.com/operate-first/opfcli/models"
	"github.com/operate-first/opfcli/utils"
	log "github.com/sirupsen/logrus"
)

func (api *API) CreateCustomResourceQuota(namespace string, customResourceQuota models.CustomResourceQuota, existsOk bool) error {
	path := filepath.Join(
		api.RepoDirectory, api.AppName,
		constants.NamespacePath, namespace, "resourcequota.yaml",
	)

	exists, err := utils.PathExists(path)
	if err != nil {
		return err
	}

	if exists {
		if existsOk {
			log.Warnf("custom resource for namespace %s already exists (continuing)", namespace)
			return nil
		}
		return fmt.Errorf("resource quota already exists in namespace %s", namespace)
	}
	customResource := models.NewCustomResourceQuota(namespace, customResourceQuota)
	customResourceOut, err := models.ToYAML(customResource)
	if err != nil {
		return err
	}

	log.Printf("writing ResourceQuota definition to %s", filepath.Dir(path))
	if exist, err := utils.PathExists(filepath.Dir(path)); err != nil || !exist {
		if exist {
			return fmt.Errorf("error writing resource quota. error: %s", err)
		}
		return fmt.Errorf("directory should already have been created for the namespace. directory %s does not exist", namespace)
	}

	err = ioutil.WriteFile(path, customResourceOut, 0644)
	if err != nil {
		return fmt.Errorf("failed to write resource quota file: %w", err)
	}

	err = utils.AddKustomizeResources(filepath.Dir(path), []string{"resourcequota.yaml"})
	if err != nil {
		return fmt.Errorf("failed to append ResourceQuota resource to kustomization file")
	}

	return nil

}
