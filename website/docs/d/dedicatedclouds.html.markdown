---
layout: "ovh"
page_title: "OVH: dedicatedclouds"
sidebar_current: "docs-ovh-datasource-dedicatedclouds"
description: |-
  Get information about OVHCloud dedicatedClouds available in the current OVHCloud account
---

# ovh_dedicatedclouds (Data Source)

Use this data source to get the list of dedicatedClouds available in the current OVHCloud account

## Example Usage

```hcl
data "ovh_dedicatedclouds" "mydedicatedclouds" {}
```

## Argument Reference

There are no arguments to this datasource.

## Attributes Reference

The following attributes are exported:

* `dedicatedclouds`: The list of dedicatedClouds available in the current OVHCloud account 
