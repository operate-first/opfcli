package utils

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/operate-first/opfcli/models"
	log "github.com/sirupsen/logrus"
)

// WriteKustomization creates a kustomization.yaml in the given path.
func WriteKustomization(path string, resources []string, components []string) error {
	kustom := models.NewKustomization(resources, components)

	kustomOut := models.ToYAML(kustom)

	log.Debugf("writing kustomization for %s", path)
	err := ioutil.WriteFile(
		filepath.Join(path, "kustomization.yaml"),
		kustomOut, 0644,
	)
	if err != nil {
		return fmt.Errorf("failed to write kustomization: %w", err)
	}

	return nil
}

// WriteComponent creates a kustomization.yaml in the given path. This
// creates a Component rather than a Kustomization.
func WriteComponent(path string, resources []string) error {
	kustom := models.NewKomponent(resources)

	kustomOut := models.ToYAML(kustom)

	log.Debugf("writing component kustomization for %s", path)
	err := ioutil.WriteFile(
		filepath.Join(path, "kustomization.yaml"),
		kustomOut, 0644,
	)
	if err != nil {
		return fmt.Errorf("failed to write component: %w", err)
	}

	return nil
}

// AddKustomizeComponent adds a component to an existing kustomization file.
func AddKustomizeComponent(path, componentPath string) error {
	kustomizePath := filepath.Join(path, "kustomization.yaml")
	log.Debugf("updating kustomization for %s", path)

	kustom, err := models.KustomizeFromYAMLPath(kustomizePath)
	if err != nil {
		return err
	}

	kustom.Components = append(kustom.Components, componentPath)
	kustomOut := models.ToYAML(kustom)
	err = ioutil.WriteFile(kustomizePath, kustomOut, 0644)
	if err != nil {
		return err
	}

	return nil
}
