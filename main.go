package main

import (
	"fmt"
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
	policy := NewCarbonPolicy(CarbonPolicySpec{
		Taints: TaintPolicy{
			Type: policyTaintTypeTest,
		},
	})
	ca := NewCAClient("https://stubs.gov")
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

	locations, err := ca.GetLocationData(nodes.GetRegions())
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

func printNodes(nodes *Nodes) {
	fmt.Printf("Found %d nodes\n", nodes.Len())
	for node := range nodes.Iterator() {
		fmt.Printf("🟣 %s (%s):\n", node.Name, getLocation(node))
		fmt.Printf("\tTaints: %s\n", node.Spec.Taints)
	}
}
