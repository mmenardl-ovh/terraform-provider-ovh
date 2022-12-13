package ovh

import (
	"fmt"
	"log"
	"time"

	// "log"
	"net/url"
	"strconv"

	//  "strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ovh/terraform-provider-ovh/ovh/helpers"
)

func resourceDedicatedCloudUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceDedicatedCloudUserCreate,
		Read:   resourceDedicatedCloudUserRead,
		Update: resourceDedicatedCloudUserRead,
		Delete: resourceDedicatedCloudUserDelete,
		// Importer: &schema.ResourceImporter{
		// 	State: resourceDedicatedCloudUserImportState,
		// },
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:        schema.TypeString,
				Description: "Service name",
				ForceNew:    true,
				Required:    true,
			},
			"login": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Description: "Login of the user",
				Required:    true,
			},

			// Optional
			"can_manage_ip_failovers": {
				Type:        schema.TypeBool,
				Description: "Defines if the user can manage ip failovers",
				ForceNew:    false,
				Optional:    true,
			},
			"can_manage_rights": {
				Type:        schema.TypeBool,
				Description: "Defines if the user can manage users rights",
				ForceNew:    false,
				Optional:    true,
			},
			"email": {
				Type:        schema.TypeString,
				Description: "Email address of the user",
				ForceNew:    false,
				Optional:    true,
			},
			"encryption_right": {
				Type:        schema.TypeBool,
				Description: "Defines if the user can manage encryption / KMS configuration",
				ForceNew:    false,
				Optional:    true,
			},
			"first_name": {
				Type:        schema.TypeString,
				Description: "First name of the user",
				ForceNew:    false,
				Optional:    true,
			},
			"is_enable_manageable": {
				Type:        schema.TypeBool,
				Description: "Check if the given Dedicated Cloud user can be enabled or disabled ?",
				ForceNew:    false,
				Optional:    true,
			},
			"is_token_validator": {
				Type:        schema.TypeBool,
				Description: "Defines if the user can confirm security tokens (if a compatible option is enabled)",
				ForceNew:    false,
				Optional:    true,
			},
			"last_name": {
				Type:        schema.TypeString,
				Description: "Last name of the user",
				ForceNew:    false,
				Optional:    true,
			},

			// Computed
			"activation_state": {
				Type:        schema.TypeString,
				Description: "Activation state of the user account",
				Computed:    true,
			},
			"activedirectory_id": {
				Type:        schema.TypeInt,
				Description: "Linked Federation Active Directory (if any)",
				Computed:    true,
			},
			"activedirectory_type": {
				Type:        schema.TypeString,
				Description: "Federation Active Directory user type (if any)",
				Computed:    true,
			},
			"full_admin_ro": {
				Type:        schema.TypeBool,
				Description: "Defines if the user is a full admin in readonly",
				Computed:    true,
			},
			"state": {
				Type:        schema.TypeString,
				Description: "State of the user account",
				Computed:    true,
			},
			"user_id": {
				Type:        schema.TypeString,
				Description: "User name",
				Computed:    true,
			},
		},
	}
}

func resourceDedicatedCloudUserCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	userLogin := d.Get("login").(string)

	opts := (&DedicatedCloudUserCreateOpts{}).FromResource(d)
	task := &DedicatedCloudTask{}

	log.Printf("[DEBUG][Create] DedicatedCloudTask (for user)")
	endpoint := fmt.Sprintf("/dedicatedCloud/%s/user", url.PathEscape(serviceName))
	err := config.OVHClient.Post(endpoint, opts, &task)
	if err != nil {
		return fmt.Errorf("failed to create DedicatedCloud user: %s", err)
	}

	log.Printf("[DEBUG][Create][WaitForArchived] DedicatedCloudTask (for user)")
	endpoint = fmt.Sprintf("/dedicatedCloud/%s/task/%s", url.PathEscape(serviceName), strconv.Itoa(*task.TaskId))
	err = WaitArchivedHostingPrivateDabaseTask(config.OVHClient, endpoint, 6*time.Minute)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s", serviceName, userLogin))
	return resourceDedicatedCloudUserRead(d, meta)
}

func resourceDedicatedCloudUserRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	userLogin := d.Get("login").(string)

	ds := &DedicatedCloudUser{}

	// lookup user's id
	lookupEndpoint := fmt.Sprintf("/dedicatedCloud/%s/user&name=%s", url.PathEscape(serviceName), url.PathEscape(userLogin))
	lookupResult := make([]int, 0)
	if err := config.OVHClient.Get(lookupEndpoint, &lookupResult); err != nil || len(lookupResult) != 1 {
		d.SetId("")
		return nil
	}

	// lookup user detail
	endpoint := fmt.Sprintf("/dedicatedCloud/%s/user/%s", url.PathEscape(serviceName), url.PathEscape(strconv.Itoa(lookupResult[0])))
	if err := config.OVHClient.Get(endpoint, &ds); err != nil {
		return helpers.CheckDeleted(d, err, endpoint)
	}

	d.SetId(fmt.Sprintf("%s/%s", serviceName, userLogin))
	d.Set("service_name", serviceName)
	for k, v := range ds.ToMap() {
		d.Set(k, v)
	}

	return nil
}

func resourceDedicatedCloudUserDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	userLogin := d.Get("login").(string)

	// Resolve login --> userId
	endpoint := fmt.Sprintf("/dedicatedCloud/%s/user&name=%s", url.PathEscape(serviceName), url.PathEscape(userLogin))
	lookupResult := make([]int, 0)
	if err := config.OVHClient.Get(endpoint, &lookupResult); err != nil || len(lookupResult) != 1 {
		d.SetId("")
		return nil
	}

	// Create delete task
	task := &DedicatedCloudTask{}

	log.Printf("[DEBUG][Delete] DedicatedCloudTask (for user)")
	endpoint = fmt.Sprintf("/dedicatedCloud/%s/user/%d", url.PathEscape(serviceName), lookupResult[0])
	if err := config.OVHClient.Delete(endpoint, task); err != nil {
		helpers.CheckDeleted(d, err, endpoint)
	}

	// Wait for delete task
	log.Printf("[DEBUG][Delete][WaitForArchived] DedicatedCloudTask (for user)")
	endpoint = fmt.Sprintf("/dedicatedCloud/%s/task/%s", url.PathEscape(serviceName), strconv.Itoa(*task.TaskId))
	err := WaitArchivedHostingPrivateDabaseTask(config.OVHClient, endpoint, 6*time.Minute)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
