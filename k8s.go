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

func (c *K8sClient) UpdateNodes(nodes *Nodes) (*Nodes, error) {
	var updated []v1.Node
	for node := range nodes.Iterator() {
		current := node
		result, err := c.underlying.CoreV1().
			Nodes().
			Update(context.Background(), &current, metav1.UpdateOptions{})
		if err != nil {
			return nil, err
		}
		updated = append(updated, *result)
	}
	return NewNodes(updated), nil
}

func (c *K8sClient) AddTaint(node v1.Node, taint v1.Taint) (v1.Node, error) {
	updated := UpdateNode(node).AddTaint(taint).Build()

	result, err := c.underlying.CoreV1().
		Nodes().
		Update(context.Background(), &updated, metav1.UpdateOptions{})
	return *result, err
}

func (c *K8sClient) RemoveTaint(node v1.Node, key string) (v1.Node, error) {
	updated := UpdateNode(node).RemoveTaint(key).Build()
	result, err := c.underlying.CoreV1().
		Nodes().
		Update(context.Background(), &updated, metav1.UpdateOptions{})
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
