## NFD

The [Node Feature Discovery](https://github.com/openshift/node-feature-discovery) (NFD) manages the detection of hardware features and configuration in OCP cluster by labeling the nodes with hardware-specific information.

NFD labels the host with node-specific attributes, such as PCI cards, kernel, operating system version, and so on. In this project NFD is used to automatically label nodes with specific Hardware features like Intel dGPU card, Intel SGX and Intel QAT etc.

### Labels discription

* Label name - label discription

### Steps to install and configure NFD operator on OCP Cluster.

1. [NFD operator](https://github.com/openshift/cluster-nfd-operator) is used to manage the NFD stack. So firstly follow [NFD operator installation instructions](https://docs.openshift.com/container-platform/4.12/hardware_enablement/psap-node-feature-discovery-operator.html#install-operator-web-console_node-feature-discovery-operator) to install NFD operator on the OCP cluster. 

2. Create NodeFeatureDiscovery instance

```$ oc apply -f https://github.com/intel/intel-technology-enabling-for-openshift/blob/main/nfd/node-feature-discovery-openshift.yaml```

3. Create NodeFeatureRule instance

```$ oc apply -f https://github.com/intel/intel-technology-enabling-for-openshift/blob/main/nfd/node-feature-rules-openshift.yaml```

