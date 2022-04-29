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
// "owner" and "displayName" parameters are used to initialize the
// "openshift.io/requester" and ""openshift.io/display-name"
// annotations.
func NewNamespace(name string, owner string, displayName string, onboardingIssue string, docs string) Namespace {
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
	if len(onboardingIssue) > 0 || len(docs) > 0 {
		rsrc.Metadata.Annotations["op1st/project-owner"] = owner
	}
	if len(displayName) > 0 {
		rsrc.Metadata.Annotations["openshift.io/display-name"] = displayName
	}

	if len(onboardingIssue) > 0 {
		rsrc.Metadata.Annotations["op1st/onboarding-issue"] = onboardingIssue
	}
	if len(docs) > 0 {
		rsrc.Metadata.Annotations["op1st/docs"] = docs
	}

	return rsrc
}
