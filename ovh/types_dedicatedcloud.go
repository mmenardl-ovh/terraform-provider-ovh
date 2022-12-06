package ovh

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
