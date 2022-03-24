package api

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/operate-first/opfcli/constants"
	"github.com/operate-first/opfcli/models"
	"github.com/operate-first/opfcli/utils"
)

func (api *API) Onboard(path string) error {
	onboardRequest, err := models.OnboardRequestFromYAMLPath(path)
	if err != nil {
		return err
	}

	if err := api.CreateGroup(
		onboardRequest.TeamName,
		onboardRequest.Users,
		true,
	); err != nil {
		return err
	}

	if err := api.CreateRoleBinding(onboardRequest.TeamName, "admin"); err != nil {
		return err
	}

	var tempQuota string
	var nsPaths []string
	emptyCustomQuota := models.CustomResourceQuota{}

	for i := range onboardRequest.Namespaces {
		var customQuotaRequested = onboardRequest.Namespaces[i].CustomQuota != emptyCustomQuota
		tempQuota = ""

		if !customQuotaRequested && onboardRequest.Namespaces[i].Quota != "" {
			tempQuota = strings.ToLower(onboardRequest.Namespaces[i].Quota)
			if err := api.ValidateQuota(tempQuota); err != nil {
				return err
			}
		}

		if err := api.CreateNamespace(
			onboardRequest.Namespaces[i].Name,
			onboardRequest.TeamName,
			onboardRequest.Namespaces[i].ProjectDisplayName,
			tempQuota,
			onboardRequest.Namespaces[i].DisableLimitRange,
			true,
		); err != nil {
			return err
		}

		if customQuotaRequested {
			if err := api.CreateCustomResourceQuota(onboardRequest.Namespaces[i].Name, onboardRequest.Namespaces[i].CustomQuota, true); err != nil {
				return err
			}
		}

		nsPaths = append(nsPaths, fmt.Sprintf("../../../../base/core/namespaces/%s", onboardRequest.Namespaces[i].Name))
	}

	prodOverlayPath := filepath.Join(
		api.RepoDirectory, api.AppName, constants.ProdOverlayPath, strings.ToLower(onboardRequest.Env), strings.ToLower(onboardRequest.TargetCluster),
	)
	prodOverlayKustomizationPath := filepath.Join(prodOverlayPath, "kustomization.yaml")

	exist, err := utils.PathExists(prodOverlayKustomizationPath)
	if err != nil {
		return fmt.Errorf("encountered error when trying to verify kustomization in cluster overlay path: %s", prodOverlayKustomizationPath)
	} else if !exist {
		return fmt.Errorf("kustomization does not exist in cluster overlay path: %s", prodOverlayKustomizationPath)
	}

	err = utils.AddKustomizeResources(prodOverlayPath, nsPaths)
	if err != nil {
		return err
	}

	commonOverlayPath := filepath.Join(api.RepoDirectory, api.AppName, constants.CommonOverlayPath)
	commonOverlayKustomizationPath := filepath.Join(commonOverlayPath, "kustomization.yaml")
	groupPath := fmt.Sprintf("../../../base/user.openshift.io/groups/%s", onboardRequest.TeamName)

	exist, err = utils.PathExists(commonOverlayKustomizationPath)
	if err != nil {
		return fmt.Errorf("encountered error when trying to verify kustomization in common overlay path: %s", commonOverlayKustomizationPath)
	} else if !exist {
		return fmt.Errorf("kustomization does not exist in common overlay path: %s", commonOverlayKustomizationPath)
	}

	err = utils.AddKustomizeResources(commonOverlayPath, []string{groupPath})
	if err != nil {
		return err
	}

	return nil

}
