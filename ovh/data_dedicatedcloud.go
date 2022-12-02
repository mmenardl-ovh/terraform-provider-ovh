package ovh

import (
	"fmt"
	"net/url"
	"sort"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ovh/terraform-provider-ovh/ovh/helpers/hashcode"
)

func dataSourceDedicatedClouds() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedCloudsRead,
		Schema: map[string]*schema.Schema{
			"result": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceDedicatedCloudsRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	result := make([]string, 0)
	err := config.OVHClient.Get("/dedicatedCloud", &result)
	if err != nil {
		return fmt.Errorf("Unable to retrieve /dedicatedCloud information:\n\t %q", err)
	}

	sort.Strings(result)
	d.SetId(hashcode.Strings(result))
	d.Set("result", result)
	return nil
}

func dataSourceDedicatedCloud() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedCloudRead,
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"advanced_security": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"bandwidth": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"billing_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"certified_interface_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"commercialrange": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"generation": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"location": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_interface": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"product_reference": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"servicepack_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"spla": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"sslv3": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_access_policy": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_limit_concurrent_session": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"user_logout_policy": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_session_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vscope_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
			},
			"web_interface_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceDedicatedCloudRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	serviceName := d.Get("service_name").(string)

	dedicatedCloud := &DedicatedCloud{}

	err := config.OVHClient.Get(
		fmt.Sprintf("/dedicatedCloud/%s", url.PathEscape(serviceName)),
		&dedicatedCloud,
	)

	if err != nil {
		d.SetId("")
		return nil
	}

	d.SetId(*dedicatedCloud.ServiceName)
	d.Set("advanced_security", *dedicatedCloud.AdvancedSecurity)
	d.Set("bandwidth", *dedicatedCloud.Bandwidth)
	d.Set("billing_type", *dedicatedCloud.BillingType)
	d.Set("certified_interface_url", *dedicatedCloud.CertifiedInterfaceUrl)
	d.Set("commercialrange", *dedicatedCloud.CommercialRange)
	d.Set("description", *dedicatedCloud.Description)
	d.Set("generation", *dedicatedCloud.Generation)
	d.Set("location", *dedicatedCloud.Location)
	d.Set("management_interface", *dedicatedCloud.ManagementInterface)
	d.Set("product_reference", *dedicatedCloud.ProductReference)
	d.Set("service_name", *dedicatedCloud.ServiceName)
	d.Set("servicepack_name", *dedicatedCloud.ServicePackName)
	d.Set("spla", *dedicatedCloud.Spla)
	d.Set("sslv3", *dedicatedCloud.SslV3)
	d.Set("state", *dedicatedCloud.State)
	d.Set("user_access_policy", *dedicatedCloud.UserAccessPolicy)
	d.Set("user_limit_concurrent_session", *dedicatedCloud.UserLimitConcurrentSession)
	d.Set("user_logout_policy", *dedicatedCloud.UserLogoutPolicy)
	d.Set("user_session_timeout", *dedicatedCloud.UserSessionTimeout)
	d.Set("vscope_url", *dedicatedCloud.VScopeUrl)
	var version = make(map[string]string)
	version["build"] = *dedicatedCloud.Version.Build
	version["major"] = *dedicatedCloud.Version.Major
	version["minor"] = *dedicatedCloud.Version.Minor
	d.Set("version", version)
	d.Set("web_interface_url", *dedicatedCloud.WebInterfaceUrl)

	return nil
}
