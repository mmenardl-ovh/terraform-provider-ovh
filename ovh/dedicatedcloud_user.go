package ovh

import (
	"fmt"
	"log"
	"net/url"

	// "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/ovh/go-ovh/ovh"
)

func getDedicatedCloudUserIds(serviceName string, userName string, c *ovh.Client) ([]int64, error) {
	userIds := make([]int64, 0)

	var endpoint string
	if userName != "" {
		endpoint = fmt.Sprintf("/dedicatedCloud/%s/user?name=%s", url.PathEscape(serviceName), url.PathEscape(userName))
		log.Printf("[INFO] Fetching userId list for %s while filtering for name %s", serviceName, userName)
	} else {
		endpoint = fmt.Sprintf("/dedicatedCloud/%s/user", url.PathEscape(serviceName))
		log.Printf("[INFO] Fetching userId list for %s", serviceName)
	}

	if err := c.Get(endpoint, userIds); err != nil {
		return nil, err
	}

	return userIds, nil
}

func getDedicatedCloudUser(serviceName string, userName string, c *ovh.Client) (*DedicatedCloudUser, error) {
	user := &DedicatedCloudUser{}

	userIds, err := getDedicatedCloudUserIds(serviceName, userName, c)
	if err != nil {
		return nil, err
	}
	// if len(userIds) == 0 {
	// 	return nil, fmt.Errorf("Error looking up user %s/%s: user not found", serviceName, userName, err)
	// }

	// if len(userIds) > 1 {
	// 	log.Printf("[INFO] Multiple hits on login %s on %s", userName, serviceName)
	// }

	var endpoint string
	for _, userId := range userIds {
		log.Printf("[INFO] Checking if DedicatedCloudUser userId %d is %s/%s", userId, serviceName, userName)
		endpoint = fmt.Sprint("/dedicatedCloud/%s/user/%d", url.PathEscape(serviceName), userId)
		if errLookup := c.Get(endpoint, &user); errLookup != nil {
			return nil, errLookup
		}
		log.Printf("[INFO] Comparing if userName %s is equal to *user.Login %s", userName, *user.Login)
		if *user.Login == userName {
			return user, nil
		}
	}

	return nil, fmt.Errorf("DedicatedCloud User %s/%s not found", serviceName, userName)
}
