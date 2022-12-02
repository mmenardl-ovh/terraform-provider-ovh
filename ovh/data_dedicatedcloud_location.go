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

func dataSourceDedicatedCloudLocationsRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	result := make([]string, 0)
	err := config.OVHClient.Get("/dedicatedCloud/location", &result)
	if err != nil {
		return fmt.Errorf("Unable to retrieve /dedicatedCloud/location information:\n\t %q", err)
	}

	sort.Strings(result)
	d.SetId(hashcode.Strings(result))
	d.Set("result", result)

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
	err := config.OVHClient.Get(
		fmt.Sprintf("/dedicatedCloud/location/%s", url.PathEscape(pccZone)),
		&location,
	)

	if err != nil {
		return fmt.Errorf("Unable to retrieve /dedicatedCloud/location/{pccZone} information:\n\t %q", err)
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

func dataSourceDedicatedCloudLocationHostProfilesRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	pccZone := d.Get("pcc_zone").(string)
	result := make([]string, 0)
	err := config.OVHClient.Get(fmt.Sprintf("/dedicatedCloud/location/%s/hostProfile", pccZone), &result)
	if err != nil {
		return fmt.Errorf("Unable to retrieve /dedicatedCloud/location/{pccZone}/hostProfile information:\n\t %q", err)
	}

	sort.Strings(result)
	d.SetId(hashcode.Strings(result))
	d.Set("result", result)

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

func dataSourceDedicatedCloudLocationHypervisorsRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	pccZone := d.Get("pcc_zone").(string)
	result := make([]string, 0)
	err := config.OVHClient.Get(fmt.Sprintf("/dedicatedCloud/location/%s/hypervisor", pccZone), &result)
	if err != nil {
		return fmt.Errorf("Unable to retrieve /dedicatedCloud/location/{pccZone}/hypervisor information:\n\t %q", err)
	}

	sort.Strings(result)
	d.SetId(hashcode.Strings(result))
	d.Set("result", result)

	return nil
}

// /dedicatedCloud/location/{pccZone}/hypervisor/{id} 	# TODO

// Location stock host - reseller call

// /dedicatedCloud/location/{pccZone}/stock/host    	# TODO

// Location stock pcc - reseller call

// /dedicatedCloud/location/{pccZone}/stock/pcc			# TODO

// Location stock zpool - reseller call

// /dedicatedCloud/location/{pccZone}/stock/zpool		# TODO
