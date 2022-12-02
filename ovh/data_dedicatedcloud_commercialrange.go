package ovh

import (
	"fmt"
	"net/url"
	"sort"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ovh/terraform-provider-ovh/ovh/helpers/hashcode"
)

func dataSourceDedicatedCloudCommercialRanges() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedCloudCommercialRangesRead,
		Schema: map[string]*schema.Schema{
			"result": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceDedicatedCloudCommercialRangesRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	result := make([]string, 0)
	err := config.OVHClient.Get("/dedicatedCloud/commercialRange", &result)
	if err != nil {
		return fmt.Errorf("Unable to retrieve /dedicatedCloud/commercialRange information:\n\t %q", err)
	}

	sort.Strings(result)
	d.SetId(hashcode.Strings(result))
	d.Set("result", result)

	return nil
}

func dataSourceDedicatedCloudCommercianRange() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedCloudCommercialRangeRead,
		Schema: map[string]*schema.Schema{
			"commercialrange_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"allowed_hypervisor_versions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"allowed_network_roles": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"dedicatedcloud_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"range": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceDedicatedCloudCommercialRangeRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	commercialRangeName := d.Get("commercialrange_name").(string)
	commercialRange := &DedicatedCloudCommercialRange{}
	err := config.OVHClient.Get(
		fmt.Sprintf("/dedicatedCloud/commercialRange/%s", url.PathEscape(commercialRangeName)),
		&commercialRange,
	)
	if err != nil {
		return fmt.Errorf("Unable to retrieve /dedicatedCloud/commercialRange/{commercialRangeName} information:\n\t %q", err)
	}

	d.SetId(*commercialRange.CommercialRangeName)
	d.Set("allowed_hypervisor_versions", commercialRange.AllowedHypervisorVersions)
	d.Set("allowed_network_roles", commercialRange.AllowedNetworkRoles)
	d.Set("commercialrange_name", commercialRange.CommercialRangeName)
	d.Set("dedicatedcloud_version", commercialRange.DedicatedCloudVersion)
	d.Set("range", commercialRange.Range)

	return nil
}
