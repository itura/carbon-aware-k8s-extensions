package main

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func main() {
	callk8sApi()
	callCAApi()
}

func callk8sApi() {
	k8s, err := NewK8sClient()

	response, err := k8s.ListNodes()
	if err != nil {
		panic(err.Error())
	}
	nodes := NewNodes(response.Items)

	fmt.Printf("There are %d nodes in the cluster\n---\n", nodes.Size())

	region := "us-central1"
	fmt.Printf("nodes in %s:\n", region)
	for i, node := range nodes.ForRegion(region) {
		zone := node.Labels["topology.kubernetes.io/zone"]
		fmt.Printf("%0d %s: %s\n", i, node.Name, zone)
		fmt.Printf("\t taints: %s\n", node.Spec.Taints)

		node, err = k8s.RemoveTaint(node, "blahblah")
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("\t taints: %s\n", node.Spec.Taints)
		node, err = k8s.AddTaint(node, v1.Taint{
			Key:    "blahblah",
			Effect: "NoSchedule",
		})
		if err != nil {
			panic(err.Error())
		}

		fmt.Printf("\t taints: %s\n", node.Spec.Taints)
	}
}

func callCAApi() {
	ca := NewCAClient("http://localhost:8080")
	intensity, err := ca.GetAverageCarbonIntensity("eastus", "2022-03-11", "2022-03-12")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		panic(err.Error())
	}
	fmt.Printf("found %.02f things", intensity)
}
