package models

import (
	log "github.com/sirupsen/logrus"
)

// Subscription represents a Kubernetes OLM Subscription resource.
type Subscription struct {
	Resource `yaml:",inline"`
	Spec     SubscriptionSpec
}

type SubscriptionSpec struct {
	Channel             string `yaml:",omitempty"`
	InstallPlanApproval string `yaml:"installPlanApproval"`
	Name                string
	Source              string
	SourceNamespace     string `yaml:"sourceNamespace"`
}

// NewSubscription creates a new Subscription object. The object's
// metadata.name is set to the value of the "name" parameter. The
// "catalog" parameter defines spec.source. The "strategy" parameter
// is allowed to be either Automatic or Manual and initialized
// spec.installPlanApproval. Optional "channel" parameters sets spec.channel
func NewSubscription(name, catalog, channel, strategy string) Subscription {
	if len(name) == 0 {
		log.Fatal("a subscription requires a name")
	}

	if len(catalog) == 0 {
		log.Fatal("a subscription requires a catalog source")
	}

	if len(channel) == 0 {
		log.Warn("no channel is set in subscription, you must define it in overlay")
	}

	if (strategy != "Automatic") && (strategy != "Manual") {
		log.Fatal("a subscription strategy must be either Automatic or Manual")
	}

	rsrc := Subscription{
		Resource: Resource{
			APIVersion: "operators.coreos.com/v1alpha1",
			Kind:       "Subscription",
			Metadata: Metadata{
				Name: name,
			},
		},
		Spec: SubscriptionSpec{
			Channel:             channel,
			InstallPlanApproval: strategy,
			Name:                name,
			Source:              catalog,
			SourceNamespace:     "openshift-marketplace",
		},
	}

	return rsrc
}
