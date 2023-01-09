package api

import (
	"fmt"
	"os"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
)

const (
	KEY                  = "KEY"
	SECRET               = "SECRET"
	CONDUCTOR_SERVER_URL = "CONDUCTOR_SERVER_URL"
)

func GetWorkflowExecutionURL(workflowId string) string {
	url := os.Getenv(CONDUCTOR_SERVER_URL)
	prefix := url[:len(url)-4]
	return fmt.Sprintf("%s/execution/%s", prefix, workflowId)
}

func GetApiClientWithAuthentication() *client.APIClient {
	return client.NewAPIClient(
		getAuthenticationSettings(),
		getHttpSettingsWithAuth(),
	)
}

func getAuthenticationSettings() *settings.AuthenticationSettings {
	return settings.NewAuthenticationSettings(
		os.Getenv(KEY),
		os.Getenv(SECRET),
	)
}

func getHttpSettingsWithAuth() *settings.HttpSettings {
	return settings.NewHttpSettings(
		os.Getenv(CONDUCTOR_SERVER_URL),
	)
}
