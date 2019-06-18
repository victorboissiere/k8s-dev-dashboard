package api

import (
	"k8s.io/api/core/v1"
)

type Node struct {
	Name        string            `json:"name"`
	Annotations map[string]string `json:"annotations"`
	Labels      map[string]string `json:"labels"`
	Status      NodeStatus        `json:"status"`
	Request     NodeRequest       `json:"request"`
}

type NodeStatus struct {
	Phase       v1.NodePhase      `json:"phase"`
	NodeInfo    v1.NodeSystemInfo `json:"nodeInfo"`
	Allocatable v1.ResourceList   `json:"allocatable"`
}

type NodeRequest struct {
	Memory int64 `json:"memory"`
	CPU    int64 `json:"cpu"`
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
			Name:        node.ObjectMeta.Name,
			Labels:      node.ObjectMeta.Labels,
			Annotations: node.ObjectMeta.Annotations,
			Request:     GetNodeRequest(node.ObjectMeta.Name),
			Status: NodeStatus{
				Phase:       node.Status.Phase,
				NodeInfo:    node.Status.NodeInfo,
				Allocatable: node.Status.Allocatable,
			},
		})
	}

	return nodeList
}

func GetNodeRequest(nodeName string) NodeRequest {
	pods := getPodsFromNode(nodeName)
	var totalMemoryRequest int64 = 0
	var totalCPURequest int64 = 0

	for _, pod := range pods {
		totalMemoryRequest += pod.Request.Memory
		totalCPURequest += pod.Request.CPU
	}

	return NodeRequest{
		Memory: totalMemoryRequest,
		CPU:    totalCPURequest,
	}
}
