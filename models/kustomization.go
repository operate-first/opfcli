package models

import (
	"io/ioutil"

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
func NewKustomization() Kustomization {
	rsrc := Kustomization{
		Resource: Resource{
			APIVersion: "kustomize.config.k8s.io/v1beta1",
			Kind:       "Kustomization",
		},
	}
	return rsrc
}

// NewKomponent creates a new Komponent object.
func NewKomponent() Komponent {
	rsrc := Komponent{
		Resource: Resource{
			APIVersion: "kustomize.config.k8s.io/v1alpha1",
			Kind:       "Component",
		},
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
