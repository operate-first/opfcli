package models

import (
	"fmt"
)

type ResourceQuotaSpec struct {
	Hard CustomResourceQuota
}

type ResourceQuota struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   Metadata
	Spec       ResourceQuotaSpec
}

type CustomResourceQuota struct {
	LimitsCPU       string `yaml:"limits.cpu,omitempty"`
	LimitsMemory    string `yaml:"limits.memory,omitempty"`
	Buckets         int    `yaml:"count/objectbucketclaims.objectbucket.io,omitempty"`
	RequestsCPU     string `yaml:"requests.cpu,omitempty"`
	RequestsMemory  string `yaml:"requests.memory,omitempty"`
	RequestsStorage string `yaml:"requests.storage,omitempty"`
}

func NewCustomResourceQuota(namespace string, customResourceQuota CustomResourceQuota) ResourceQuota {
	resourceQuotaResource := ResourceQuota{
		APIVersion: "v1",
		Kind:       "ResourceQuota",
		Metadata: Metadata{
			Name: fmt.Sprintf("%s-custom", namespace),
		},
		Spec: ResourceQuotaSpec{
			Hard: CustomResourceQuota{
				LimitsCPU:       customResourceQuota.LimitsCPU,
				LimitsMemory:    customResourceQuota.LimitsMemory,
				RequestsCPU:     customResourceQuota.RequestsCPU,
				RequestsMemory:  customResourceQuota.RequestsMemory,
				RequestsStorage: customResourceQuota.RequestsStorage,
				Buckets:         customResourceQuota.Buckets,
			},
		},
	}
	return resourceQuotaResource
}
