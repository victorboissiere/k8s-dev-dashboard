package api

import (
	"k8s.io/api/core/v1"
)

type NodeStatus struct {
	Phase       v1.NodePhase `json:"phase"`
	NodeInfo    v1.NodeSystemInfo `json:"nodeInfo"`
	Allocatable v1.ResourceList `json:"allocatable"`
}

type Node struct {
	Name        string `json:"name"`
	Annotations map[string]string `json:"annotations"`
	Labels      map[string]string `json:"labels"`
	Status NodeStatus `json:"status"`
}

type NodeList = []Node

func GetNodes() NodeList {
	result := &v1.NodeList{}

	err := client.CoreV1().RESTClient().Get().
		Resource("nodes").
		Do().
		Into(result)

	checkError(err)


	var nodeList = NodeList{}
	for i := 0; i < len(result.Items); i++ {
		node := result.Items[i]
		nodeList = append(nodeList, Node{
			Name: node.ObjectMeta.Name,
			Labels: node.ObjectMeta.Labels,
			Annotations: node.ObjectMeta.Annotations,
			Status: NodeStatus{
				Phase: node.Status.Phase,
				NodeInfo: node.Status.NodeInfo,
				Allocatable: node.Status.Allocatable,
			},
		})
	}

	return nodeList
}

