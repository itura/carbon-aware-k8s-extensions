package main

import (
	"ch22/caapi"
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"net/http"
	"time"
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

	fmt.Printf("There are %d nodes in the cluster\n", nodes.Size())

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

const layout = "2006-01-02"

func callCAApi() {
	config := caapi.NewConfiguration()
	config.Servers = caapi.ServerConfigurations{
		{
			URL: "http://localhost:8080",
			Variables: map[string]caapi.ServerVariable{
				"basePath": {
					DefaultValue: "emissions",
				},
			},
		},
	}
	ca := caapi.NewAPIClient(config).CarbonAwareApi
	startTime, _ := time.Parse(layout, "2022-03-11")
	endTime, _ := time.Parse(layout, "2022-03-12")
	data, response, err := ca.GetAverageCarbonIntensity(context.Background()).
		Location("eastus").
		StartTime(startTime).
		EndTime(endTime).
		Execute()
	fmt.Println(response.Request.URL.String())

	if response.StatusCode != http.StatusOK {
		fmt.Printf("%s", response.StatusCode)
		panic(err.Error())
	}
	fmt.Printf("found %.02f things", *data.CarbonIntensity)
}
