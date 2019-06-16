package main

import (
	"./controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"os"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		DisableStackAll: true,
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPatch},
	}))

	e.GET("/v1/tracker", dashboard.GetTracker)
	e.GET("/v1/namespaces", dashboard.GetNamespaces)
	e.GET("/v1/namespaces/applications", dashboard.GetApplicationsNamespaces)
	e.GET("/v1/namespaces/:namespace/deployments", dashboard.GetDeployments)
	e.GET("/v1/namespaces/:namespace/deployments/:deployment/pods", dashboard.GetPods)
	e.GET("/v1/namespaces/:namespace/deployments/:deployment/events", dashboard.GetEvents)
	e.PATCH("/v1/namespaces/:namespace/deployments/:deployment/arbitrary-rolling-update", dashboard.ForceArbitraryRollingUpdate)
	e.GET("/v1/nodes", dashboard.GetNodes)

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Healthy!")
	})

	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "3000"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
