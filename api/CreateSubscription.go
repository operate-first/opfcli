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

func (api *API) CreateSubscription(name, catalog, namespace, channel string, manual bool) error {
	path := filepath.Join(
		api.RepoDirectory, api.AppName,
		constants.SubscriptionPath, name, "subscription.yaml",
	)

	exists, err := utils.PathExists(filepath.Dir(path))
	if err != nil {
		return err
	}

	if exists {
		log.Warnf("subscription already exists (continuing)")
		return nil
	}

	strategy := "Automatic"
	if manual {
		strategy = "Manual"
	}

	sub := models.NewSubscription(
		name, catalog, channel, strategy,
	)

	subOut := models.ToYAML(sub)

	log.Printf("writing subscription definition to %s", filepath.Dir(path))
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("failed to create subscription directory: %w", err)
	}

	err = ioutil.WriteFile(path, subOut, 0644)
	if err != nil {
		return fmt.Errorf("failed to write subscription: %w", err)
	}

	kustomize := models.NewKustomization(
		[]string{"subscription.yaml"}, nil, "",
	)
	kustomize.Namespace = namespace

	err = kustomize.Write(filepath.Dir(path))
	if err != nil {
		return err
	}

	return nil
}
