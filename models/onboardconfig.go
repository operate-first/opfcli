package models

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type OnboardNamespace struct {
	EnableMonitoring   bool
	Name               string
	Quota              string
	CustomQuota        CustomResourceQuota `yaml:"custom_quota,omitempty"`
	DisableLimitRange  bool                `yaml:"disable_limit_range,omitempty"`
	ProjectDisplayName string              `yaml:"project_display_name,omitempty"`
}

type OnboardingRequest struct {
	Env                string `yaml:",omitempty"`
	Namespaces         []OnboardNamespace
	ProjectDescription string   `yaml:"project_description"`
	TargetCluster      string   `yaml:"target_cluster"`
	TeamName           string   `yaml:"team_name"`
	Users              []string `yaml:",omitempty"`
	OnboardingIssue    string   `yaml:"onboarding_issue,omitempty"`
	Docs               string   `yaml:"docs,omitempty"`
}

func OnboardRequestFromYAMLPath(path string) (OnboardingRequest, error) {

	var OnboardRequest OnboardingRequest

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return OnboardingRequest{}, err
	}

	err = yaml.Unmarshal(content, &OnboardRequest)
	if err != nil {
		return OnboardingRequest{}, err
	}

	return OnboardRequest, nil
}
