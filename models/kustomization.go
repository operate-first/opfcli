package models

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// Komponent represents a Kustomize Component. A Component is a collection of
// resources that can be included in a Kustomization file.
type Komponent struct {
	Resource  `yaml:",inline"`
	Resources []string `yaml:",omitempty"`
}

// Kustomization represents a kustomization file.
type Kustomization struct {
	Resource   `yaml:",inline"`
	Resources  []string `yaml:",omitempty"`
	Components []string `yaml:",omitempty"`
}

// NewKustomization creates a new Kustomization object.
func NewKustomization(resources, components []string) Kustomization {
	rsrc := Kustomization{
		Resource: Resource{
			APIVersion: "kustomize.config.k8s.io/v1beta1",
			Kind:       "Kustomization",
		},
		Resources:  resources,
		Components: components,
	}
	return rsrc
}

// NewKomponent creates a new Komponent object.
func NewKomponent(resources []string) Komponent {
	rsrc := Komponent{
		Resource: Resource{
			APIVersion: "kustomize.config.k8s.io/v1alpha1",
			Kind:       "Component",
		},
		Resources: resources,
	}
	return rsrc
}

// KustomizeFromYAMLPath reads the file at path and returns a
// Kustomization object.
func KustomizeFromYAMLPath(path string) (Kustomization, error) {
	var kustom Kustomization

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return Kustomization{}, err
	}

	err = yaml.Unmarshal(content, &kustom)
	if err != nil {
		return Kustomization{}, err
	}

	return kustom, nil
}

func (kustom *Kustomization) Write(path string) error {
	kustomOut := ToYAML(kustom)

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

func (komp *Komponent) Write(path string) error {
	kompOut := ToYAML(komp)

	log.Debugf("writing component kustomization for %s", path)
	err := ioutil.WriteFile(
		filepath.Join(path, "kustomization.yaml"),
		kompOut, 0644,
	)
	if err != nil {
		return fmt.Errorf("failed to write kustomization: %w", err)
	}

	return nil
}
