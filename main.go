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

func callCAApi() {
	//config := caapi.NewConfiguration()
	//config.Host = "localhost:8080"
	//config.Scheme = "http"
	//ca := caapi.NewAPIClient(config).CarbonAwareApi
	//emissionsForecastBatchDTO := []caapi.EmissionsForecastBatchDTO{*caapi.NewEmissionsForecastBatchDTO(time.Now(), "eastus")} // []EmissionsForecastBatchDTO | Array of requested forecasts. (optional)
	//resp, r, err := ca.BatchForecastDataAsync(context.Background()).EmissionsForecastBatchDTO(emissionsForecastBatchDTO).Execute()

	//d, r, err := ca.GetCurrentForecastData(context.Background()).
	//	Location([]string{"eastus"}).
	//	Execute()
	//if err != nil {
	//	fmt.Printf("%s", r.StatusCode)
	//	panic(err.Error())
	//}
	//fmt.Printf("found %d things", len(resp))
}
