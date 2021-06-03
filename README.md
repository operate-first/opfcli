# opfcli

## Building

To build this tool from a checked out copy of the repository, run:

```
make
```

This will produce an executable named `opfcli`.

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

Flags:
  -a, --app-name string   application name (default "cluster-scope")
  -f, --config string     configuration file
  -h, --help              help for opfcli
  -R, --repodir string    path to opf repository
```

### create-group

```
Create a group.

Create the group resource and associated kustomization file

Usage:
  opfcli create-group group [flags]

Flags:
  -h, --help                 help for create-project
```

### create-project

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
Enable monitoring fora Kubernetes namespace.

This will add a RoleBinding to the target namespace that permits
Prometheus to access certain metrics about pods, services, etc.

Usage:
  opfcli enable-monitoring namespace [flags]

Flags:
  -h, --help   help for enable-monitoring
```

### grant-access

```
Grant a group access to a namespace.

Grant a group access to a namespace with the specifed role
(admin, edit, or view).

Usage:
  opfcli grant-access namespace group role [flags]

Flags:
  -h, --help   help for grant-access
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
