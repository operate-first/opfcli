# opfcli

## Building

To build this tool from a checked out copy of the repository, run:

```
make
```

This will produce an executable named `opfcli-<os>-<arch>` (for
example, `opfcli-linux-amd64`).

## Usage

```
A command line tool for Operate First GitOps.

Use opfcli to interact with an Operate First style Kubernetes
configuration repository.

Usage:
  opfcli [command]

Available Commands:
  create-group      Create a group
  create-project    Onboard a new project into Operate First
  enable-monitoring Enable monitoring for a Kubernetes namespace
  grant-access      Grant a group access to a namespace
  help              Help about any command
  onboard           Creates Groups, Namespaces, and Roles for a new Operate First project

Flags:
  -a, --app-name string   application name (default "cluster-scope")
  -f, --config string     configuration file
  -h, --help              help for opfcli
  -R, --repodir string    path to opf repository
```

## create-group

```
Create a group.

Create the group resource and associated kustomization file

Usage:
  opfcli create-group group [flags]

Flags:
  -h, --help                  help for create-project
  -d, --display-name          short team description for easy identification of project
  -n                          do not set a limitrange on this project
  -u, --users                 comma seperated list of users to add to the group
```

## create-project

```
Onboard a new project into Operate First.

- Register a new group
- Register a new namespace with appropriate role bindings for your group

Usage:
  opfcli create-project projectName projectOwner [flags]

Flags:
  -d, --description string   Team description
  -h, --help                 help for create-project
  -n, --no-limitrange        Do not set a limitrange on this project
  -q, --quota string         Set a quota on this project

Global Flags:
  -a, --app-name string      application name (default "cluster-scope")
  -f, --config-file string   configuration file
  -r, --repo-dir string      path to opf repository
```

## enable-monitoring

```
Enable monitoring for a Kubernetes namespace.

This will add a RoleBinding to the target namespace that permits
Prometheus to access certain metrics about pods, services, etc.

Usage:
  opfcli enable-monitoring namespace [flags]

Flags:
  -h, --help   help for enable-monitoring
```

## grant-access

```
Grant a group access to a namespace.

Grant a group access to a namespace with the specifed role
(admin, edit, or view).

Usage:
  opfcli grant-access namespace group role [flags]

Flags:
  -h, --help   help for grant-access
```

## Onboard

Onboard everything necessary for a new application into the Operate-First environment.

The onboard command expects an onboarding configuration file. These are the necessary and optional paramaters that make up the onboarding configuration file.

Necessary parameters:
  - `team_name` = name of the group looking to be onboarded
  - `<namespaces[i]>.name` = the name of the namespace(s) for you project (defined per namespace)
  - `project_description` = short summary of your project
  - `target_cluster` = name of the cluster to onboard to
  - `env` = name of the environment in which the `target_cluster` lives
  
Optional parameters:
  - `<namespaces[i]>.quota` = Namespace Quota
  - `<namespaces[i]>.custom_quota` = a customizeable resource quota to be applied to its namespace
  - `users` = users to be given access to the project when onboarding


This is an example of a valid configuration file using every option. Below each component will be discussed.

```
env: MOC
namespaces:
  - enable_monitoring: false
    name: testproject
    quota: testquota
    disable_limit_range: false
    project_display_name: testprojectdisplayname
    custom_quota:
      limits.cpu: '28'
      requests.cpu: '28'
      limits.memory: 32Gi
      requests.memory: 32Gi
      requests.storage: 100Gi
      count/objectbucketclaims.objectbucket.io: 1
project_description: This is the configuration for a sample project / app to onboard
target_cluster: Smaug
team_name: testgroup
users:
  - testing_gh_handle_1
  - testing_gh_handle_2

```

### `target_cluster` and `env`

Values for `target_cluster` can be obtained by running:
```
kustomize build https://github.com/operate-first/apps/acm/overlays/moc/infra/managedclusters?ref=master | yq e -N '.metadata.name' -
```
  - Requires: [yq](https://github.com/mikefarah/yq#install) and [kustomize](https://kubectl.docs.kubernetes.io/installation/kustomize/)

Similarly, values for `env` can be obtained by running:

```
curl -sX GET https://api.github.com/repos/operate-first/apps/contents/argocd/overlays/moc-infra/applications/envs | yq e '.[].name' -
```
  - Requires: [yq](https://github.com/mikefarah/yq#install)

NOTE: the `opfcli` is made for use in `operate-first/apps` which is why these urls are either `curl`ed or used in kustomize build.

If you cannot install `kustomize` and or `yq`, our last known options for `env` and `target_cluster` are as follows:

- env: `moc`
  - target_cluster: `smaug`
  - target_cluster: `infra`
- env: `osc`
  - target_cluster: `osc-cl1`
- env: `emea`:
  - target_cluster: `rick`
  - target_cluster: `morty`

### Namespace-Scoped Configurations

1. `quota`s: Quota's are optional values that refer to names of common [ResourceQuotas in `operate-first/apps`](https://github.com/operate-first/apps/tree/master/cluster-scope/components/resourcequotas). They limit the resources able to be used in a specific namespace.

2. `custom_quota`s: Custom quotas are optional configurations for a custom `ResourceQuota` to be used in a namespace. The values of a custom quota reflect those of our resourcequotas defined per namespace in operate-first/apps [see example](https://github.com/operate-first/apps/blob/master/cluster-scope/base/core/namespaces/sandbox/resourcequota.yaml).
    - If a valid `quota` is selected but a `custom_quota` is also defined in your configuration file, it will default to creating the namespace using the `custom_quota`.
    - All values are optional in a `custom_quota`.
    - `custom_quota` options: `[limits.cpu, requests.cpu, limits.memory, requests.memroy, requests.storage, count/objectbucketclaims.objectbucket.io]`

3. `disable_limit_range`: All requests to create and modify resources in Openshift are evaluated against each `LimitRange` object in the project. If the resource violates any of the enumerated constraints, the resource is rejected. Setting `disable_limit_range` to `true`, creates a namespace without the [operate-first default limit range](https://github.com/operate-first/apps/blob/master/cluster-scope/components/limitranges/default/limitrange.yaml).

4. `display_project_name`: display-names in openshift are user-defined strings which provide an alternative, more human-readable way to refer to a namespace. For more information on this checkout the [openshift docs](https://docs.openshift.com/online/pro/architecture/core_concepts/projects_and_users.html#projects) on projects.

### Users

GitHub handles that will be used to authenticate OCP clusters via GitHub. These GitHub usernames are converted to OCP users and are also used to extend permissions via Openshift RBAC.

### Usage

```
Usage:
  opfcli onboard <file-path>

Flags:
  -d, --display-name    provide a shorter name to use to refer to the project in openshift
  -n, --no-limitrange   Do not set a resource limitrange on this project

```

Use "opfcli [command] --help" for more information about a command.

## Configuration

The `opfcli` command will look for a configuration file `.opfcli.yaml`
in two places:

- It first checks in the top level of the current git repository. If
  you are running the `opfcli` command outside of a git repository it
  will instead check the current directory.

- If it doesn't find a local configuration file, it will look for
  `~/.opfcli.yaml`.

Use the `OPF_LOGLEVEL` environment variable to set logging verbosity.
The default is `1` (informational), but you can also set it to `0`
(warnings only) or `2` (or greater), which enables debug output.

### Available configuration options

- `app-name` -- sets the name of the directory containing your YAML
  resources. This defaults to `cluster-scope`.

## Examples

### Create a project

```
opfcli create-project project1 group1 -d "This is project1"
```

This will result in:

```
cluster-scope/
├── base
│   ├── core
│   │   └── namespaces
│   │       └── project1
│   │           ├── kustomization.yaml
│   │           └── namespace.yaml
│   └── user.openshift.io
│       └── groups
│           └── group1
│               ├── group.yaml
│               └── kustomization.yaml
└── components
    └── project-admin-rolebindings
        └── group1
            ├── kustomization.yaml
            └── rbac.yaml
```

### Create a group

```
opfcli create-group group2
```

This will result in:

```
cluster-scope/
└── base
    └── user.openshift.io
        └── groups
            └── group1
                ├── group.yaml
                └── kustomization.yaml
```

### Grant access to a project

```
opfcli grant-access project1 group2 view
```

This will result in:

```
cluster-scope/components/project-view-rolebindings/
└── group2
    ├── kustomization.yaml
    └── rbac.yaml
```

(And will modify
`cluster-scope/base/core/namespaces/project1/kustomization.yaml`)

## License

opfcli -- A tool for managing an Operate First style configuration repository.  
Copyright (C) 2021 Operate First Team

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
