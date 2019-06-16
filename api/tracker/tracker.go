package tracker

import (
	"github.com/jasonlvhit/gocron"
)


type Tracker struct {
	Deployments []DeploymentTracker `json:"deployments"`
}

var (
	tracker Tracker
)

func init() {
	tracker = Tracker{
		Deployments: []DeploymentTracker{},
	}
}

func GetTracker() Tracker {
	return tracker
}

func trackDeploymentsErrors(namespaces []string) {
	tracker.Deployments = getDeployments(namespaces)
}

func Setup(namespaces []string) {
	gocron.Every(30).Seconds().Do(trackDeploymentsErrors, namespaces)
	<- gocron.Start()
}

