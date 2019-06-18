package api

type TopNodes struct {
	NodeCount int         `json:"nodeCount"`
	Spec      TopNodeSpec `json:"spec"`
}

type TopNodeSpec struct {
	Request Resource `json:"request"`
	Limit   Resource `json:"limit"`
}

type Resource struct {
	Memory int64 `json:"memory"`
	CPU    int64 `json:"cpu"`
}

func GetTopNodes() TopNodes {
	nodeList := GetNodes()

	return TopNodes{
		NodeCount: len(nodeList),
		Spec:      getTopNodeSpec(nodeList),
	}
}

func getTopNodeSpec(nodeList NodeList) TopNodeSpec {
	limit := Resource{Memory: 0, CPU: 0}
	request := Resource{Memory: 0, CPU: 0}

	for _, node := range nodeList {
		limit.Memory += node.Status.Allocatable.Memory().Value()
		limit.CPU += node.Status.Allocatable.Cpu().Value()

		pods := getPodsFromNode(node.Name)
		podsResource := getRequestFromPods(pods)

		request.Memory += podsResource.Memory
		request.CPU += podsResource.CPU
	}

	return TopNodeSpec{
		Request: request,
		Limit:   limit,
	}
}

func getRequestFromPods(podList PodList) Resource {
	request := Resource{Memory: 0, CPU: 0}

	for _, node := range podList {
		request.Memory += node.Request.Memory
		request.CPU += node.Request.CPU
	}

	request.CPU = request.CPU / 1000

	return request
}
