package api

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNamespaces() []string {
	namespaceList, err := client.CoreV1().Namespaces().List(metav1.ListOptions{})
	checkError(err)

	var namespaces []string
	for i:= 0; i < len(namespaceList.Items); i++ {
		namespaces = append(namespaces, namespaceList.Items[i].ObjectMeta.Name)
	}

	return namespaces
}



