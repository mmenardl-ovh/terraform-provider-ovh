---
layout: "ovh"
page_title: "OVH: dedicatedcloud allowed network"
sidebar_current: "docs-ovh-datasource-dedicatedcloud-allowednetwork"
description: |-
  Get detail information about OVHCloud dedicatedCloud allowed network
---

# ovh_dedicatedcloud_allowednetwork (Data Source)

Use this data source to get the list of networks trusted to connect to your Dedicated Cloud

## Example Usage

```hcl
data "ovh_dedicatedcloud_allowednetwork" "my_allowed_network" {
  service_name = "pcc-1-2-3-4"
  network_access_id = 1
}
```

## Argument Reference

- `service_name`: (Required) Your Dedicated Cloud service name
- `network_access_id`: (Required) Allowed network id

## Attributes Reference

The following attributes are exported:

* `network_access_id`: Allowed network access id
* `description`: Description of your trusted network
* `network`: Network trusted on your Dedicated Cloud
* `service_name`: Dedicated Cloud where the network is trusted
* `state`: State of the trusted network
