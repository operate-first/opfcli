package api

func (api *API) CreateProject(projectName, projectOwner, projectDescription string) error {
	if err := api.CreateGroup(
		projectOwner,
		false,
	); err != nil {
		return err
	}

	if err := api.CreateRoleBinding(projectOwner, "admin"); err != nil {
		return err
	}

	return api.CreateNamespace(
		projectName,
		projectOwner,
		projectDescription,
	)
}
