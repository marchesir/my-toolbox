# Leran Kubernetes

## Goal
- Kubernetes fundamentals
- Troubleshooting
- Stateful workloads
- Services & networking
- Helm
- Kustomize
- Production debugging

Main lab environment:
- Local multi-node kind cluster
- Lightweight containers (BusyBox, NGINX)
- Intentionally broken scenarios

---

# Chapter 1 — Local Multi-Node kind Cluster

## Objective
Create a realistic Kubernetes lab locally.

We will use:
- kind
- Docker
- kubectl

Cluster setup:
- 1 control plane
- 2 worker nodes

---

## Install Tools

### macOS

```bash
brew install kind
```

## Create Multi-Node Cluster

Create file:

```yaml
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
- role: worker
- role: worker
```

Create cluster:

```bash
kind create cluster --config kind-config.yaml --name devops-lab
```

Verify:

```bash
kubectl get nodes
kubectl cluster-info
```

---

# Chapter 2 — Pods, Deployments & Scaling

## Objective
Understand Kubernetes core workload objects.

Main container used:
- BusyBox

Why BusyBox?
- Small
- Fast
- Great for networking/debugging labs

---

## Part A — Pod Basics

### Topics
- What is a Pod?
- Pod lifecycle
- Declarative YAML
- kubectl basics
- Logs and exec

---

## Create BusyBox Pod

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: busybox-test
spec:
  containers:
  - name: busybox
    image: busybox
    command: ['sh', '-c', 'sleep 3600']
```

Apply:

```bash
kubectl apply -f pod.yaml
```

Useful commands:

```bash
kubectl get pods
kubectl describe pod busybox-test
kubectl logs busybox-test
kubectl exec -it busybox-test -- sh
```

---

## Part B — Deployments

### Topics
- ReplicaSets
- Deployments
- Self-healing
- Rolling updates
- Scaling

---

## Create Deployment

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: busybox-deploy
spec:
  replicas: 2
  selector:
    matchLabels:
      app: busybox
  template:
    metadata:
      labels:
        app: busybox
    spec:
      containers:
      - name: busybox
        image: busybox
        command: ['sh', '-c', 'sleep 3600']
```

---

## Scaling

Scale deployment:

```bash
kubectl scale deployment busybox-deploy --replicas=5
```

Check scheduling:

```bash
kubectl get pods -o wide
```

Observe pods across worker nodes.

---

## Self-Healing Test

Delete one pod:

```bash
kubectl delete pod <pod-name>
```

Observe:
- ReplicaSet recreates pod automatically

---

# Chapter 3 — ConfigMaps

## Objective
Separate configuration from containers.

---

## Create ConfigMap

```bash
kubectl create configmap app-config \
  --from-literal=APP_ENV=dev
```

Verify:

```bash
kubectl get configmap
kubectl describe configmap app-config
```

---

## Use ConfigMap in Pod

```yaml
env:
- name: APP_ENV
  valueFrom:
    configMapKeyRef:
      name: app-config
      key: APP_ENV
```

Verify:

```bash
kubectl exec -it busybox-test -- env
```

---

# Chapter 4 — Services Deep Dive

## Topics
- ClusterIP
- NodePort
- LoadBalancer
- Service discovery
- kube-proxy basics
- Endpoints

---

## ClusterIP

Default service type.

Used for:
- Internal communication
- Pod-to-pod traffic

---

## Deploy NGINX

```bash
kubectl create deployment nginx --image=nginx
```

Expose service:

```bash
kubectl expose deployment nginx --port=80 --type=ClusterIP
```

Inspect:

```bash
kubectl get svc
kubectl describe svc nginx
kubectl get endpoints
```

---

## Test Service from BusyBox

```bash
kubectl run testbox --image=busybox -it --rm -- sh
```

Inside pod:

```bash
wget -qO- http://nginx
```

---

# Chapter 5 — NodePort Services

## Objective
Expose applications externally.

---

## Create NodePort Service

```bash
kubectl expose deployment nginx \
  --type=NodePort \
  --port=80
```

Inspect:

```bash
kubectl get svc
kubectl describe svc nginx
```

Access:

```bash
curl http://<node-ip>:<node-port>
```

---

# Chapter 6 — Troubleshooting Broken Services

## Objective
Learn debugging patterns used in real DevOps work.

This section intentionally creates bugs.

---

## Scenario 1 — Wrong Labels

### Problem
Service selector does not match pod labels.

### Symptoms
- Service unreachable
- No endpoints

### Troubleshooting

```bash
kubectl get svc
kubectl get endpoints
kubectl get pods --show-labels
kubectl describe svc nginx
```

### Root Cause
Selector mismatch.

---

## Scenario 2 — Wrong targetPort

### Problem
Service forwards traffic to wrong container port.

### Symptoms
- Timeout
- Connection refused

### Troubleshooting

```bash
kubectl describe svc
kubectl describe pod
kubectl logs
```

---

## Scenario 3 — CrashLoopBackOff

### Problem
Container continuously crashes.

### Troubleshooting

```bash
kubectl get pods
kubectl describe pod
kubectl logs <pod>
```

---

## Scenario 4 — DNS Failure

### Symptoms
Service name not resolving.

### Troubleshooting

```bash
nslookup nginx
kubectl get pods -n kube-system
kubectl logs -n kube-system deployment/coredns
```

---

## Scenario 5 — NetworkPolicy Blocking Traffic

### Symptoms
Pods cannot communicate.

### Troubleshooting

```bash
kubectl get networkpolicy
kubectl describe networkpolicy
kubectl exec
```

---

# Chapter 7 — StatefulSets

## Objective
Understand stateful workloads.

---

## Topics
- StatefulSets
- Stable identities
- Ordered startup
- Persistent storage
- Stateful DNS

---

## Difference from Deployments

Deployments:
- Stateless
- Random pod names

StatefulSets:
- Stable pod names
- Stable storage
- Ordered deployment

Example:

```text
mysql-0
mysql-1
mysql-2
```

---

# Chapter 8 — Headless Services

## Objective
Understand pod-specific DNS.

---

## What is a Headless Service?

A service with:

```yaml
clusterIP: None
```

Used for:
- Databases
- Kafka
- Redis clusters
- Elasticsearch

---

## Create Headless Service

```yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx-headless
spec:
  clusterIP: None
  selector:
    app: nginx
  ports:
  - port: 80
```

---

## StatefulSet Example

```yaml
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: nginx-stateful
spec:
  serviceName: nginx-headless
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx
```

---

## Verify DNS

```bash
kubectl run dns-test --image=busybox -it --rm -- sh
```

Inside pod:

```bash
nslookup nginx-stateful-0.nginx-headless
```

---

# Chapter 9 — Kustomize

## Objective
Manage Kubernetes environments cleanly.

---

## Topics
- Bases
- Overlays
- Patching
- Environment configs
- Secret generators
- Config generators

---

## Recommended Structure

```text
k8s/
├── base/
│   ├── deployment.yaml
│   ├── service.yaml
│   └── kustomization.yaml
├── overlays/
│   ├── dev/
│   ├── stage/
│   └── prod/
```

---

## Base Example

```yaml
resources:
- deployment.yaml
- service.yaml
```

---

## Dev Overlay

```yaml
resources:
- ../../base

namePrefix: dev-

replicas:
- name: nginx
  count: 1
```

---

## Prod Overlay

```yaml
resources:
- ../../base

namePrefix: prod-

replicas:
- name: nginx
  count: 5
```

---

## Deploy with Kustomize

```bash
kubectl apply -k overlays/dev
kubectl apply -k overlays/prod
```

---

# Chapter 10 — Helm

## Objective
Package reusable Kubernetes applications.

---

## Topics
- Helm charts
- values.yaml
- Templates
- Releases
- Upgrades
- Rollbacks

---

## Create Chart

```bash
helm create demo-app
```

---

## Install Chart

```bash
helm install demo demo-app
```

---

## Upgrade Release

```bash
helm upgrade demo demo-app
```

---

## Rollback

```bash
helm rollback demo 1
```

---

## Troubleshooting Helm

### Failed templates

```bash
helm lint
helm template
```

### Failed release

```bash
helm history demo
helm rollback demo
```

---


## Troubleshooting

- Pending pods
- CrashLoopBackOff
- ImagePullBackOff
- DNS failures
- Broken services
- NetworkPolicy issues
- Resource exhaustion
- OOMKilled

---

# Recommended Tools

## Kubernetes
- kind
- kubectl
- k9s
- stern

## DevOps
- Docker
- Helm
- Kustomize

---

# Chapter 11 — Node Scheduling, Taints & Tolerations

# Core Concepts

## Node Selector

Simplest scheduling method.

Pods only run on nodes with matching labels.

Example:
- GPU workloads
- Database nodes
- Monitoring nodes

---

## Node Labels

View labels:

```bash
kubectl get nodes --show-labels
```

Add label:

```bash
kubectl label nodes worker-node-1 workload=backend
```

---

## Node Selector Example

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx-selector
spec:
  nodeSelector:
    workload: backend
  containers:
  - name: nginx
    image: nginx
```

Verify pod placement:

```bash
kubectl get pods -o wide
```

---

# Taints & Tolerations

## Objective

Control which pods are allowed onto nodes.

Very common in production clusters.

---

# Taints

A taint repels pods from a node.

Think:

> "Do not schedule pods here unless explicitly allowed."

---

## Add Taint

```bash
kubectl taint nodes worker-node-1 dedicated=database:NoSchedule
```

Meaning:
- Node reserved for database workloads
- Other pods blocked

---

## View Taints

```bash
kubectl describe node worker-node-1
```

---

# Tolerations

A toleration allows pods to run on tainted nodes.

---

## Example Toleration

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: db-pod
spec:
  tolerations:
  - key: "dedicated"
    operator: "Equal"
    value: "database"
    effect: "NoSchedule"

  containers:
  - name: nginx
    image: nginx
```

---

# Taint Effects

## NoSchedule

Pods will NOT schedule unless tolerated.

---

## PreferNoSchedule

Kubernetes tries to avoid scheduling pods there.

Soft rule.

---

## NoExecute

Pods already running may get evicted.

---

# Real Production Use Cases

## Dedicated Database Nodes

Only DB pods allowed.

---

## GPU Nodes

Only ML/AI workloads allowed.

---

## Spot/Preemptible Nodes

Specific workloads tolerate interruption.

---

## Infra Nodes

Monitoring/logging components isolated.

Examples:
- ingress controllers
- monitoring stack
- logging agents

---

# Hands-on Labs

## Lab 1 — Create Tainted Node

```bash
kubectl taint nodes worker-node-1 dedicated=backend:NoSchedule
```

Deploy normal pod.

Observe:
- Pod remains Pending

Check:

```bash
kubectl describe pod <pod-name>
```

---

## Lab 2 — Add Toleration

Deploy pod with toleration.

Observe:
- Pod schedules successfully

---

## Lab 3 — Remove Taint

```bash
kubectl taint nodes worker-node-1 dedicated=backend:NoSchedule-
```

Observe:
- Scheduling returns to normal

---

# Advanced Scheduling Concepts

## Node Affinity

More powerful than nodeSelector.

Supports:
- required rules
- preferred rules
- expressions

---

## Pod Affinity

Place pods together.

Example:
- frontend near backend

---

## Pod Anti-Affinity

Spread pods across nodes.

Used for:
- High availability
- Avoiding single point of failure

---

# Node Affinity Example

```yaml
affinity:
  nodeAffinity:
    requiredDuringSchedulingIgnoredDuringExecution:
      nodeSelectorTerms:
      - matchExpressions:
        - key: workload
          operator: In
          values:
          - backend
```

---

# Troubleshooting Scheduling Problems

## Common Issues

- Pod Pending
- Missing labels
- Taint mismatch
- Resource exhaustion
- Affinity conflicts

---

## Troubleshooting Commands

```bash
kubectl describe pod <pod-name>
kubectl get nodes --show-labels
kubectl describe node
kubectl top nodes
```

---

- Explain Kubernetes scheduling in interviews
