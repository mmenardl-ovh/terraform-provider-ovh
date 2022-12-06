package ovh

import (
	"fmt"
	"net/url"
	"sort"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ovh/terraform-provider-ovh/ovh/helpers/hashcode"
)

// Users

func dataSourceDedicatedCloudUsers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedCloudUsersRead,
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"users": {
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

	endpoint := fmt.Sprintf("/dedicatedCloud/%s/user", url.PathEscape(serviceName))

	result := make([]int, 0)
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
	d.Set("users", result)

	return nil
}

// User

func dataSourceDedicatedCloudUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedCloudUserRead,
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"activation_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"activedirectory_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"activedirectory_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"can_manage_ip_failovers": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"can_manage_rights": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"encryption_right": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"first_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"full_admin_ro": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_enable_manageable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_token_validator": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"last_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"login": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nsx_right": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"phone_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"receive_alerts": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceDedicatedCloudUserRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	userId := d.Get("user_id").(int)
	user := &DedicatedCloudUser{}

	endpoint := fmt.Sprintf("/dedicatedCloud/%s/user/%d", serviceName, userId)
	err := config.OVHClient.Get(endpoint, &user)
	if err != nil {
		return fmt.Errorf("Error calling GET %s:\n\t %q", endpoint, err)
	}

	d.SetId(fmt.Sprintf("%s/%d", d.Get("service_name").(string), *user.UserId))
	d.Set("activation_state", *user.ActivationState)
	d.Set("activedirectory_id", *user.ActiveDirectoryId)
	d.Set("activedirectory_type", *user.ActiveDirectoryType)
	d.Set("can_manage_ip_failovers", *user.CanManageIpFailOvers)
	d.Set("can_manage_network", *user.CanManageNetwork)
	d.Set("can_manage_rights", *user.CanManageRights)
	d.Set("email", *user.Email)
	d.Set("encryption_right", *user.EncryptionRight)
	if user.FirstName != nil {
		d.Set("first_name", *user.FirstName)
	}
	d.Set("full_admin_ro", *user.FullAdminRo)
	d.Set("is_enable_manageable", *user.IsEnableManageable)
	d.Set("is_token_validator", *user.IsTokenValidator)
	if user.LastName != nil {
		d.Set("last_name", *user.LastName)
	}
	d.Set("login", *user.Login)
	d.Set("name", *user.Name)
	d.Set("nsx_right", *user.NsxRight)
	if user.PhoneNumber != nil {
		d.Set("phone_number", *user.PhoneNumber)
	}
	d.Set("receive_alerts", *user.ReceiveAlerts)
	d.Set("state", *user.State)
	d.Set("user_id", d.Get("user_id"))

	return nil
}

// User tasks

func dataSourceDedicatedCloudUserTasks() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedCloudUserTasksRead,
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_id": {
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

func dataSourceDedicatedCloudUserTasksRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	userId := d.Get("user_id").(int)
	result := make([]int, 0)

	endpoint := fmt.Sprintf("/dedicatedCloud/%s/user/%d/task", serviceName, userId)
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

// User task

func dataSourceDedicatedCloudUserTask() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDedicatedCloudUserTaskRead,
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_id": {
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
			"datacenter_id": {
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

func dataSourceDedicatedCloudUserTaskRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	userId := d.Get("user_id").(int)
	taskId := d.Get("task_id").(int)
	task := &DedicatedCloudTask{}

	endpoint := fmt.Sprintf("/dedicatedCloud/%s/user/%d/task/%d", serviceName, userId, taskId)
	err := config.OVHClient.Get(endpoint, &task)
	if err != nil {
		return fmt.Errorf("Error calling GET %s:\n\t %q", endpoint, err)
	}

	d.SetId(fmt.Sprintf("%s/%d/%d", d.Get("service_name"), d.Get("user_id"), d.Get("task_id")))
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
