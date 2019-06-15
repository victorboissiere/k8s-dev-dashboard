package dashboard

import (
	"net/http"
	"github.com/labstack/echo"
	"../api"
)

func Namespaces(c echo.Context) error {
	return c.JSON(http.StatusOK, api.GetNamespaces())
}

func Deployments(c echo.Context) error {
	return c.JSON(http.StatusOK, api.GetDeployments(c.Param("namespace")))
}

func Pods(c echo.Context) error {
	rawDeployment := api.GetRawDeployment(c.Param("namespace"), c.Param("deployment"))

	return c.JSON(http.StatusOK, api.GetPods(c.Param("namespace"), rawDeployment.Spec.Selector.MatchLabels))
}

func Nodes(c echo.Context) error {
	return c.JSON(http.StatusOK, api.GetNodes())
}
