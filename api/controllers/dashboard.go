package dashboard

import (
	"../api"
	"../tracker"
	"github.com/labstack/echo"
	"net/http"
)

func init() {
	go tracker.Setup(getApplicationsNamespaces())
}

func GetTracker(c echo.Context) error {
	return c.JSON(http.StatusOK, tracker.GetTracker())
}

func GetNamespaces(c echo.Context) error {
	return c.JSON(http.StatusOK, api.GetNamespaces())
}

func GetApplicationsNamespaces(c echo.Context) error {
	return c.JSON(http.StatusOK, getApplicationsNamespaces())
}

func GetDeployments(c echo.Context) error {
	return c.JSON(http.StatusOK, api.GetDeployments(c.Param("namespace")))
}

func GetPods(c echo.Context) error {
	rawDeployment := api.GetRawDeployment(c.Param("namespace"), c.Param("deployment"))

	return c.JSON(http.StatusOK, api.GetPods(c.Param("namespace"), rawDeployment.Spec.Selector.MatchLabels))
}

func GetEvents(c echo.Context) error {
	return c.JSON(http.StatusOK, api.GetEvents(c.Param("namespace"), c.Param("deployment")))
}

func GetNodes(c echo.Context) error {
	return c.JSON(http.StatusOK, api.GetNodes())
}

func ForceArbitraryRollingUpdate(c echo.Context) error {
	namespace := c.Param("namespace")
	if !isNamespaceAuthorizedForArbitraryRollingUpdate(namespace) {
		return c.JSON(http.StatusForbidden, "")
	}

	api.ForceDeploymentArbitraryRollingUpdate(namespace, c.Param("deployment"))
	return c.JSON(http.StatusOK, "")
}
