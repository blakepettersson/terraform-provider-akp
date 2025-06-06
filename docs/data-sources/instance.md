---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "akp_instance Data Source - akp"
subcategory: ""
description: |-
  Gets information about an Argo CD instance by its name
---

# akp_instance (Data Source)

Gets information about an Argo CD instance by its name

## Example Usage

```terraform
data "akp_instance" "example" {
  name = "test"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Instance name

### Read-Only

- `application_set_secret` (Map of String) stores secret key-value that will be used by `ApplicationSet`. For an example of how to use this in your ApplicationSet's pull request generator, see [here](https://github.com/argoproj/argo-cd/blob/master/docs/operator-manual/applicationset/Generators-Pull-Request.md#github). In this example, `tokenRef.secretName` would be application-set-secret.
- `argocd` (Attributes) Argo CD instance (see [below for nested schema](#nestedatt--argocd))
- `argocd_cm` (Map of String) is aligned with the options in `argocd-cm` ConfigMap as described in the [ArgoCD Atomic Configuration](https://argo-cd.readthedocs.io/en/stable/operator-manual/declarative-setup/#atomic-configuration). For a concrete example, refer to [this documentation](https://argo-cd.readthedocs.io/en/stable/operator-manual/argocd-cm-yaml/).
- `argocd_image_updater_config` (Map of String) configures Argo CD image updater, and it is aligned with `argocd-image-updater-config` ConfigMap of Argo CD, for available options and examples, refer to [this documentation](https://argocd-image-updater.readthedocs.io/en/stable/).
- `argocd_image_updater_secret` (Map of String) contains sensitive data (e.g., credentials for image updater to access registries) of Argo CD image updater, for available options and examples, refer to [this documentation](https://argocd-image-updater.readthedocs.io/en/stable/).
- `argocd_image_updater_ssh_config` (Map of String) contains the ssh configuration for Argo CD image updater, and it is aligned with `argocd-image-updater-ssh-config` ConfigMap of Argo CD, for available options and examples, refer to [this documentation](https://argocd-image-updater.readthedocs.io/en/stable/).
- `argocd_notifications_cm` (Map of String) configures Argo CD notifications, and it is aligned with `argocd-notifications-cm` ConfigMap of Argo CD, for more details and examples, refer to [this documentation](https://argocd-notifications.readthedocs.io/en/stable/).
- `argocd_notifications_secret` (Map of String) contains sensitive data of Argo CD notifications, and it is aligned with `argocd-notifications-secret` Secret of Argo CD, for more details and examples, refer to [this documentation](https://argocd-notifications.readthedocs.io/en/stable/services/overview/#sensitive-data).
- `argocd_rbac_cm` (Map of String) is aligned with the options in `argocd-rbac-cm` ConfigMap as described in the [ArgoCD Atomic Configuration](https://argo-cd.readthedocs.io/en/stable/operator-manual/declarative-setup/#atomic-configuration). For a concrete example, refer to [this documentation](https://argo-cd.readthedocs.io/en/stable/operator-manual/argocd-rbac-cm-yaml/).
- `argocd_secret` (Map of String) is aligned with the options in `argocd-secret` Secret as described in the [ArgoCD Atomic Configuration](https://argo-cd.readthedocs.io/en/stable/operator-manual/declarative-setup/#atomic-configuration). For a concrete example, refer to [this documentation](https://argo-cd.readthedocs.io/en/stable/operator-manual/argocd-secret-yaml/).
- `argocd_ssh_known_hosts_cm` (Map of String) is aligned with the options in `argocd-ssh-known-hosts-cm` ConfigMap as described in the [ArgoCD Atomic Configuration](https://argo-cd.readthedocs.io/en/stable/operator-manual/declarative-setup/#atomic-configuration). For a concrete example, refer to [this documentation](https://argo-cd.readthedocs.io/en/stable/operator-manual/argocd-ssh-known-hosts-cm-yaml/).
- `argocd_tls_certs_cm` (Map of String) is aligned with the options in `argocd-tls-certs-cm` ConfigMap as described in the [ArgoCD Atomic Configuration](https://argo-cd.readthedocs.io/en/stable/operator-manual/declarative-setup/#atomic-configuration). For a concrete example, refer to [this documentation](https://argo-cd.readthedocs.io/en/stable/operator-manual/argocd-tls-certs-cm-yaml/).
- `config_management_plugins` (Attributes Map) is a map of [Config Management Plugins](https://argo-cd.readthedocs.io/en/stable/operator-manual/config-management-plugins/#config-management-plugins), the key of map entry is the `name` of the plugin, and the value is the definition of the Config Management Plugin(v2). (see [below for nested schema](#nestedatt--config_management_plugins))
- `id` (String) Instance ID
- `repo_credential_secrets` (Map of Map of String) is a map of repo credential secrets, the key of map entry is the `name` of the secret, and the value is the aligned with options in `argocd-repositories.yaml.data` as described in the [ArgoCD Atomic Configuration](https://argo-cd.readthedocs.io/en/stable/operator-manual/declarative-setup/#atomic-configuration). For a concrete example, refer to [this documentation](https://argo-cd.readthedocs.io/en/stable/operator-manual/argocd-repositories-yaml/).
- `repo_template_credential_secrets` (Map of Map of String) is a map of repository credential templates secrets, the key of map entry is the `name` of the secret, and the value is the aligned with options in `argocd-repo-creds.yaml.data` as described in the [ArgoCD Atomic Configuration](https://argo-cd.readthedocs.io/en/stable/operator-manual/declarative-setup/#atomic-configuration). For a concrete example, refer to [this documentation](https://argo-cd.readthedocs.io/en/stable/operator-manual/argocd-repo-creds.yaml/).

<a id="nestedatt--argocd"></a>
### Nested Schema for `argocd`

Read-Only:

- `spec` (Attributes) Argo CD instance spec (see [below for nested schema](#nestedatt--argocd--spec))

<a id="nestedatt--argocd--spec"></a>
### Nested Schema for `argocd.spec`

Read-Only:

- `description` (String) Instance description
- `instance_spec` (Attributes) Argo CD instance spec (see [below for nested schema](#nestedatt--argocd--spec--instance_spec))
- `version` (String) Argo CD version. Should be equal to any Akuity [`argocd` image tag](https://quay.io/repository/akuity/argocd?tab=tags).

<a id="nestedatt--argocd--spec--instance_spec"></a>
### Nested Schema for `argocd.spec.instance_spec`

Read-Only:

- `agent_permissions_rules` (Attributes List) The ability to configure agent permissions rules. (see [below for nested schema](#nestedatt--argocd--spec--instance_spec--agent_permissions_rules))
- `app_in_any_namespace_config` (Attributes) App in any namespace config (see [below for nested schema](#nestedatt--argocd--spec--instance_spec--app_in_any_namespace_config))
- `app_set_delegate` (Attributes) Select cluster in which you want to Install Application Set controller (see [below for nested schema](#nestedatt--argocd--spec--instance_spec--app_set_delegate))
- `appset_plugins` (Attributes List) Application Set plugins (see [below for nested schema](#nestedatt--argocd--spec--instance_spec--appset_plugins))
- `appset_policy` (Attributes) Configures Application Set policy settings. (see [below for nested schema](#nestedatt--argocd--spec--instance_spec--appset_policy))
- `assistant_extension_enabled` (Boolean) Enable Powerful AI-powered assistant Extension. It helps analyze Kubernetes resources behavior and provides suggestions about resolving issues.
- `audit_extension_enabled` (Boolean) Enable Audit Extension. Set this to `true` to install Audit Extension to Argo CD instance.
- `backend_ip_allow_list_enabled` (Boolean) Enable ip allow list for cluster agents
- `cluster_customization_defaults` (Attributes) Default values for cluster agents (see [below for nested schema](#nestedatt--argocd--spec--instance_spec--cluster_customization_defaults))
- `crossplane_extension` (Attributes) Custom Resource Definition group name that identifies the Crossplane resource in kubernetes. We will include built-in crossplane resources. Note that you can use glob pattern to match the group. ie. *.crossplane.io (see [below for nested schema](#nestedatt--argocd--spec--instance_spec--crossplane_extension))
- `declarative_management_enabled` (Boolean) Enable Declarative Management
- `extensions` (Attributes List) Extensions (see [below for nested schema](#nestedatt--argocd--spec--instance_spec--extensions))
- `fqdn` (String) Configures the FQDN for the argocd instance, for ingress URL, domain suffix, etc.
- `host_aliases` (Attributes List) Host Aliases that override the DNS entries for control plane Argo CD components such as API Server and Dex. (see [below for nested schema](#nestedatt--argocd--spec--instance_spec--host_aliases))
- `image_updater_delegate` (Attributes) Select cluster in which you want to Install Image Updater (see [below for nested schema](#nestedatt--argocd--spec--instance_spec--image_updater_delegate))
- `image_updater_enabled` (Boolean) Enable Image Updater
- `ip_allow_list` (Attributes List) IP allow list (see [below for nested schema](#nestedatt--argocd--spec--instance_spec--ip_allow_list))
- `multi_cluster_k8s_dashboard_enabled` (Boolean) Enable the KubeVision feature
- `repo_server_delegate` (Attributes) In case some clusters don't have network access to your private Git provider you can delegate these operations to one specific cluster. (see [below for nested schema](#nestedatt--argocd--spec--instance_spec--repo_server_delegate))
- `subdomain` (String) Instance subdomain. By default equals to instance id
- `sync_history_extension_enabled` (Boolean) Enable Sync History Extension. Sync count and duration graphs as well as event details table on Argo CD application details page.

<a id="nestedatt--argocd--spec--instance_spec--agent_permissions_rules"></a>
### Nested Schema for `argocd.spec.instance_spec.agent_permissions_rules`

Read-Only:

- `api_groups` (List of String) API groups of the rule.
- `resources` (List of String) Resources of the rule.
- `verbs` (List of String) Verbs of the rule.


<a id="nestedatt--argocd--spec--instance_spec--app_in_any_namespace_config"></a>
### Nested Schema for `argocd.spec.instance_spec.app_in_any_namespace_config`

Read-Only:

- `enabled` (Boolean) Whether the app in any namespace config is enabled or not.


<a id="nestedatt--argocd--spec--instance_spec--app_set_delegate"></a>
### Nested Schema for `argocd.spec.instance_spec.app_set_delegate`

Read-Only:

- `managed_cluster` (Attributes) Use managed cluster (see [below for nested schema](#nestedatt--argocd--spec--instance_spec--app_set_delegate--managed_cluster))

<a id="nestedatt--argocd--spec--instance_spec--app_set_delegate--managed_cluster"></a>
### Nested Schema for `argocd.spec.instance_spec.app_set_delegate.managed_cluster`

Read-Only:

- `cluster_name` (String) Cluster name



<a id="nestedatt--argocd--spec--instance_spec--appset_plugins"></a>
### Nested Schema for `argocd.spec.instance_spec.appset_plugins`

Read-Only:

- `base_url` (String) Plugin base URL
- `name` (String) Plugin name
- `request_timeout` (Number) Plugin request timeout
- `token` (String) Plugin token


<a id="nestedatt--argocd--spec--instance_spec--appset_policy"></a>
### Nested Schema for `argocd.spec.instance_spec.appset_policy`

Read-Only:

- `override_policy` (Boolean) Allows per `ApplicationSet` sync policy.
- `policy` (String) Policy restricts what types of modifications will be made to managed Argo CD `Application` resources.
Available options: `sync`, `create-only`, `create-delete`, and `create-update`.
  - Policy `sync`(default): Update and delete are allowed.
  - Policy `create-only`: Prevents ApplicationSet controller from modifying or deleting Applications.
  - Policy `create-update`: Prevents ApplicationSet controller from deleting Applications. Update is allowed.
  - Policy `create-delete`: Prevents ApplicationSet controller from modifying Applications, Delete is allowed.


<a id="nestedatt--argocd--spec--instance_spec--cluster_customization_defaults"></a>
### Nested Schema for `argocd.spec.instance_spec.cluster_customization_defaults`

Read-Only:

- `app_replication` (Boolean) Enables Argo CD state replication to the managed cluster that allows disconnecting the cluster from Akuity Platform without losing core Argocd features
- `auto_upgrade_disabled` (Boolean) Disable Agents Auto Upgrade. On resource update terraform will try to update the agent if this is set to `true`. Otherwise agent will update itself automatically
- `kustomization` (String) Kustomize configuration that will be applied to generated agent installation manifests
- `redis_tunneling` (Boolean) Enables the ability to connect to Redis over a web-socket tunnel that allows using Akuity agent behind HTTPS proxy


<a id="nestedatt--argocd--spec--instance_spec--crossplane_extension"></a>
### Nested Schema for `argocd.spec.instance_spec.crossplane_extension`

Read-Only:

- `resources` (Attributes List) Glob patterns of the resources to match. (see [below for nested schema](#nestedatt--argocd--spec--instance_spec--crossplane_extension--resources))

<a id="nestedatt--argocd--spec--instance_spec--crossplane_extension--resources"></a>
### Nested Schema for `argocd.spec.instance_spec.crossplane_extension.resources`

Read-Only:

- `group` (String) Glob pattern of the group to match.



<a id="nestedatt--argocd--spec--instance_spec--extensions"></a>
### Nested Schema for `argocd.spec.instance_spec.extensions`

Read-Only:

- `id` (String) Extension ID
- `version` (String) Extension version


<a id="nestedatt--argocd--spec--instance_spec--host_aliases"></a>
### Nested Schema for `argocd.spec.instance_spec.host_aliases`

Read-Only:

- `hostnames` (List of String) Hostnames
- `ip` (String) IP address


<a id="nestedatt--argocd--spec--instance_spec--image_updater_delegate"></a>
### Nested Schema for `argocd.spec.instance_spec.image_updater_delegate`

Read-Only:

- `control_plane` (Boolean) If use control plane or not
- `managed_cluster` (Attributes) If use managed cluster or not (see [below for nested schema](#nestedatt--argocd--spec--instance_spec--image_updater_delegate--managed_cluster))

<a id="nestedatt--argocd--spec--instance_spec--image_updater_delegate--managed_cluster"></a>
### Nested Schema for `argocd.spec.instance_spec.image_updater_delegate.managed_cluster`

Read-Only:

- `cluster_name` (String) Cluster name



<a id="nestedatt--argocd--spec--instance_spec--ip_allow_list"></a>
### Nested Schema for `argocd.spec.instance_spec.ip_allow_list`

Read-Only:

- `description` (String) IP description
- `ip` (String) IP address


<a id="nestedatt--argocd--spec--instance_spec--repo_server_delegate"></a>
### Nested Schema for `argocd.spec.instance_spec.repo_server_delegate`

Read-Only:

- `control_plane` (Boolean) If use control plane or not
- `managed_cluster` (Attributes) If use managed cluster or not (see [below for nested schema](#nestedatt--argocd--spec--instance_spec--repo_server_delegate--managed_cluster))

<a id="nestedatt--argocd--spec--instance_spec--repo_server_delegate--managed_cluster"></a>
### Nested Schema for `argocd.spec.instance_spec.repo_server_delegate.managed_cluster`

Read-Only:

- `cluster_name` (String) Cluster name






<a id="nestedatt--config_management_plugins"></a>
### Nested Schema for `config_management_plugins`

Read-Only:

- `enabled` (Boolean) Whether this plugin is enabled or not. Default to false.
- `image` (String) Image to use for the plugin
- `spec` (Attributes) Plugin spec (see [below for nested schema](#nestedatt--config_management_plugins--spec))

<a id="nestedatt--config_management_plugins--spec"></a>
### Nested Schema for `config_management_plugins.spec`

Read-Only:

- `discover` (Attributes) The discovery config is applied to a repository. If every configured discovery tool matches, then the plugin may be used to generate manifests for Applications using the repository. If the discovery config is omitted then the plugin will not match any application but can still be invoked explicitly by specifying the plugin name in the app spec. Only one of fileName, find.glob, or find.command should be specified. If multiple are specified then only the first (in that order) is evaluated. (see [below for nested schema](#nestedatt--config_management_plugins--spec--discover))
- `generate` (Attributes) The generate command runs in the Application source directory each time manifests are generated. Standard output must be ONLY valid Kubernetes Objects in either YAML or JSON. A non-zero exit code will fail manifest generation. Error output will be sent to the UI, so avoid printing sensitive information (such as secrets). (see [below for nested schema](#nestedatt--config_management_plugins--spec--generate))
- `init` (Attributes) The init command runs in the Application source directory at the beginning of each manifest generation. The init command can output anything. A non-zero status code will fail manifest generation. Init always happens immediately before generate, but its output is not treated as manifests. This is a good place to, for example, download chart dependencies. (see [below for nested schema](#nestedatt--config_management_plugins--spec--init))
- `parameters` (Attributes) The parameters config describes what parameters the UI should display for an Application. It is up to the user to actually set parameters in the Application manifest (in spec.source.plugin.parameters). The announcements only inform the "Parameters" tab in the App Details page of the UI. (see [below for nested schema](#nestedatt--config_management_plugins--spec--parameters))
- `preserve_file_mode` (Boolean) Whether the plugin receives repository files with original file mode. Dangerous since the repository might have executable files. Set to true only if you trust the CMP plugin authors. Set to false by default.
- `version` (String) Plugin version

<a id="nestedatt--config_management_plugins--spec--discover"></a>
### Nested Schema for `config_management_plugins.spec.discover`

Read-Only:

- `file_name` (String) A glob pattern (https://pkg.go.dev/path/filepath#Glob) that is applied to the Application's source directory. If there is a match, this plugin may be used for the Application.
- `find` (Attributes) Find config (see [below for nested schema](#nestedatt--config_management_plugins--spec--discover--find))

<a id="nestedatt--config_management_plugins--spec--discover--find"></a>
### Nested Schema for `config_management_plugins.spec.discover.find`

Read-Only:

- `args` (List of String) Arguments for the find command
- `command` (List of String) The find command runs in the repository's root directory. To match, it must exit with status code 0 and produce non-empty output to standard out.
- `glob` (String) This does the same thing as `file_name`, but it supports double-start (nested directory) glob patterns.



<a id="nestedatt--config_management_plugins--spec--generate"></a>
### Nested Schema for `config_management_plugins.spec.generate`

Read-Only:

- `args` (List of String) Arguments of the command
- `command` (List of String) Command


<a id="nestedatt--config_management_plugins--spec--init"></a>
### Nested Schema for `config_management_plugins.spec.init`

Read-Only:

- `args` (List of String) Arguments of the command
- `command` (List of String) Command


<a id="nestedatt--config_management_plugins--spec--parameters"></a>
### Nested Schema for `config_management_plugins.spec.parameters`

Read-Only:

- `dynamic` (Attributes) Dynamic parameter announcements are announcements specific to an Application handled by this plugin. For example, the values for a Helm chart's values.yaml file could be sent as parameter announcements. (see [below for nested schema](#nestedatt--config_management_plugins--spec--parameters--dynamic))
- `static` (Attributes List) Static parameter announcements are sent to the UI for all Applications handled by this plugin. Think of the `string`, `array`, and `map` values set here as defaults. It is up to the plugin author to make sure that these default values actually reflect the plugin's behavior if the user doesn't explicitly set different values for those parameters. (see [below for nested schema](#nestedatt--config_management_plugins--spec--parameters--static))

<a id="nestedatt--config_management_plugins--spec--parameters--dynamic"></a>
### Nested Schema for `config_management_plugins.spec.parameters.dynamic`

Read-Only:

- `args` (List of String) Arguments of the command
- `command` (List of String) The command will run in an Application's source directory. Standard output must be JSON matching the schema of the static parameter announcements list.


<a id="nestedatt--config_management_plugins--spec--parameters--static"></a>
### Nested Schema for `config_management_plugins.spec.parameters.static`

Read-Only:

- `array` (List of String) This field communicates the parameter's default value to the UI if the parameter is an `array`.
- `collection_type` (String) Collection Type describes what type of value this parameter accepts (string, array, or map) and allows the UI to present a form to match that type. Default is `string`. This field must be present for non-string types. It will not be inferred from the presence of an `array` or `map` field.
- `item_type` (String) Item type tells the UI how to present the parameter's value (or, for arrays and maps, values). Default is `string`. Examples of other types which may be supported in the future are `boolean` or `number`. Even if the itemType is not `string`, the parameter value from the Application spec will be sent to the plugin as a string. It's up to the plugin to do the appropriate conversion.
- `map` (Map of String) This field communicates the parameter's default value to the UI if the parameter is a `map`.
- `name` (String) Parameter name
- `required` (Boolean) Whether the Parameter is required or not. If this field is set to true, the UI will indicate to the user that they must set the value. Default to false.
- `string` (String) This field communicates the parameter's default value to the UI if the parameter is a `string`.
- `title` (String) Title and description of the parameter
- `tooltip` (String) Tooltip of the Parameter, will be shown when hovering over the title
