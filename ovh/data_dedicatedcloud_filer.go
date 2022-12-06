package ovh

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ovh/terraform-provider-ovh/ovh/helpers/hashcode"
)

// Filers

func dataSourceDedicatedCloudFilers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedCloudFilersRead,
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"filers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
		},
	}
}

func dataSourceDedicatedCloudFilersRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	result := make([]int, 0)

	endpoint := fmt.Sprintf("/dedicatedCloud/%s/filer", serviceName)
	err := config.OVHClient.Get(endpoint, &result)
	if err != nil {
		return fmt.Errorf("Error calling GET %s:\n\t %q", endpoint, err)
	}

	sort.Ints(result)
	var stringResults []string
	for _, i := range result {
		stringResults = append(stringResults, strconv.Itoa(i))
	}
	d.SetId(hashcode.Strings(stringResults))
	d.Set("filers", result)

	return nil
}

// Filer

func dataSourceDedicatedCloudFiler() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedCloudFilerRead,
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"filer_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"active_node": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"billing_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"connection_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"full_profile": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_managed_by_ovh": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"master": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"profile": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
			},
			"slave": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"space_free": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"space_provisionned": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"space_used": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vm_total": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceDedicatedCloudFilerRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	filerId := d.Get("filer_id").(int)
	filer := &DedicatedCloudFiler{}

	endpoint := fmt.Sprintf("/dedicatedCloud/%s/filer/%d", serviceName, filerId)
	err := config.OVHClient.Get(endpoint, &filer)
	if err != nil {
		return fmt.Errorf("Error calling GET %s:\n\t %q", endpoint, err)
	}

	d.SetId(fmt.Sprintf("%s/%d", d.Get("service_name"), d.Get("filer_id")))
	d.Set("active_node", *filer.ActiveNode)
	d.Set("billing_type", *filer.BillingType)
	d.Set("connection_state", *filer.ConnectionState)
	d.Set("filer_id", d.Get("filer_id"))
	d.Set("full_profile", *filer.FullProfile)
	d.Set("is_managed_by_ovh", *filer.IsManagedByOvh)
	d.Set("master", *filer.Master)
	d.Set("name", *filer.Name)
	d.Set("profile", *filer.Profile)
	if filer.ResourceName != nil {
		d.Set("resource_name", *filer.ResourceName)
	}
	if filer.Size != nil {
		var size = make(map[string]interface{})
		size["unit"] = *filer.Size.Unit
		size["value"] = strconv.Itoa(*filer.Size.Value)
		d.Set("size", size)
	}
	if filer.Slave != nil {
		d.Set("slave", *filer.Slave)
	}
	d.Set("space_free", *filer.SpaceFree)
	d.Set("space_provisionned", *filer.SpaceProvisionned)
	d.Set("space_used", *filer.SpaceUsed)
	d.Set("state", *filer.State)
	d.Set("vm_total", *filer.VmTotal)

	return nil
}
