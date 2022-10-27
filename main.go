package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/util/yaml"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"os"
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		fmt.Fprintf(os.Stderr, "😭\n")
		os.Exit(1)
	}

	fmt.Printf("😎\n")
}

func run() error {
	policy, err := initPolicy()
	if err != nil {
		return err
	}
	isInCluster := len(os.Getenv("KUBERNETES_SERVICE_HOST")) > 0
	k8s, err := NewK8sClient(isInCluster)
	if err != nil {
		return err
	}

	response, err := k8s.ListNodes()
	if err != nil {
		return err
	}
	nodes := NewNodes(response.Items)
	printNodes(nodes)

	dataSource, err := initDataSource(policy.Spec)
	if err != nil {
		return err
	}
	locations, err := dataSource.GetLocationData(nodes.GetRegions())
	if err != nil {
		return err
	}
	fmt.Printf("Found %d locations\n", locations.Len())

	fmt.Println("Updating nodes...")
	nodes, err = policy.
		SetNodes(nodes).
		SetLocations(locations).
		UpdateNodes()
	if err != nil {
		return err
	}

	nodes, err = k8s.UpdateNodes(nodes)
	if err != nil {
		return err
	}
	fmt.Println("done.")
	printNodes(nodes)
	return nil
}

func initPolicy() (*CarbonPolicy, error) {
	var spec CarbonPolicySpec
	dat, err := os.ReadFile("./config.yaml")
	if err != nil {
		spec = CarbonPolicySpec{}
	} else {
		err = yaml.Unmarshal(dat, &spec)
		if err != nil {
			return nil, err
		}
	}
	policy := NewCarbonPolicy(spec)
	return policy, nil
}

func initDataSource(spec CarbonPolicySpec) (EmissionsDataSource, error) {
	switch spec.DataSource.Type {
	case optStub:
		return NewStubEDS(), nil
	case optCAAPI:
		return NewCAClient("https://stubs.gov"), nil
	default:
		return nil, fmt.Errorf("invalid value for .dataSource.type")
	}
}

func printNodes(nodes *Nodes) {
	fmt.Printf("Found %d nodes\n", nodes.Len())
	for node := range nodes.Iterator() {
		fmt.Printf("🟣 %s (%s):\n", node.Name, getLocation(node))
		fmt.Printf("\tTaints: %s\n", node.Spec.Taints)
		fmt.Printf("\tLabels: %s\n", node.Labels)
	}
}
