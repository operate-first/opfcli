package models

import (
	log "github.com/sirupsen/logrus"
)

// Namespace represents a Kubernetes Namespace.
type Namespace struct {
	Resource `yaml:",inline"`
}

// NewNamespace creates a new Namespace object. The object's
// metadata.name is set to the value of the "name" parameter. The
// "owner" and "description" parameters are used to initialize the
// "openshift.io/requester" and ""openshift.io/display-name"
// annotations.
func NewNamespace(name, owner, description string) Namespace {
	if len(name) == 0 {
		log.Fatal("a namespace requires a name")
	}

	if len(owner) == 0 {
		log.Fatal("a namespace requires an owner")
	}

	rsrc := Namespace{
		Resource: Resource{
			APIVersion: "v1",
			Kind:       "Namespace",
			Metadata: Metadata{
				Name:        name,
				Annotations: make(map[string]string),
			},
		},
	}
	rsrc.Metadata.Annotations["openshift.io/requester"] = owner
	if len(description) > 0 {
		rsrc.Metadata.Annotations["openshift.io/display-name"] = description
	}

	return rsrc
}
