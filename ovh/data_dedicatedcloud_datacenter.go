package ovh

import (
	"fmt"
	"math/rand"
	"sort"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Datacenters

func dataSourceDedicatedCloudDatacenters() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedCloudDatacentersRead,
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"datacenters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
		},
	}
}

func dataSourceDedicatedCloudDatacentersRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	result := make([]int, 0)

	endpoint := fmt.Sprintf("/dedicatedCloud/%s/datacenter", serviceName)
	err := config.OVHClient.Get(endpoint, &result)
	if err != nil {
		return fmt.Errorf("Error calling GET %s:\n\t %q", endpoint, err)
	}

	sort.Ints(result)
	d.SetId(fmt.Sprintf("%d", rand.Intn(1000000)))
	d.Set("datacenters", result)

	return nil
}

// Datacenter

func dataSourceDedicatedCloudDatacenter() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedCloudDatacenterRead,
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"datacenter_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"commercial_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"commercialrange_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"horizon_view_name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				Default:  nil,
			},
			"is_removable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceDedicatedCloudDatacenterRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	datacenterId := d.Get("datacenter_id").(int)
	datacenter := &DedicatedCloudDatacenter{}

	endpoint := fmt.Sprintf("/dedicatedCloud/%s/datacenter/%d", serviceName, datacenterId)
	err := config.OVHClient.Get(endpoint, &datacenter)
	if err != nil {
		return fmt.Errorf("Error calling GET %s:\n\t %q", endpoint, err)
	}

	d.SetId(fmt.Sprintf("%s/%d", d.Get("service_name").(string), *datacenter.DatacenterId))
	d.Set("datacenter_id", d.Get("datacenter_id").(int))
	d.Set("commercial_name", *datacenter.CommercialName)
	d.Set("commercialrange_name", *datacenter.CommercialRangeName)
	d.Set("description", *datacenter.Description)
	if datacenter.HorizonViewName != nil {
		d.Set("horizon_view_name", *datacenter.HorizonViewName)
	}
	d.Set("name", *datacenter.Name)
	d.Set("is_removable", *datacenter.IsRemovable)
	d.Set("version", *datacenter.Version)

	return nil
}

// Datacenter tasks

func dataSourceDedicatedCloudDatacenterTasks() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedCloudDatacenterTasksRead,
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"datacenter_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"tasks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
		},
	}
}

func dataSourceDedicatedCloudDatacenterTasksRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	datacenterId := d.Get("datacenter_id").(int)
	result := make([]int, 0)

	endpoint := fmt.Sprintf("/dedicatedCloud/%s/datacenter/%d/task", serviceName, datacenterId)
	err := config.OVHClient.Get(endpoint, &result)
	if err != nil {
		return fmt.Errorf("Error calling GET %s:\n\t %q", endpoint, err)
	}

	sort.Ints(result)
	d.SetId(fmt.Sprintf("%d", rand.Intn(1000000)))
	d.Set("tasks", result)

	return nil
}

// Datacenter task

func dataSourceDedicatedCloudDatacenterTask() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedCloudDatacenterTaskRead,
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"datacenter_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"task_id": {
				Type:     schema.TypeInt,
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
			"network_access_id": {
				Type:     schema.TypeInt,
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

func dataSourceDedicatedCloudDatacenterTaskRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	datacenterId := d.Get("datacenter_id").(int)
	taskId := d.Get("task_id").(int)
	task := &DedicatedCloudTask{}

	endpoint := fmt.Sprintf("/dedicatedCloud/%s/datacenter/%d/task/%d", serviceName, datacenterId, taskId)
	err := config.OVHClient.Get(endpoint, &task)
	if err != nil {
		return fmt.Errorf("Error calling GET %s:\n\t %q", endpoint, err)
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
