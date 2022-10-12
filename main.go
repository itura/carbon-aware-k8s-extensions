package main

import (
	"context"
	"flag"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func main() {
	client, err := createK8sClient()

	response, err := client.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	nodes := NewNodes(response.Items)

	fmt.Printf("There are %d nodes in the cluster\n", nodes.Size())

	region := "us-central1"
	fmt.Printf("nodes in %s:\n", region)
	for i, node := range nodes.ForRegion(region) {
		zone := node.Labels["topology.kubernetes.io/zone"]
		fmt.Printf("%0d %s: %s\n", i, node.Name, zone)
	}
}

type Mapping[T any] map[string]T

type Nodes struct {
	nodes []v1.Node
	m     Mapping[[]v1.Node]
}

func NewNodes(nodes []v1.Node) *Nodes {
	m := Mapping[[]v1.Node]{}
	for _, node := range nodes {
		region := node.Labels["topology.kubernetes.io/region"]
		_, exists := m[region]
		if !exists {
			m[region] = []v1.Node{node}
		} else {
			m[region] = append(m[region], node)
		}
	}

	return &Nodes{
		nodes: nodes,
		m:     m,
	}
}

func (n *Nodes) ForRegion(region string) []v1.Node {
	nodes, present := n.m[region]
	if !present {
		return []v1.Node{}
	}

	return nodes
}

func (n *Nodes) GetAll() []v1.Node {
	return n.nodes
}

func (n *Nodes) Size() int {
	return len(n.nodes)
}

func createK8sClient() (*kubernetes.Clientset, error) {
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
	return clientset, err
}
