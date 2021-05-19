package models

import (
	log "github.com/sirupsen/logrus"
)

// Subject represents a participant in a RoleBinding.
type Subject struct {
	APIGroup string `yaml:"apiGroup"`
	Kind     string
	Name     string
}

// RoleBinding represents a Kubernetes RoleBinding object. It links
// one or more subjects to a role.
type RoleBinding struct {
	Resource `yaml:",inline"`
	RoleRef  Subject `yaml:"roleRef"`
	Subjects []Subject
}

// NewRoleBinding creates a new RoleBinding object. The
// "name" parameter is used to initialize "metadata.name". The "role"
// parameter is used to set the name of the "roleRef".
func NewRoleBinding(name string, role string) RoleBinding {
	if len(name) == 0 {
		log.Fatal("a group requires a name")
	}

	rsrc := RoleBinding{
		Resource: Resource{
			APIVersion: "rbac.authorization.k8s.io/v1",
			Kind:       "RoleBinding",
			Metadata: Metadata{
				Name: name,
			},
		},
		RoleRef: Subject{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     role,
		},
		Subjects: make([]Subject, 0),
	}
	return rsrc
}

// NewGroupSubject creates a new Subject referring
// to the named group.
func NewGroupSubject(groupName string) Subject {
	rsrc := Subject{
		APIGroup: "rbac.authorization.k8s.io",
		Kind:     "Group",
		Name:     groupName,
	}

	return rsrc
}

// AddGroup adds an entry for the named Group to the
// list of subjects in the RoleBinding.
func (rolebinding *RoleBinding) AddGroup(groupName string) {
	sub := NewGroupSubject(groupName)
	rolebinding.Subjects = append(rolebinding.Subjects, sub)
}
