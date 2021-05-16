// Package models contains a collection of structs that represent Kubernetes
// resources
package models

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// ToYAML converts a struct into a YAML string. If unable to marshal the
// struct, it prints an error and exits.
func ToYAML(resource interface{}) []byte {
	s, err := yaml.Marshal(&resource)
	if err != nil {
		log.Fatalf("failed converting resource to YAML: %v", err)
	}
	return s
}
