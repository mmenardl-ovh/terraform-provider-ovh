package ovh

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

type DedicatedCloudCommercialRange struct {
	DedicatedCloudVersion     *string   `json:"dedicatedCloudVersion"`
	AllowedNetworkRoles       *[]string `json:"allowedNetworkRoles"`
	CommercialRangeName       *string   `json:"commercialRangeName"`
	AllowedHypervisorVersions *[]string `json:"allowedHypervisorVersions"`
	Range                     *string   `json:"range"`
}

type DedicatedCloudLocation struct {
	Id          *int    `json:"id"`
	PccZone     *string `json:"pccZone"`
	City        *string `json:"city"`
	CountryCode *string `json:"countryCode"`
}

type DedicatedCloudHostProfile struct {
	Id   *int    `json:"id"`
	Name *string `json:"name"`
}

type DedicatedCloudOs struct {
	FullName             *string `json:"fullName"`
	LastModificationDate *string `json:"lastModificationDate"`
	ShortName            *string `json:"shortName"`
}
