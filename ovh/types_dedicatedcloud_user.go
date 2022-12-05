package ovh

import "fmt"

const (
	Disabled ActivationStateType = iota
	Disabling
	Enabled
	Enabling
	ToDisable
	ToEnable
)

const (
	Creating StateType = iota
	Deleting
	Delivered
	Error
)

type ActivationStateType int

type StateType int

type DedicatedCloudUser struct {
	ActivationState      *ActivationStateType `json:"activationState"`
	ActiveDirectoryId    *int                 `json:"activeDirectoryId"`
	ActiveDirectoryType  *string              `json:"activeDirectoryType"`
	CanManageIpFailOvers *bool                `json:"canManageIpFailOvers"`
	CanManageNetwork     *bool                `json:"canManageNetwork"`
	CanManageRights      *bool                `json:"canManageRights"`
	Email                *string              `json:"email"`
	EncryptionRight      *bool                `json:"encryptionRight"`
	FirstName            *bool                `json:"firstName"`
	FullAdminRo          *bool                `json:"fullAdminRo"`
	IsEnableManageable   *bool                `json:"isEnableManageable"`
	IsTokenValidator     *bool                `json:"isTokenValidator"`
	LastName             *string              `json:"lastName"`
	Login                *string              `json:"login"`
	Name                 *string              `json:"name"`
	NsxRight             *string              `json:"nsxRight"`
	PhoneNumber          *string              `json:"phoneNumber"`
	ReceiveAlerts        *string              `json:"receiveAlerts"`
	State                *StateType           `json:"state"`
}

func (s ActivationStateType) toString() {
	switch s {
	case Disabled:
		fmt.Println("disabled")
	case Disabling:
		fmt.Println("disabling")
	case Enabled:
		fmt.Println("enabled")
	case Enabling:
		fmt.Println("enabling")
	case ToDisable:
		fmt.Println("todisable")
	case ToEnable:
		fmt.Println("toenable")
	default:
		fmt.Println("Invalid ActivationStateType entry")
	}
}

func (s StateType) toString() {
	switch s {
	case Creating:
		fmt.Println("creating")
	case Deleting:
		fmt.Println("deleting")
	case Delivered:
		fmt.Println("delivered")
	case Error:
		fmt.Println("error")
	default:
		fmt.Println("Invalid StateType entry")
	}
}
