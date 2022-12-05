package ovh

import (
	"fmt"
	"net/url"
	"sort"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	// "github.com/ovh/terraform-provider-ovh/ovh/helpers/hashcode"
)

func dataSourceDedicatedCloudUsers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedCloudUsersRead,
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"result": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func dataSourceDedicatedCloudUsersRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)

	result := make([]int, 0)
	err := config.OVHClient.Get(
		fmt.Sprintf("/dedicatedCloud/%s/user", url.PathEscape(serviceName)),
		&result,
	)
	if err != nil {
		return fmt.Errorf("Unable to retrieve /dedicatedCloud/{serviceName} information:\n\t %q", err)
	}

	sort.Ints(result)
	d.SetId(serviceName)
	d.Set("result", result)

	return nil
}
