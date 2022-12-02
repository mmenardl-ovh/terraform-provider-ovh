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

func dataSourceDedicatedCloudAllowedNetworkTask() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedCloudAllowedNetworkTaskRead,
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network_access_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"task_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_from": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"datacenter_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"end_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"execution_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"filer_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"host_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"last_modification_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"maintenance_date_from": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"maintenance_date_to": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"network": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"order_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"parent_task_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"progress": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vlan_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceDedicatedCloudAllowedNetworkTaskRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	allowedNetworkId := d.Get("network_access_id").(string)
	taskId := d.Get("task_id").(string)
	task := &DedicatedCloudTask{}

	err := config.OVHClient.Get(
		fmt.Sprintf("/dedicatedCloud/%s/allowedNetwork/%s/task/%s", serviceName, allowedNetworkId, taskId),
		&task,
	)
	if err != nil {
		d.SetId("")
		return nil
	}

	d.SetId(fmt.Sprintf("%d", *task.TaskId))
	d.Set("created_by", task.CreatedBy)
	d.Set("created_from", task.CreatedFrom)
	d.Set("datacenter_id", task.DatacenterId)
	d.Set("description", task.Description)
	d.Set("end_date", task.EndDate)
	d.Set("execution_date", task.ExecutionDate)
	d.Set("filer_id", task.FilerId)
	d.Set("host_id", task.HostId)
	d.Set("last_modification_date", task.LastModificationDate)
	d.Set("maintenance_date_from", task.MaintenanceDateFrom)
	d.Set("maintenance_date_to", task.MaintenanceDateTo)
	d.Set("name", task.Name)
	d.Set("network", task.Network)
	d.Set("network_access_id", task.NetworkAccessId)
	d.Set("order_id", task.OrderId)
	d.Set("parent_task_id", task.ParentTaskId)
	d.Set("progress", task.Progress)
	d.Set("state", task.State)
	d.Set("task_id", task.TaskId)
	d.Set("type", task.Type)
	d.Set("user_id", task.UserId)
	d.Set("vlan_id", task.VlanId)

	return nil
}
