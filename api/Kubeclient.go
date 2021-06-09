package api

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func (api *API) Kubeclient() (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", api.Kubeconfig)
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}
