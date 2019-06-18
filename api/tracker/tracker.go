package tracker

import (
	"../api"
	"github.com/jasonlvhit/gocron"
)

type Tracker struct {
	Deployments []DeploymentTracker `json:"deployments"`
	TopNodes    api.TopNodes        `json:"topNodes"`
}

var (
	tracker Tracker
)

func init() {
	tracker = Tracker{
		Deployments: []DeploymentTracker{},
		TopNodes:    api.TopNodes{},
	}
}

func GetTracker() Tracker {
	return tracker
}

func trackDeploymentsErrors(namespaces []string) {
	tracker.Deployments = getDeployments(namespaces)
}

func trackTopNodes() {
	tracker.TopNodes = api.GetTopNodes()
}

func Setup(namespaces []string) {
	gocron.Every(30).Seconds().Do(trackDeploymentsErrors, namespaces)
	gocron.Every(10).Minutes().Do(trackTopNodes, namespaces)
	trackDeploymentsErrors(namespaces)
	trackTopNodes()

	<-gocron.Start()
}
