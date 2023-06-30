# Setting up Node Feature Discovery
[Node Feature Discovery (NFD) Operator](https://docs.openshift.com/container-platform/4.12/hardware_enablement/psap-node-feature-discovery-operator.html) manages the deployment and lifecycle of the NFD add-on to detect hardware features and system configuration, such as PCI cards, kernel, operating system version, etc.

# Prerequisites
- Provisioned RHOCP 4.12 cluster. Follow steps [here](/README.md#provisioning-rhocp-cluster).

# Install NFD Operator
Follow the guide below to install the NFD operator using CLI or web console. 
- [Install from the CLI](https://docs.openshift.com/container-platform/4.12/hardware_enablement/psap-node-feature-discovery-operator.html#install-operator-cli_node-feature-discovery-operator)
- [Install from the web console](https://docs.openshift.com/container-platform/4.12/hardware_enablement/psap-node-feature-discovery-operator.html#install-operator-web-console_node-feature-discovery-operator)

# Configure NFD Operator
Note: As RHOCP cluster administrator, you might need to merge the NFD operator config from the following Custom Resources (CRs) with other NFD operator configs that are already applied on your cluster.  

1. Create `NodeFeatureDiscovery` CR instance.
```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/nfd/node-feature-discovery-openshift.yaml 
```

2.	Create `NodeFeatureRule` CR instance.
```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/nfd/node-feature-rules-openshift.yaml
```

# Verification 
Use the command shown below to verify whether the nodes are labeled properly by NFD:
```
$ oc describe node node_name | grep intel.feature.node.kubernetes.io
```
Example output: 
```
intel.feature.node.kubernetes.io/dgpu-canary=true
intel.feature.node.kubernetes.io/gpu=true
```

# Labels table
| Label | Intel hardware feature | 
| ----- | ---------------------- |
| `intel.feature.node.kubernetes.io/gpu=true` | Intel® Data Center GPU Flex Series | 
| `intel.feature.node.kubernetes.io/sgx=true` | Intel® SGX | 

# See Also