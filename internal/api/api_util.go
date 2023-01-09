package api

import (
	"os"

	"github.com/conductor-sdk/conductor-go/sdk/client"
	"github.com/conductor-sdk/conductor-go/sdk/settings"
)

const (
	KEY                  = "KEY"
	SECRET               = "SECRET"
	CONDUCTOR_SERVER_URL = "CONDUCTOR_SERVER_URL"
)

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
