# Carbon Hack 22
- [hackathon website](https://taikai.network/gsf/hackathons/carbonhack22)
- [idea board](https://docs.google.com/document/d/14VQZwFe-Q8bxf1TbsNNOXfTT37BFGVfUfk0MzP7rE6c/edit#heading=h.68fgvwg50ibg)
- [Carbon Aware SDK](https://github.com/Green-Software-Foundation/carbon-aware-sdk)

## Usage
Install prereqs:
  - [go >=1.18](https://go.dev/dl/)

```
go mod download
go run .
```

## Ideas
- carbon aware operations in k8s
  - CRDs: carbon aware cluster extensions
    - applying taints to nodes in dirty regions
    - auto scaler
    - Deployment wrapper
  - scheduler plugin: bake carbon awareness into the control plane
    - score nodes based on carbon intensity
    - ref
      - https://kubernetes.io/docs/concepts/scheduling-eviction/kube-scheduler/
      - https://kubernetes.io/docs/concepts/scheduling-eviction/scheduling-framework
      - https://medium.com/@juliorenner123/k8s-creating-a-kube-scheduler-plugin-8a826c486a1
      - https://github.com/kubernetes-sigs/scheduler-plugins/blob/master/doc/install.md
      - https://kubernetes.io/docs/tasks/extend-kubernetes/configure-multiple-schedulers/
      - https://github.com/kubernetes-sigs/scheduler-plugins

### Nodes
GKE labels
  - topology.gke.io/zone=us-central1-b
  - topology.kubernetes.io/region=us-central1
  - topology.kubernetes.io/zone=us-central1-b




