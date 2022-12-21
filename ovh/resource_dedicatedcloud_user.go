package ovh

import (
	"context"
	"fmt"
	"log"
	"time"

	"net/url"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDedicatedCloudUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDedicatedCloudUserCreate,
		ReadContext:   resourceDedicatedCloudUserRead,
		UpdateContext: resourceDedicatedCloudUserUpdate,
		DeleteContext: resourceDedicatedCloudUserDelete,
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
				Computed:    true,
			},
			"can_manage_rights": {
				Type:        schema.TypeBool,
				Description: "Defines if the user can manage users rights",
				ForceNew:    false,
				Optional:    true,
				Computed:    true,
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
				Computed:    true,
			},
			"first_name": {
				Type:        schema.TypeString,
				Description: "First name of the user",
				ForceNew:    false,
				Optional:    true,
				Computed:    true,
			},
			"full_admin_ro": {
				Type:        schema.TypeBool,
				Description: "Defines if the user is a full admin in readonly",
				ForceNew:    false,
				Optional:    true,
				Computed:    true,
			},
			"is_token_validator": {
				Type:        schema.TypeBool,
				Description: "Defines if the user can confirm security tokens (if a compatible option is enabled)",
				ForceNew:    false,
				Optional:    true,
				Computed:    true,
			},
			"last_name": {
				Type:        schema.TypeString,
				Description: "Last name of the user",
				ForceNew:    false,
				Optional:    true,
				Computed:    true,
			},
			"last_updated": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nsx_right": {
				Type:        schema.TypeBool,
				Description: "Is this User able to access nsx interface (requires NSX option)",
				ForceNew:    false,
				Optional:    true,
				Computed:    true,
			},
			"phone_number": {
				Type:        schema.TypeString,
				Description: "Mobile phone number of the user",
				ForceNew:    false,
				Optional:    true,
				Computed:    true,
			},
			"receive_alerts": {
				Type:        schema.TypeBool,
				Description: "Defines if the user receives technical alerts",
				ForceNew:    false,
				Optional:    true,
				Computed:    true,
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
			"can_manage_network": {
				Type:        schema.TypeBool,
				Description: "Defines if the user can manage the network",
				ForceNew:    false,
				Computed:    true,
			},
			"is_enable_manageable": {
				Type:        schema.TypeBool,
				Description: "Check if the given Dedicated Cloud user can be enabled or disabled ?",
				Computed:    true,
			},
			"state": {
				Type:        schema.TypeString,
				Description: "State of the user account",
				Computed:    true,
			},
			"user_id": {
				Type:        schema.TypeInt,
				Description: "User ID",
				Computed:    true,
			},
		},
	}
}

func resourceDedicatedCloudUserCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	userLogin := d.Get("login").(string)

	opts := (&DedicatedCloudUserCreateOpts{}).FromResource(d)
	task := &DedicatedCloudTask{}

	log.Printf("[DEBUG][Create] DedicatedCloudTask addUser for %s/%s", serviceName, userLogin)
	endpoint := fmt.Sprintf("/dedicatedCloud/%s/user", url.PathEscape(serviceName))
	if err := config.OVHClient.Post(endpoint, opts, &task); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[DEBUG][Create][WaitForDone] DedicatedCloudTask addUser for %s/%s", serviceName, userLogin)
	if err := waitForDedicatedCloudTask(6*time.Minute, serviceName, task, config.OVHClient); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%s/%s", serviceName, userLogin))
	return resourceDedicatedCloudUserRead(ctx, d, meta)
}

func resourceDedicatedCloudUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	userLogin := d.Get("login").(string)

	var diags diag.Diagnostics

	user, err := getDedicatedCloudUser(serviceName, userLogin, config.OVHClient)
	if err != nil {
		d.SetId("")
		return diag.FromErr(err)
	}

	log.Printf("[INFO] Found DedicatedCloud user %s/%s with %d", serviceName, *user.Login, *user.UserId)
	d.SetId(fmt.Sprintf("%s/%s", serviceName, *user.Login))
	d.Set("service_name", serviceName)
	for k, v := range user.ToMap() {
		if k != "name" {
			d.Set(k, v)
		}
	}

	return diags
}

func resourceDedicatedCloudUserUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	userId := d.Get("user_id").(int)

	// var diags diag.Diagnostics
	optionalAttributes := []string{"can_manage_rights", "encryption_right", "first_name", "is_token_validator", "last_name", "receive_alerts"}
	if d.HasChanges(optionalAttributes...) {
		opts := (&DedicatedCloudUserUpdateOpts{}).FromResource(d)
		task := &DedicatedCloudTask{}

		endpoint := fmt.Sprintf("/dedicatedCloud/%s/user/%d/changeProperties", serviceName, userId)
		if err := config.OVHClient.Post(endpoint, opts, &task); err != nil {
			return diag.FromErr(err)
		}

		if err := waitForDedicatedCloudTask(10*time.Minute, serviceName, task, config.OVHClient); err != nil {
			return diag.FromErr(err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC3339))
	}

	return resourceDedicatedCloudUserRead(ctx, d, meta)
}

func resourceDedicatedCloudUserDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	userLogin := d.Get("login").(string)

	var diags diag.Diagnostics

	user, err := getDedicatedCloudUser(serviceName, userLogin, config.OVHClient)
	if err != nil {
		d.SetId("")
		return diag.FromErr(err)
	}

	// Create delete task
	task := &DedicatedCloudTask{}

	log.Printf("[DEBUG][Delete] Create DedicatedCloudTask removeUser for %s/%s (%d)", serviceName, userLogin, *user.UserId)
	endpoint := fmt.Sprintf("/dedicatedCloud/%s/user/%d", url.PathEscape(serviceName), *user.UserId)
	if err := config.OVHClient.Delete(endpoint, task); err != nil {
		return diag.FromErr(err)
	}

	// Wait for delete task
	log.Printf("[DEBUG][Create][WaitForDone] DedicatedCloudTask removeUser %d to transit to state done.", *task.TaskId)
	if err := waitForDedicatedCloudTask(20*time.Minute, serviceName, task, config.OVHClient); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return diags
}
