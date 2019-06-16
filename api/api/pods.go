package api

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes/scheme"
)

type PodStatus struct {
	StartTime         *metav1.Time         `json:"startTime"`
	Phase             v1.PodPhase          `json:"phase"`
	Reason            string               `json:"reason"`
	Message           string               `json:"message"`
	ContainerStatuses []v1.ContainerStatus `json:"containerStatuses"`
}

type Pod struct {
	Name   string    `json:"name"`
	HostIP string    `json:"hostIP"`
	PodIP  string    `json:"podIP"`
	Status PodStatus `json:"status"`
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

	var podList = PodList{}
	for i := 0; i < len(result.Items); i++ {
		pod := result.Items[i]
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
		})
	}

	return podList
}
