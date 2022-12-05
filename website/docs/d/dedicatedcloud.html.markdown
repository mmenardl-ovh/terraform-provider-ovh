---
layout: "ovh"
page_title: "OVH: dedicatedcloud"
sidebar_current: "docs-ovh-datasource-dedicatedcloud"
description: |-
  Get information about current dedicatedCloud
---

# ovh_dedicatedcloud (Data Source)

Use this data source to get the information for the current dedicatedCloud

## Example Usage

```hcl
data "ovh_dedicatedcloud" "mydedicatedcloud" {
  service_name = "pcc-1-2-3-4"
}
```

## Argument Reference

* `service_name`: (Required) Dedicated Cloud service name

## Attributes Reference

The following attributes are exported:

* `advanced_security` : advanced security state
* `bandwidth`: The current bandwidth of you Dedicated Cloud
* `billing_type`: Billing type of your Dedicated Cloud
* `certified_interface_url`: Url to the Dedicated Cloud certified interface
* `commercialrange`: The current version of your Dedicated Cloud
* `description`: Description of your Dedicated Cloud
* `generation`: Generation of your Dedicated Cloud
* `location`: Datacenter were your Dedicated Cloud is physically located
* `management_interface`: The management interface name
* `product_reference`: The reference universe information for your Dedicated Cloud
* `service_name`: Service name of your Dedicated Cloud
* `servicepack_name`: Name of the current service pack
* `spla`: SPLA licensing state
* `sslv3`: Enable SSL v3 support. Warning : this option is not recommended as it was recognized as a security breach. If this is enabled, we advise you to enable the filtered User access policy
* `state`: Current state of your Dedicated Cloud
* `user_access_policy`: Access policy of your Dedicated Cloud : opened to every IPs or filtered
* `user_limit_concurrent_session`: The maximum amount of connected users allowed on the Dedicated Cloud management interface
* `user_logout_policy`: Which user will be disconnected first in case of quota of maximum connection is reached
* `user_session_timeout`: The timeout (in seconds) for the user sessions on the Dedicated Cloud management interface. 0 value disable the timeout
* `vscope_url`: Url to the Dedicated Cloud vScope interface
* `version`: Version of the management interface
* `web_interface_url`: Url to the Dedicated Cloud web interface
