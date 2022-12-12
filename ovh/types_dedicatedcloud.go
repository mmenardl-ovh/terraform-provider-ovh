package ovh

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	// "github.com/ovh/terraform-provider-ovh/ovh/helpers"
)

// DedicatedCloud

type DedicatedCloud struct {
	AdvancedSecurity           *bool                  `json:"advancedSecurity"`
	Bandwidth                  *string                `json:"bandwidth"`
	BillingType                *string                `json:"billingType"`
	CertifiedInterfaceUrl      *string                `json:"certifiedInterfaceUrl"`
	CommercialRange            *string                `json:"commercialRange"`
	Description                *string                `json:"description"`
	Generation                 *string                `json:"generation"`
	Location                   *string                `json:"location"`
	ManagementInterface        *string                `json:"managementInterface"`
	ProductReference           *string                `json:"productReference"`
	ServiceName                *string                `json:"serviceName"`
	ServicePackName            *string                `json:"servicePackName"`
	Spla                       *bool                  `json:"spla"`
	SslV3                      *bool                  `json:"sslV3"`
	State                      *string                `json:"state"`
	UserAccessPolicy           *string                `json:"userAccessPolicy"`
	UserLimitConcurrentSession *int                   `json:"userLimitConcurrentSession"`
	UserLogoutPolicy           *string                `json:"userLogoutPolicy"`
	UserSessionTimeout         *int                   `json:"userSessionTimeout"`
	VScopeUrl                  *string                `json:"vScopeUrl"`
	Version                    *DedicatedCloudVersion `json:"version"`
	WebInterfaceUrl            *string                `json:"webInterfaceUrl"`
}

type DedicatedCloudVersion struct {
	Build *string `json:"build"`
	Major *string `json:"major"`
	Minor *string `json:"minor"`
}

// Commercial Range

type DedicatedCloudCommercialRange struct {
	DedicatedCloudVersion     *string   `json:"dedicatedCloudVersion"`
	AllowedNetworkRoles       *[]string `json:"allowedNetworkRoles"`
	CommercialRangeName       *string   `json:"commercialRangeName"`
	AllowedHypervisorVersions *[]string `json:"allowedHypervisorVersions"`
	Range                     *string   `json:"range"`
}

// Location

type DedicatedCloudLocation struct {
	Id          *int    `json:"id"`
	PccZone     *string `json:"pccZone"`
	City        *string `json:"city"`
	CountryCode *string `json:"countryCode"`
}

// Host Profile

type DedicatedCloudHostProfile struct {
	Id   *int    `json:"id"`
	Name *string `json:"name"`
}

// Os - Hypervisor

type DedicatedCloudOs struct {
	FullName             *string `json:"fullName"`
	LastModificationDate *string `json:"lastModificationDate"`
	ShortName            *string `json:"shortName"`
}

// Allowed Network

type DedicatedCloudAllowedNetwork struct {
	Description     *string `json:"description"`
	Network         *string `json:"network"`
	NetworkAccessId *int    `json:"networkAccessId"`
	State           *string `json:"state"`
}

// Task

type DedicatedCloudTask struct {
	CreatedBy            *string `json:"createdBy"`
	CreatedFrom          *string `json:"createdFrom"`
	DatacenterId         *int    `json:"datacenterId"`
	Description          *string `json:"description"`
	EndDate              *string `json:"endDate"`
	ExecutionDate        *string `json:"executionDate"`
	FilerId              *int    `json:"filerId"`
	HostId               *int    `json:"hostId"`
	LastModificationDate *string `json:"lastModificationDate"`
	MaintenanceDateFrom  *string `json:"maintenanceDateFrom"`
	MaintenanceDateTo    *string `json:"maintenanceDateTo"`
	Name                 *string `json:"name"`
	Network              *string `json:"network"`
	NetworkAccessId      *int    `json:"networkAccessId"`
	OrderId              *int    `json:"orderId"`
	ParentTaskId         *int    `json:"parentTaskId"`
	Progress             *int    `json:"progress"`
	State                *string `json:"state"`
	TaskId               *int    `json:"taskId"`
	Type                 *string `json:"type"`
	UserId               *int    `json:"userId"`
	VlanId               *int    `json:"vlanId"`
}

// Datacenter

type DedicatedCloudDatacenter struct {
	CommercialName      *string `json:"commercialName"`
	CommercialRangeName *string `json:"commercialRangeName"`
	DatacenterId        *int    `json:"datacenterId"`
	Description         *string `json:"description"`
	HorizonViewName     *string `json:"horizonViewName"`
	IsRemovable         *bool   `json:"isRemovable"`
	Name                *string `json:"name"`
	Version             *string `json:"version"`
}

// Cluster

type DedicatedCloudCluster struct {
	Id              *int                     `json:"id"`
	Autoscale       *DedicatedCloudAutoscale `json:"autoscale"`
	DrsMode         *string                  `json:"drsMode"`
	DrsStatus       *string                  `json:"drsStatus"`
	EvcMode         *string                  `json:"evcMode"`
	HaStatus        *string                  `json:"haStatus"`
	Name            *string                  `json:"name"`
	VmwareClusterId *string                  `json:"vmwareClusterId"`
}

type DedicatedCloudAutoscale struct {
	AutoScaleInHost     *string `json:"autoScaleInHost"`
	AutoScaleOutHost    *string `json:"autoScaleOutHost"`
	AutoScaleOutStorage *string `json:"autoScaleOutStorage"`
	ConfigId            *int    `json:"configId"`
	Id                  *int    `json:"id"`
	InMaintenanceMode   *bool   `json:"inMaintenanceMode"`
	State               *string `json:"state"`
}

// User

type DedicatedCloudUser struct {
	ActivationState      *string `json:"activationState"`
	ActiveDirectoryId    *int    `json:"activeDirectoryId"`
	ActiveDirectoryType  *string `json:"activeDirectoryType"`
	CanManageIpFailOvers *bool   `json:"canManageIpFailOvers"`
	CanManageNetwork     *bool   `json:"canManageNetwork"`
	CanManageRights      *bool   `json:"canManageRights"`
	Email                *string `json:"email"`
	EncryptionRight      *bool   `json:"encryptionRight"`
	FirstName            *string `json:"firstName"`
	FullAdminRo          *bool   `json:"fullAdminRo"`
	IsEnableManageable   *bool   `json:"isEnableManageable"`
	IsTokenValidator     *bool   `json:"isTokenValidator"`
	LastName             *string `json:"lastName"`
	Login                *string `json:"login"`
	Name                 *string `json:"name"`
	NsxRight             *bool   `json:"nsxRight"`
	PhoneNumber          *string `json:"phoneNumber"`
	ReceiveAlerts        *bool   `json:"receiveAlerts"`
	State                *string `json:"state"`
	UserId               *int    `json:"userId"`
}

func (v DedicatedCloudUser) ToMap() map[string]interface{} {
	obj := make(map[string]interface{})
	obj["login"] = v.Login
	obj["user_id"] = v.UserId
	if v.ActivationState != nil {
		obj["activation_state"] = v.ActivationState
	}
	if v.ActiveDirectoryId != nil {
		obj["activedirectory_id"] = v.ActiveDirectoryId
	}
	if v.ActiveDirectoryType != nil {
		obj["activedirectory_type"] = v.ActiveDirectoryType
	}
	obj["can_manage_ip_failovers"] = v.CanManageIpFailOvers
	obj["can_manage_rights"] = v.CanManageRights
	if v.Email != nil {
		obj["email"] = v.Email
	}
	obj["encryption_right"] = v.EncryptionRight
	if v.FirstName != nil {
		obj["first_name"] = v.FirstName
	}
	obj["full_admin_ro"] = v.FullAdminRo
	obj["is_enable_manageable"] = v.IsEnableManageable
	obj["is_token_validator"] = v.IsTokenValidator
	if v.LastName != nil {
		obj["last_name"] = v.LastName
	}
	obj["name"] = v.Name
	obj["nsx_right"] = v.NsxRight
	if v.PhoneNumber != nil {
		obj["phone_number"] = v.PhoneNumber
	}
	obj["receive_alerts"] = v.ReceiveAlerts
	obj["state"] = v.State

	return obj
}

type DedicatedCloudUserCreateOpts struct {
	ServiceName string `json:"serviceName"`
	// CanAddResource *bool `json:"canAddRessource` // ... And no, this is not a typo >_>
	// CanManageRights *bool `json:"canManageRights"`
	// Email *string `json:"email"`
	// EncryptionRight *bool `json:"encryptionRight"`
	// ExpirationDate *string `json:"expirationDate"`
	// FirstName *string `json:"firstName"`
	// LastName *string `json:"lastName"`
	Login string `json:"name"`
	// NetworkRole *string `json:"networkRole"`
	// NsxRight *bool `json:"nsxRight"`
	// Password *string `json:"password"`
	// PhoneNumber *string `json:"phoneNumber"`
	// ReceiveAlerts *bool `json:"receiveAlerts"`
	// Right *string `json:"right"`
	// TokenValidator *bool `json:"tokenValidator"`
	// VmNetworkRole *string `json:"vmNetworkRole"`
}

func (opts *DedicatedCloudUserCreateOpts) FromResource(d *schema.ResourceData) *DedicatedCloudUserCreateOpts {
	opts.ServiceName = d.Get("service_name").(string)
	opts.Login = d.Get("login").(string)

	return opts
}

// Filer

type DedicatedCloudFiler struct {
	ActiveNode        *string                  `json:"activeNode"`
	BillingType       *string                  `json:"billingType"`
	ConnectionState   *string                  `json:"connectionState"`
	FilerId           *int                     `json:"filerId"`
	FullProfile       *string                  `json:"fullProfile"`
	IsManagedByOvh    *bool                    `json:"isManagedByOvh"`
	Master            *string                  `json:"master"`
	Name              *string                  `json:"name"`
	Profile           *string                  `json:"profile"`
	ResourceName      *string                  `json:"resourceName"`
	Size              *DedicatedCloudFilerSize `json:"size"`
	Slave             *string                  `json:"slave"`
	SpaceFree         *int                     `json:"spaceFree"`
	SpaceProvisionned *int                     `json:"spaceProvisionned"`
	SpaceUsed         *int                     `json:"spaceUsed"`
	State             *string                  `json:"state"`
	VmTotal           *int                     `json:"vmTotal"`
}

type DedicatedCloudFilerSize struct {
	Unit  *string `json:"unit"`
	Value *int    `json:"value"`
}
