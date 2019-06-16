package dashboard

import (
	"os"
	"strings"
	"../api"
)

const DefaultDeveloperNamespaces = "integration,validation,staging,qualification,preproduction,production"

func getEnvValue(envKey string, defaultEnvValue string) string {
	envValue, exists := os.LookupEnv(envKey)
	if exists {
		return envValue
	}

	return defaultEnvValue
}

func selectNamespaces(namespaces []string, envKey string, defaultEnvValue string) []string {
	userSelectedNamespaces := getEnvValue(envKey, defaultEnvValue) + ","

	selectedNamespaces := make([]string, 0)
	for i := 0; i < len(namespaces); i++ {
		if strings.Contains(userSelectedNamespaces, namespaces[i] + ",") {
			selectedNamespaces = append(selectedNamespaces, namespaces[i])
		}
	}

	return selectedNamespaces
}

func isNamespaceInEnv(namespace string, envKey string, defaultEnvValue string) bool {
	namespaces := getEnvValue(envKey, defaultEnvValue) + ","

	return strings.Contains(namespaces, namespace + ",")
}

func getApplicationsNamespaces() []string {
	return selectNamespaces(api.GetNamespaces(),"APPLICATION_NAMESPACES", DefaultDeveloperNamespaces)
}

func isNamespaceAuthorizedForArbitraryRollingUpdate(namespace string) bool {
	return isNamespaceInEnv(namespace, "ARBITRARY_ROLLING_UPDATE_NAMESPACES", "")
}

