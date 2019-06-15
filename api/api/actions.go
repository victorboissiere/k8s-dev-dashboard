package api

import (
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/json"
	"time"
	"k8s.io/api/core/v1"
)

type PatchDeployment struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value v1.EnvVar `json:"value"`
}

func ForceDeploymentArbitraryRollingUpdate(namespace string, deploymentName string) {
	payload := []PatchDeployment{{
		Op:    "add",
		Path:  "/spec/template/spec/containers/0/env/-",
		Value: v1.EnvVar{
			Name: "LAST_FORCED_ROLLING_UPDATE",
			Value: time.Now().String(),
		},
	}}

	payloadBytes, err := json.Marshal(payload)
	checkError(err)

	_, err = client.AppsV1().
		Deployments(namespace).
		Patch(deploymentName, types.JSONPatchType, payloadBytes)

	checkError(err)
}
