package ovh

import (
	"fmt"
	"math/rand"
	"sort"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Allowed networks

func dataSourceDedicatedCloudAllowedNetworks() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedCloudAllowedNetworksRead,
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"result": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
		},
	}
}

func dataSourceDedicatedCloudAllowedNetworksRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	result := make([]int, 0)

	err := config.OVHClient.Get(fmt.Sprintf("/dedicatedCloud/%s/allowedNetwork", serviceName), &result)
	if err != nil {
		return fmt.Errorf("Unable to retrieve /dedicatedCloud/{serviceName}/allowedNetwork information:\n\t %q", err)
	}

	sort.Ints(result)
	d.SetId(fmt.Sprintf("%d", rand.Intn(1000000)))
	d.Set("result", result)

	return nil
}

// Allowed network

func dataSourceDedicatedCloudAllowedNetwork() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedCloudAllowedNetworkRead,
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network_access_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"network": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceDedicatedCloudAllowedNetworkRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	allowedNetworkId := d.Get("network_access_id").(string)
	allowedNetwork := &DedicatedCloudAllowedNetwork{}

	err := config.OVHClient.Get(
		fmt.Sprintf("/dedicatedCloud/%s/allowedNetwork/%s", serviceName, allowedNetworkId),
		&allowedNetwork,
	)
	if err != nil {
		d.SetId("")
		return nil
	}

	d.SetId(fmt.Sprintf("%d", *allowedNetwork.NetworkAccessId))
	d.Set("description", *allowedNetwork.Description)
	d.Set("network", *allowedNetwork.Network)
	d.Set("network_access_id", *allowedNetwork.NetworkAccessId)
	d.Set("state", *allowedNetwork.State)

	return nil
}

// Allowed network tasks

func dataSourceDedicatedCloudAllowedNetworkTasks() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedCloudAllowedNetworksRead,
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"result": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
		},
	}
}

func dataSourceDedicatedCloudAllowedNetworkTasksRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	allowedNetworkId := d.Get("network_access_id").(string)
	result := make([]int, 0)

	err := config.OVHClient.Get(fmt.Sprintf("/dedicatedCloud/%s/allowedNetwork/%s/task", serviceName, allowedNetworkId), &result)
	if err != nil {
		return fmt.Errorf("Unable to retrieve /dedicatedCloud/{serviceName}/allowedNetwork/{network_access_id}/task information:\n\t %q", err)
	}

	sort.Ints(result)
	d.SetId(fmt.Sprintf("%d", rand.Intn(1000000)))
	d.Set("result", result)

	return nil
}

// Allowed network task
