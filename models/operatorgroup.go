package models

import (
	log "github.com/sirupsen/logrus"
)

// OperatorGroup represents a Kubernetes OLM OperatorGroup resource.
type OperatorGroup struct {
	Resource `yaml:",inline"`
	Spec     OperatorGroupSpec
}

type OperatorGroupSpec struct {
	TargetNamespaces []string `yaml:"targetNamespaces,omitempty"`
}

// NewOperatorGroup creates a new OperatorGroup object. The object's
// metadata.name is set to the value of the "namespace" parameter.
// If the "singleNamespace" is true, spec.targetNamespaces is set to include
// the "namespace" string
func NewOperatorGroup(namespace string, singleNamespace bool) OperatorGroup {
	if len(namespace) == 0 {
		log.Fatal("an operator group requires a namespace")
	}

	targetNamespaces := []string{}
	if singleNamespace {
		targetNamespaces = append(targetNamespaces, namespace)
	}

	rsrc := OperatorGroup{
		Resource: Resource{
			APIVersion: "operators.coreos.com/v1",
			Kind:       "OperatorGroup",
			Metadata: Metadata{
				Name: namespace,
			},
		},
		Spec: OperatorGroupSpec{
			TargetNamespaces: targetNamespaces,
		},
	}

	return rsrc
}
