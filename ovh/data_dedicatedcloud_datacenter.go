package ovh

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ovh/terraform-provider-ovh/ovh/helpers/hashcode"
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
	var stringResults []string
	for _, i := range result {
		stringResults = append(stringResults, strconv.Itoa(i))
	}
	d.SetId(hashcode.Strings(stringResults))
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
	var stringResults []string
	for _, i := range result {
		stringResults = append(stringResults, strconv.Itoa(i))
	}
	d.SetId(hashcode.Strings(stringResults))
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

	d.SetId(fmt.Sprintf("%s/%d/%d", d.Get("service_name"), d.Get("datacenter_id"), d.Get("task_id")))
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

// Datacenter clusters

func dataSourceDedicatedCloudDatacenterClusters() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedCloudDatacenterClustersRead,
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"datacenter_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"clusters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
		},
	}
}

func dataSourceDedicatedCloudDatacenterClustersRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	datacenterId := d.Get("datacenter_id").(int)
	result := make([]int, 0)

	endpoint := fmt.Sprintf("/dedicatedCloud/%s/datacenter/%d/cluster", serviceName, datacenterId)
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
	d.Set("clusters", result)

	return nil
}

// Datacenter Cluster

func dataSourceDedicatedCloudDatacenterCluster() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedCloudDatacenterClusterRead,
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"datacenter_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"cluster_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"autoscale": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"drs_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"drs_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"evc_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ha_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vmware_cluster_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceDedicatedCloudDatacenterClusterRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	datacenterId := d.Get("datacenter_id").(int)
	clusterId := d.Get("cluster_id").(int)
	cluster := &DedicatedCloudCluster{}

	endpoint := fmt.Sprintf("/dedicatedCloud/%s/datacenter/%d/cluster/%d", serviceName, datacenterId, clusterId)
	err := config.OVHClient.Get(endpoint, &cluster)
	if err != nil {
		return fmt.Errorf("Error calling GET %s:\n\t %q", endpoint, err)
	}

	d.SetId(fmt.Sprintf("%s/%d/%d", d.Get("service_name"), d.Get("datacenter_id"), d.Get("cluster_id")))
	d.Set("drs_mode", *cluster.DrsMode)
	d.Set("drs_status", *cluster.DrsStatus)
	d.Set("evc_mode", *cluster.EvcMode)
	d.Set("ha_status", *cluster.HaStatus)
	d.Set("name", *cluster.Name)
	d.Set("vmware_cluster_id", *cluster.VmwareClusterId)
	if cluster.Autoscale != nil {
		var autoscale = make(map[string]interface{})
		autoscale["autoscale_in_host"] = *cluster.Autoscale.AutoScaleInHost
		autoscale["autoscale_out_host"] = *cluster.Autoscale.AutoScaleOutHost
		autoscale["autoscale_out_storage"] = *cluster.Autoscale.AutoScaleOutStorage
		autoscale["config_id"] = strconv.Itoa(*cluster.Autoscale.ConfigId)
		autoscale["id"] = strconv.Itoa(*cluster.Autoscale.Id)
		autoscale["in_maintenance_mode"] = *cluster.Autoscale.InMaintenanceMode
		autoscale["state"] = *cluster.Autoscale.State

		d.Set("autoscale", autoscale)
	}

	return nil
}

// Datacenter Filers

func dataSourceDedicatedCloudDatacenterFilers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedCloudDatacenterFilersRead,
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"datacenter_id": {
				Type:     schema.TypeInt,
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

func dataSourceDedicatedCloudDatacenterFilersRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	datacenterId := d.Get("datacenter_id").(int)
	result := make([]int, 0)

	endpoint := fmt.Sprintf("/dedicatedCloud/%s/datacenter/%d/filer", serviceName, datacenterId)
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

// Datacenter Filer

func dataSourceDedicatedCloudDatacenterFiler() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedCloudDatacenterFilerRead,
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"datacenter_id": {
				Type:     schema.TypeInt,
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

func dataSourceDedicatedCloudDatacenterFilerRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	datacenterId := d.Get("datacenter_id").(int)
	filerId := d.Get("filer_id").(int)
	filer := &DedicatedCloudFiler{}

	endpoint := fmt.Sprintf("/dedicatedCloud/%s/datacenter/%d/filer/%d", serviceName, datacenterId, filerId)
	err := config.OVHClient.Get(endpoint, &filer)
	if err != nil {
		return fmt.Errorf("Error calling GET %s:\n\t %q", endpoint, err)
	}

	d.SetId(fmt.Sprintf("%s/%d/%d", d.Get("service_name"), d.Get("datacenter_id"), d.Get("filer_id")))
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
