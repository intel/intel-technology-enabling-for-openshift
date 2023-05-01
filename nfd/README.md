## NFD

The [Node Feature Discovery](https://github.com/openshift/node-feature-discovery) (NFD) manages the detection of hardware features and configuration in a OCP cluster by labeling the nodes with hardware-specific information.

NFD labels the host with node-specific attributes, such as PCI cards, kernel, operating system version, and so on. In this project, NFD is used to automatically label nodes with specific hardware features such as Intel® Data Center GPU Flex Series, Intel® SGX and Intel® QAT etc.

### Steps to install and configure NFD operator:

1. Follow the [NFD operator installation instructions](https://docs.openshift.com/container-platform/4.12/hardware_enablement/psap-node-feature-discovery-operator.html#install-operator-web-console_node-feature-discovery-operator) to install the NFD operator. 

2. Create NodeFeatureDiscovery instance

   ```
   $ oc apply -f https://github.com/intel/intel-technology-enabling-for-openshift/blob/main/nfd/node-feature-discovery-openshift.yaml
   ```

3. Create NodeFeatureRule instance

   ```
   $ oc apply -f https://github.com/intel/intel-technology-enabling-for-openshift/blob/main/nfd/node-feature-rules-openshift.yaml
   ```

4. If NFD has been configured properly, NFD will detect and label Intel® hardware features present on each node in the cluster.

    Feature | Label
    --- | ---
    Intel® Data Center GPU Flex Series | intel.feature.node.kubernetes.io/gpu=true
    Intel® QAT | intel.feature.node.kubernetes.io/qat=true
    Intel® SGX | intel.feature.node.kubernetes.io/sgx=true
