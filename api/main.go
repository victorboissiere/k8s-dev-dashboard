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
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))

	e.GET("/v1/namespaces", dashboard.Namespaces)
	e.GET("/v1/namespaces/:namespace/deployments", dashboard.Deployments)
	e.GET("/v1/namespaces/:namespace/deployment/:deployment/pods", dashboard.Pods)
	e.GET("/v1/nodes", dashboard.Nodes)

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Healthy!")
	})

	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "3000"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
