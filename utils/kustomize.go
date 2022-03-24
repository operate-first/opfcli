package utils

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/operate-first/opfcli/models"
	log "github.com/sirupsen/logrus"
	kustomizationAPITypes "sigs.k8s.io/kustomize/api/types"
)

// WriteKustomization creates a kustomization.yaml in the given path. This will work for a Kustomization or Components.
func WriteKustomization(path string, kustom kustomizationAPITypes.Kustomization) error { // move sort here
	kustomSorted := models.SortKustomization(kustom)
	kustomOut, err := models.ToYAML(kustomSorted)
	if err != nil {
		return err
	}
	log.Debugf("writing kustomization for %s", path)
	err = ioutil.WriteFile(
		filepath.Join(path, "kustomization.yaml"),
		kustomOut, 0644,
	)

	if err != nil {
		return fmt.Errorf("failed to write kustomization: %w", err)
	}

	return nil
}

// AddKustomizeComponent adds a component to an existing kustomization file.
func AddKustomizeComponent(path string, componentPaths []string) error {
	kustomizePath := filepath.Join(path, "kustomization.yaml")
	log.Debugf("updating kustomization for %s", path)

	kustom, err := models.KustomizeFromYAMLPath(kustomizePath)
	if err != nil {
		return err
	}

	for _, componentPath := range componentPaths {
		flagContainsComponentPath := false
		for _, kustomComponent := range kustom.Components {
			if kustomComponent == componentPath {
				flagContainsComponentPath = true
			}
		}
		if !flagContainsComponentPath {
			kustom.Components = append(kustom.Components, componentPath)
		}
	}
	kustomSorted := models.SortKustomization(kustom)

	kustomOut, err := models.ToYAML(kustomSorted)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(kustomizePath, kustomOut, 0644)
	if err != nil {
		return err
	}

	return nil
}

func AddKustomizeResources(path string, resourcePaths []string) error {
	kustomizePath := filepath.Join(path, "kustomization.yaml")
	log.Debugf("updating kustomization for %s", path)
	kustom, err := models.KustomizeFromYAMLPath(kustomizePath)
	if err != nil {
		return err
	}

	for _, resourcePath := range resourcePaths {
		flagContainsResourcePath := false
		for _, kustomResource := range kustom.Resources {
			if kustomResource == resourcePath {
				flagContainsResourcePath = true
			}
		}
		if !flagContainsResourcePath {
			kustom.Resources = append(kustom.Resources, resourcePath)
		}
	}
	kustomSorted := models.SortKustomization(kustom)

	kustomOut, err := models.ToYAML(kustomSorted)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(kustomizePath, kustomOut, 0644)
	if err != nil {
		return err
	}

	return nil
}
