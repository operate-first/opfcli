package models

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"
	yaml "sigs.k8s.io/yaml"
)

// Group represents a group of users.
type Group struct {
	Resource `yaml:",inline"`
	Users    []string
}

// NewGroup creates a new Group object.
func NewGroup(name string, users []string) Group {
	if len(name) == 0 {
		log.Fatal("a group requires a name")
	}

	rsrc := Group{
		Resource: Resource{
			APIVersion: "user.openshift.io/v1",
			Kind:       "Group",
			Metadata: Metadata{
				Name: name,
			},
		},
		Users: users,
	}
	return rsrc
}

func (g *Group) Unmarshal(y []byte) error {
	j, err := yaml.YAMLToJSON(y)
	if err != nil {
		return err
	}
	dec := json.NewDecoder(bytes.NewReader(j))
	dec.DisallowUnknownFields()
	var ng Group
	err = dec.Decode(&ng)
	if err != nil {
		return err
	}
	*g = ng
	return nil
}

func GroupFromYamlPath(path string) (Group, error) {
	var group Group

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return group, err
	}

	err = group.Unmarshal(content)
	if err != nil {
		return Group{}, err
	}

	return group, nil
}

func SortGroup(group Group) Group {
	users := make([]string, len(group.Users))
	copy(users, group.Users)
	if len(users) > 0 {
		if !sort.StringsAreSorted(users) {
			sort.Strings(users)
		}
	}
	return Group{
		Resource: group.Resource,
		Users:    users,
	}
}

func AddUsersToGroup(path string, users []string) error {
	groupPath := filepath.Join(path, "group.yaml")
	log.Debugf("updating kustomization for %s", path)
	group, err := GroupFromYamlPath(groupPath)
	if err != nil {
		return err
	}

	for _, userToAdd := range users {
		userToAdd = strings.Trim(userToAdd, " ")
		flagContainsUser := false
		for _, user := range group.Users {
			if user == userToAdd {
				flagContainsUser = true
			}
		}
		if !flagContainsUser {
			group.Users = append(group.Users, userToAdd)
		}
	}
	groupSorted := SortGroup(group)

	groupOut, err := ToYAML(groupSorted)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(groupPath, groupOut, 0644)
	if err != nil {
		return err
	}

	return nil
}
