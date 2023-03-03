## Machine Configuration

Machine Configuration is used to configure the RHCOS on each node in a OCP cluster.

[Machine Configuration Operator](https://github.com/openshift/machine-config-operator) (MCO) is provided by Red Hat to handle machine configuration.

### Prerequisites:

Make sure NFD operator has been installed and configured properly 

Please refer to [instructions](/nfd/README.md#steps-to-install-and-configure-nfd-operator-on-ocp-cluster) to install and configure NFD operator on OCP Cluster.

### General Configuration for Provisioning Intel Hardware Features
* Set up an alternative firmware path for the Cluster 

The following command sets `/var/lib/firmware` as an alternative firmware path for RHCOS kernel to load OOT firmware. Because the default firmware directory `/lib/firmware` is mounted as read-only directory on OCP cluster.

```$ oc apply -f https://github.com/intel/intel-technology-enabling-for-openshift/blob/main/machine_configuration/100-alternative-fw-path-for-worker-nodes.yaml```

Note: This command might trigger nodes rebooting in the worker machine configuration pool. It is not preferred by this project. A better solution is in the plan.

### Machine Configuration for Provisioning Intel dGPU
* Create dGPU machine config pool:

Below command creates a custom machine config pool for worker nodes with Intel dGPU card. These nodes are labeled `intel.feature.node.kubernetes.io/gpu: 'true'` by NFD. 

```$ oc apply -f https://github.com/intel/intel-technology-enabling-for-openshift/blob/main/machine_configuration/intel-dgpu-machine-config-pool.yaml```

* Disable in-tree ast driver 

```$ oc apply -f https://github.com/intel/intel-technology-enabling-for-openshift/blob/main/machine_configuration/100-intel-dgpu-machine-config-disable-ast.yaml```

Becuase of the known incompatible issue from intel i915 driver, This command is used to blacklist (disable loading) of ast driver on worker nodes in Intel dGPU machine config pool. This is needed prior to loading Intel dGPU drivers and firmware. 

Note: This known incompatible issue will be fixed when upgrade to intel data center gpud driver - 1.0.0 using RHEL-9.x i915 driver

### Machine Configuration for Provisioning Intel QAT

* Create QAT machine config pool:

Following command creates a custom machine config pool for worker nodes with Intel QAT feature. These nodes are labeled `intel.feature.node.kubernetes.io/qat: 'true'` by NFD. 

```$ oc apply -f https://github.com/intel/intel-technology-enabling-for-openshift/blob/main/machine_configuration/intel-qat-machine-config-pool.yaml```

* Turn on intel_iommu kernel parameter for QAT

For QAT, intel_iommu is a kernel parameter that needs to be turned on using the following command.

```$ oc apply -f https://github.com/intel/intel-technology-enabling-for-openshift/blob/main/machine_configuration/100-intel-qat-intel-iommu-on.yaml```

Note: This will reboot the worker nodes in the Intel QAT machine config pool one by one to turn on intel_iommu kernel parameter. Rebooting the node is not preferred A better solution is in the plan.
