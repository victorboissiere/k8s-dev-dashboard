package api

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes/scheme"
)

type Pod struct {
	Name   string    `json:"name"`
	HostIP string    `json:"hostIP"`
	PodIP  string    `json:"podIP"`
	Status PodStatus `json:"status"`
	Request PodRequest `json:"request"`
}

type PodStatus struct {
	StartTime         *metav1.Time         `json:"startTime"`
	Phase             v1.PodPhase          `json:"phase"`
	Reason            string               `json:"reason"`
	Message           string               `json:"message"`
	ContainerStatuses []v1.ContainerStatus `json:"containerStatuses"`
}

type PodRequest struct {
	Memory int64 `json:"memory"`
	CPU int64 `json:"cpu"`
}

type PodList = []Pod

func GetPods(namespace string, matchLabels map[string]string) PodList {
	result := &v1.PodList{}
	options := &metav1.ListOptions{LabelSelector: labels.FormatLabels(matchLabels)}
	err := client.CoreV1().RESTClient().Get().
		Namespace(namespace).
		Resource("pods").
		VersionedParams(options, scheme.ParameterCodec).
		Do().
		Into(result)

	checkError(err)

	return mapPods(result)
}

func getPodsFromNode(nodeName string) PodList {
	result := &v1.PodList{}
	options := &metav1.ListOptions{ FieldSelector: "spec.nodeName=" + nodeName }
	err := client.CoreV1().RESTClient().Get().
		Resource("pods").
		VersionedParams(options, scheme.ParameterCodec).
		Do().
		Into(result)

	checkError(err)

	return mapPods(result)
}


func mapPods(rawPodList *v1.PodList) PodList {
	var podList = PodList{}
	for i := 0; i < len(rawPodList.Items); i++ {
		pod := rawPodList.Items[i]
		podList = append(podList, Pod{
			Name:   pod.Name,
			HostIP: pod.Status.HostIP,
			PodIP:  pod.Status.PodIP,
			Status: PodStatus{
				Phase:             pod.Status.Phase,
				StartTime:         pod.Status.StartTime,
				Reason:            pod.Status.Reason,
				Message:           pod.Status.Message,
				ContainerStatuses: pod.Status.ContainerStatuses,
			},
			Request: getPodRequests(pod),
		})
	}

	return podList
}

func getPodRequests (pod v1.Pod) PodRequest {
	var totalMemoryRequest int64 = 0
	var totalCPURequest int64 = 0

	for _, container := range pod.Spec.Containers {
		request := container.Resources.Requests
		totalMemoryRequest += request.Memory().Value()
		totalCPURequest += request.Cpu().MilliValue()
	}

	return PodRequest{
		Memory: totalMemoryRequest,
		CPU: totalCPURequest,
	}
}
