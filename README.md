# Carbon Aware Kubernetes Extensions (cake)
Enable carbon aware operations in your Kubernetes cluster

## cake-node-analyzer

![cake-node-analyzer architecture](docs/cake-node-analyzer-architecture.png)

Consider a hypothetical k8s cluster with autoscaling node pools in 3 regions: A, B and C. The Carbon Aware Operator has gathered data from the Emissions Data Source and determined that region C is the least green. It applies a [taint](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/) to the nodes in region C:

```yaml
# ...
taints:
- key: greensoftware.foundation/carbon-intensity
  value: high
  effect: NoSchedule
```

Now, if resources are submitted to the scheduler, they will not be scheduled on any nodes in region C unless they provide a matching toleration:

```yaml
# ...
tolerations:
- key: greensoftware.foundation/carbon-intensity
  operator: Equal
  value: high
  effect: NoSchedule
```

With this we have a way of using a built-in k8s mechanism to bring carbon awareness to our daily operations. In addition to taints, the CAO could be configured to apply labels to nodes, allowing [node (anti-)affinity](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node) to be carbon aware as well.

The intended way to use the tool would be to configure a set of taints and/or labels that would allow you to naturally move workloads from dirty regions to clean ones. Once utilization of dirty regions is low or gone, you can remove it from your configs entirely.


## Usage
Install prereqs:
  - [go >=1.18](https://go.dev/dl/)

```
# dev
go mod download
go run .
./scripts/verify.sh

# run
docker build -t my.repo/cao:latest .
helm install cao k8s/cao --set image=my.repo/cao:latest
```

## Misc
- [hackathon website](https://taikai.network/gsf/hackathons/carbonhack22)
- [Carbon Aware SDK](https://github.com/Green-Software-Foundation/carbon-aware-sdk)
- scheduler plugin
  - https://kubernetes.io/docs/concepts/scheduling-eviction/kube-scheduler/
  - https://kubernetes.io/docs/concepts/scheduling-eviction/scheduling-framework
  - https://medium.com/@juliorenner123/k8s-creating-a-kube-scheduler-plugin-8a826c486a1
  - https://github.com/kubernetes-sigs/scheduler-plugins/blob/master/doc/install.md
  - https://kubernetes.io/docs/tasks/extend-kubernetes/configure-multiple-schedulers/
  - https://github.com/kubernetes-sigs/scheduler-plugins
-GKE labels
  - topology.gke.io/zone=us-central1-b
  - topology.kubernetes.io/region=us-central1
  - topology.kubernetes.io/zone=us-central1-b




