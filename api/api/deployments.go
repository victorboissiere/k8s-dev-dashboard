package api

import (
	v1beta12 "k8s.io/api/extensions/v1beta1"
)

type Deployment struct {
	Name                string            `json:"name"`
	Labels              map[string]string `json:"labels"`
	Replicas            int32             `json:"replicas"`
	AvailableReplicas   int32             `json:"availableReplicas"`
	UnavailableReplicas int32             `json:"unavailableReplicas"`
	UpdatedReplicas     int32             `json:"updatedReplicas"`
}

type DeploymentList = []Deployment
func GetDeployments(namespace string) DeploymentList {
	result := &v1beta12.DeploymentList{}
	err := client.ExtensionsV1beta1().RESTClient().Get().
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
			AvailableReplicas: deployment.Status.AvailableReplicas,
			UnavailableReplicas: deployment.Status.UnavailableReplicas,
			UpdatedReplicas: deployment.Status.UpdatedReplicas,
			Replicas: deployment.Status.Replicas,
		})
	}

	return deploymentList
}

func GetRawDeployment(namespace string, deployment string)  *v1beta12.Deployment {
	result := &v1beta12.Deployment{}
	err := client.ExtensionsV1beta1().RESTClient().Get().
		Namespace(namespace).
		Resource("deployments").
		Name(deployment).
		Do().
		Into(result)

	checkError(err)

	return result
}


