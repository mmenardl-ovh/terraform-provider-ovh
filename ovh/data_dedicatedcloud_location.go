package ovh

import (
	"fmt"
	"net/url"
	"sort"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ovh/terraform-provider-ovh/ovh/helpers/hashcode"
)

func dataSourceDedicatedCloudLocations() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedCloudLocationsRead,
		Schema: map[string]*schema.Schema{
			"locations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceDedicatedCloudLocationsRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	result := make([]string, 0)

	endpoint := "/dedicatedCloud/location"

	err := config.OVHClient.Get(endpoint, &result)
	if err != nil {
		return fmt.Errorf("Error calling GET %s:\n\t %q", endpoint, err)
	}

	sort.Strings(result)
	d.SetId(hashcode.Strings(result))
	d.Set("locations", result)

	return nil
}

func dataSourceDedicatedCloudLocation() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedCloudLocationRead,
		Schema: map[string]*schema.Schema{
			"pcc_zone": {
				Type:     schema.TypeString,
				Required: true,
			},
			"city": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"country_code": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceDedicatedCloudLocationRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	pccZone := d.Get("pcc_zone").(string)
	location := &DedicatedCloudLocation{}
	endpoint := fmt.Sprintf("/dedicatedCloud/location/%s", url.PathEscape(pccZone))
	err := config.OVHClient.Get(
		endpoint,
		&location,
	)
	if err != nil {
		return fmt.Errorf("Error calling GET %s:\n\t %q", endpoint, err)
	}

	d.SetId(fmt.Sprintf("%d", *location.Id))
	d.Set("city", *location.City)
	d.Set("country_code", *location.CountryCode)
	d.Set("id", *location.Id)
	d.Set("pcc_zone", *location.PccZone)

	return nil
}

// Location host profile - reseller call
func dataSourceDedicatedCloudLocationHostProfiles() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedCloudLocationHostProfilesRead,
		Schema: map[string]*schema.Schema{
			"pcc_zone": {
				Type:     schema.TypeString,
				Required: true,
			},
			"host_profiles": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceDedicatedCloudLocationHostProfilesRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	pccZone := d.Get("pcc_zone").(string)
	result := make([]string, 0)
	endpoint := fmt.Sprintf("/dedicatedCloud/location/%s/hostProfile", pccZone)
	err := config.OVHClient.Get(endpoint, &result)
	if err != nil {
		return fmt.Errorf("Error calling GET %s:\n\t %q", endpoint, err)
	}

	sort.Strings(result)
	d.SetId(hashcode.Strings(result))
	d.Set("host_profiles", result)

	return nil
}

// /dedicatedCloud/location/{pccZone}/hostProfile/{id}	# TODO

// Location hypervisor - reseller call
func dataSourceDedicatedCloudLocationHypervisors() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedCloudLocationHypervisorsRead,
		Schema: map[string]*schema.Schema{
			"pcc_zone": {
				Type:     schema.TypeString,
				Required: true,
			},
			"hypervisors": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceDedicatedCloudLocationHypervisorsRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	pccZone := d.Get("pcc_zone").(string)
	result := make([]string, 0)
	endpoint := fmt.Sprintf("/dedicatedCloud/location/%s/hypervisor", pccZone)
	err := config.OVHClient.Get(endpoint, &result)
	if err != nil {
		return fmt.Errorf("Error calling GET %s:\n\t %q", endpoint, err)
	}

	sort.Strings(result)
	d.SetId(hashcode.Strings(result))
	d.Set("hypervisors", result)

	return nil
}

// /dedicatedCloud/location/{pccZone}/hypervisor/{id} 	# TODO

// Location stock host - reseller call

// /dedicatedCloud/location/{pccZone}/stock/host    	# TODO

// Location stock pcc - reseller call

// /dedicatedCloud/location/{pccZone}/stock/pcc			# TODO

// Location stock zpool - reseller call

// /dedicatedCloud/location/{pccZone}/stock/zpool		# TODO
