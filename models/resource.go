package models

type kvmap map[string]string

// Metadata represents the metadata section of a Kubernetes resource.
type Metadata struct {
	Name        string
	Annotations kvmap `yaml:",omitempty"`
	Labels      kvmap `yaml:",omitempty"`
}

// Resource represents the core attributes of a Kubernetes resource. This
// struct is embedded in all other Kubernetes resource models.
type Resource struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   Metadata `yaml:",omitempty"`
}
