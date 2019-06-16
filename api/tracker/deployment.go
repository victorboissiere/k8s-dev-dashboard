package tracker

import (
	"../api"
)

type DeploymentTrackerStatus string

const (
	DeploymentTrackerPending  DeploymentTrackerStatus = "Pending"
	DeploymentTrackerFailing  DeploymentTrackerStatus = "Failing"
)

type DeploymentTracker struct {
	Namespace string `json:"namespace"`
	Name string `json:"name"`
	Status DeploymentTrackerStatus `json:"status"`
}

func getDeployments(namespaces []string) []DeploymentTracker {
	var deploymentTrackers []DeploymentTracker
	for i := 0; i < len(namespaces); i++ {
		deployments := api.GetDeployments(namespaces[i]);
		for j := 0; j < len(deployments); j++ {
			deployment := deployments[j]
			if deployment.Status.HasErrors {
				deploymentTrackers = append(deploymentTrackers, DeploymentTracker{
					Namespace: namespaces[i],
					Name: deployment.Name,
					Status: getDeploymentTrackerStatus(deployment),
				})
			}
		}
	}

	return deploymentTrackers
}

func getDeploymentTrackerStatus(deployment api.Deployment) DeploymentTrackerStatus {
	if deployment.Status.HasErrors {
		return DeploymentTrackerFailing
	}

	return DeploymentTrackerPending
}
