---
layout: "ovh"
page_title: "OVH: dedicatedcloud allowed networks"
sidebar_current: "docs-ovh-datasource-dedicatedcloud-allowednetworks"
description: |-
  Get information about OVHCloud dedicatedCloud allowed networks
---

# ovh_dedicatedcloud_allowednetworks (Data Source)

Use this data source to get the list of networks trusted to connect to your Dedicated Cloud

## Example Usage

```hcl
data "ovh_dedicatedcloud_allowednetworks" "my_allowed_networks" {
  service_name = "pcc-1-2-3-4"
}
```

## Argument Reference

- `service_name`: (Required) Your Dedicated Cloud service name

## Attributes Reference

The following attributes are exported:

* `allowed_networks`: The list of the allowed network ids for your Dedicated Cloud
