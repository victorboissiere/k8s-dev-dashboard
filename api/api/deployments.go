package api

import (
	"k8s.io/api/core/v1"
	v12 "k8s.io/api/apps/v1"
)

type DeploymentStatus struct {
	HasErrors bool `json:"hasErrors"`
	IsInProgress bool `json:"isInProgress"`
}
type Deployment struct {
	Name                string            `json:"name"`
	Labels              map[string]string `json:"labels"`
	Replicas            int32             `json:"replicas"`
	AvailableReplicas   int32             `json:"availableReplicas"`
	UnavailableReplicas int32             `json:"unavailableReplicas"`
	Status           DeploymentStatus `json:"status"`
}

type DeploymentList = []Deployment

func GetDeployments(namespace string) DeploymentList {
	result := &v12.DeploymentList{}
	err := client.AppsV1().RESTClient().Get().
		Namespace(namespace).
		Resource("deployments").
		Do().
		Into(result)

	checkError(err)

	var deploymentList = DeploymentList{}
	for i := 0; i < len(result.Items); i++ {
		deployment := result.Items[i]
		deploymentList = append(deploymentList, Deployment{
			Name: deployment.ObjectMeta.Name,
			Labels: deployment.ObjectMeta.Labels,
			Replicas: deployment.Status.Replicas,
			AvailableReplicas: deployment.Status.AvailableReplicas,
			UnavailableReplicas: deployment.Status.UnavailableReplicas,
			Status: getDeploymentStatus(deployment),
		})
	}

	return deploymentList
}

func getDeploymentStatus(deployment v12.Deployment) DeploymentStatus {
	status := DeploymentStatus{ HasErrors: false, IsInProgress: false }

	conditions := deployment.Status.Conditions
	for i := 0; i < len(conditions); i++ {
		condition := conditions[i]
		if condition.Type == v12.DeploymentProgressing && condition.Status == v1.ConditionFalse {
				status.HasErrors = true
		} else if condition.Type == v12.DeploymentReplicaFailure {
			status.HasErrors = true
		}
	}

	if status.HasErrors {
		return status
	}

	if deployment.Status.ReadyReplicas < *deployment.Spec.Replicas {
		status.HasErrors = true
	} else if deployment.Status.Replicas > *deployment.Spec.Replicas {
		status.IsInProgress = true
	}

	return status
}

func GetRawDeployment(namespace string, deployment string)  *v12.Deployment {
	result := &v12.Deployment{}
	err := client.ExtensionsV1beta1().RESTClient().Get().
		Namespace(namespace).
		Resource("deployments").
		Name(deployment).
		Do().
		Into(result)

	checkError(err)

	return result
}


