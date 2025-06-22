package tests

import (
	"fmt"
	"testing"

	"github.com/CoreKitMDK/corekit-service-configuration/v2/pkg/configuration"
)

func TestConfiguration(t *testing.T) {
	config := configuration.NewConfiguration()

	config.UseConfigFile = false
	config.UseEnv = false
	config.UseConfigString = false

	config.UseApi = true
	config.ApiUrl = "http://config"
	config.ApiNamespace = "testing-dev"

	client := config.Init()
	if client == nil {
		fmt.Println("Configuration client is nil")
		return
	}

	key := "foo"
	get, err := client.Get(key)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("For Key :" + key + " Value :" + get)
}
