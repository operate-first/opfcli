package models

import (
	"io/ioutil"
	"sort"

	kustomizationAPITypes "sigs.k8s.io/kustomize/api/types"
)

func NewKustomization(resources []string, components []string, namespace string) kustomizationAPITypes.Kustomization {
	rsrc := kustomizationAPITypes.Kustomization{
		TypeMeta: kustomizationAPITypes.TypeMeta{
			APIVersion: "kustomize.config.k8s.io/v1beta1",
			Kind:       "Kustomization",
		},
		Namespace:  namespace,
		Resources:  resources,
		Components: components,
	}
	return rsrc
}

// NewKomponent creates a new Komponent object.
func NewKomponent(resources []string) kustomizationAPITypes.Kustomization {
	rsrc := kustomizationAPITypes.Kustomization{
		TypeMeta: kustomizationAPITypes.TypeMeta{
			APIVersion: "kustomize.config.k8s.io/v1alpha1",
			Kind:       "Component",
		},
		Resources: resources,
	}
	return rsrc
}

// KustomizeFromYAMLPath reads the file at path and returns a
// Kustomization object.
func KustomizeFromYAMLPath(path string) (kustomizationAPITypes.Kustomization, error) {
	var kustom kustomizationAPITypes.Kustomization

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return kustom, err
	}

	err = kustom.Unmarshal(content)
	if err != nil {
		return kustomizationAPITypes.Kustomization{}, err
	}

	return kustom, nil
}

func SortKustomization(kustom kustomizationAPITypes.Kustomization) kustomizationAPITypes.Kustomization {
	resources := make([]string, len(kustom.Resources))
	copy(resources, kustom.Resources)
	if len(resources) > 0 {
		if !sort.StringsAreSorted(resources) {
			sort.Strings(resources)
		}
	}

	components := make([]string, len(kustom.Components))
	copy(components, kustom.Components)
	if len(components) > 0 {
		if !sort.StringsAreSorted(components) {
			sort.Strings(components)
		}
	}

	rsrc := kustomizationAPITypes.Kustomization{
		TypeMeta: kustomizationAPITypes.TypeMeta{
			APIVersion: kustom.APIVersion,
			Kind:       kustom.Kind,
		},
		Namespace:             kustom.Namespace,
		Resources:             resources,
		Components:            components,
		Generators:            kustom.Generators,
		PatchesStrategicMerge: kustom.PatchesStrategicMerge,
		PatchesJson6902:       kustom.PatchesJson6902,
	}
	return rsrc

}
