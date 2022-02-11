# Terraform Provider Multipass

This provider is a plugin for Terraform that allows interactions with the [Multipass](https://multipass.run/) cli
utility.

## To Do
1. Add a lot more tests including [Acceptance Tests](https://www.terraform.io/plugin/sdkv2/testing/acceptance-tests/testcase)
2. Imports
3. Improve documentation

## Limitations

Resizing an instance isn't currently an option due to a limitation in Multipass [https://github.com/canonical/multipass/issues/1158](https://github.com/canonical/multipass/issues/1158).

## Provider Configuration

The provider configuration uses the current Multipass configuration by default. It uses the `multipass set flag=value`
to configure Multipass and requires elevated privileges. The parameters match the available options in the
[Multipass Set Command](https://multipass.run/docs/set-command).

| Parameter                             | Multipass Key                         |
|---------------------------------------|---------------------------------------|
| client_apps_windows_terminal_profiles | client.apps_windows-terminal_profiles |
| client_gui_autostart                  | client.gui_autostart                  |
| client_gui_hotkey                     | client.gui_hotkey                     |
| client_primary_name                   | client.primary-name                   |
| local_bridged_network                 | local.bridged_network                 |
| local_driver                          | local.driver                          |
| local_privileged_mounts               | local.privileged-mounts               |


## Examples

#### Default Instance

```terraform
resource "multipass_instance" "default" {
  name = "default-instance" // uses the default settings for Multipass
}
```

#### Specify The Image

```terraform
data "multipass_image" "bionic" {
  search = "18.04" // search looks in name and alias fields by default (only one alias must match if fuzzy == false)
  fuzzy  = false // true will match part of the search while false it must match exactly
}

resource "multipass_instance" "bionic" {
  name = "bionic-instance"

  image = data.multipass_image.bionic.name
}
```

#### Specify The Cores, Disk, and Memory

```terraform
resource "multipass_instance" "sized_instance" {
  cpus = 2
  disk = "20G"
  mem  = "4G"
  name = "sized-instance"
}
```

#### Specify Network Options

```terraform
resource "multipass_instance" "network_instance" {
  cpus = 2
  disk = "20G"
  mem  = "4G"
  name = "network-instance"

  network {
    name     = data.multipass_network.local_nat.name
    mode     = "auto"
    mac      = "00:11:22:33:44:55"
    required = false // If true, and the backend doesn't support the networks feature then the apply will fail
  }
}

resource "multipass_instance" "bridged_instance" {
  cpus = 2
  disk = "20G"
  mem  = "4G"
  name = "bridged-instance"

  bridged = true
}
```

#### Specify Mounts

```terraform
resource "multipass_instance" "mount_instance" {
  cpus = 2
  disk = "20G"
  mem  = "4G"
  name = "mount-instance"
  
  bridged = true

  mount {
    source_path = "/home/user/kubernetes"
    mount_path  = "/opt/kubernetes"
  }
}
```

#### Cluster

```terraform
resource "multipass_instance" "nodes" {
  count = 4

  cpus = 2
  disk = "20G"
  mem  = "4G"
  name = count.index == 0 ? "control-plane" : "worker-node-0${count.index}"

  bridged = true

  mount {
    source_path = "/home/ansible"
    mount_path  = "/opt/ansible"
  }

  mount {
    source_path = "/home/user/kubernetes"
    mount_path  = "/opt/kubernetes"
  }

  image = data.multipass_image.focal_img.name
}
```

#### Create An Alias
```terraform
resource "multipass_instance" "default" {
  name = "default-instance" // uses the default settings for Multipass
}

resource "multipass_alias" "default" {
  alias    = "bootstrap-${multipass_instance.default.name}"
  command  = "/usr/local/bin/bootstap"
  instance = multipass_instance.default.name
}
```