package main

import (
	"context"
	"flag"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

type K8sClient struct {
	underlying *kubernetes.Clientset
}

func (c *K8sClient) ListNodes() (*v1.NodeList, error) {
	response, err := c.underlying.CoreV1().
		Nodes().
		List(context.Background(), metav1.ListOptions{})
	return response, err
}

func (c *K8sClient) AddTaint(node v1.Node, taint v1.Taint) (v1.Node, error) {
	node.Spec.Taints = append(node.Spec.Taints, taint)

	result, err := c.underlying.CoreV1().
		Nodes().
		Update(context.Background(), &node, metav1.UpdateOptions{})
	return *result, err
}

func (c *K8sClient) RemoveTaint(node v1.Node, key string) (v1.Node, error) {
	var updated []v1.Taint
	for _, taint := range node.Spec.Taints {
		if taint.Key != key {
			updated = append(updated, taint)
		}
	}
	node.Spec.Taints = updated

	result, err := c.underlying.CoreV1().
		Nodes().
		Update(context.Background(), &node, metav1.UpdateOptions{})
	return *result, err
}

func NewK8sClient() (*K8sClient, error) {
	//// creates the in-cluster config
	//config, err := rest.InClusterConfig()
	//if err != nil {
	//	panic(err.Error())
	//}
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return &K8sClient{underlying: clientset}, err
}
