## Machine Configuration

Machine Configuration is used to configure the RHCOS on each node in a OCP cluster.

[Machine Configuration Operator](https://github.com/openshift/machine-config-operator) (MCO) is provided by Red Hat to handle machine configuration.

MCO provides a MachineConfig CRD. 
The MachineConfig CRD yaml file contains the recipe of the configuration user would like to apply on particular nodes. 
Once deployed, MCO creates a statically rendered MachineConfig file which includes the MachineConfigs for each node. It then applies that configuration to each node. 

MCO applies changes to operating systems in pools of machines. 
All OpenShift Container Platform clusters start with worker and master (control plane) node pools. 
OCP allows users to set up a custom pool of worker nodes that includes particular hardware features needed by an application.

## OCP Day 0-2:
1. Day 0: Define the requirements of the OCP platform and design it.
2. Day 1: Provision the OCP platform and configure it to a working state.
3. Day 2.0: Day 2.0 is an iOCP project concept which means "Zero O’Clock of Day 2".
4. Day 2: The OCP platform is installed and ready to begin providing services.

## Prerequisites:

Please refer to [NFD README](https://github.com/intel-sandbox/intel.dgpu.operator.prototype/tree/main/nfd/README.md) to setup NFD on OCP Cluster.

## General Configuration for Intel Features
### 1. Deploy alternative firmware path machine configuration. 
This machine configuration sets an alternative firmware path /var/lib/firmware so that user can load OOT firmware
on OCP node. This is needed because the default firmware directory is read-only on RHCOS OCP nodes.

To deploy, use the below command:

```$ oc apply -f https://github.com/intel-sandbox/intel.dgpu.operator.prototype/blob/main/machine_configuration/100-alternative-fw-path-for-worker-nodes.yaml```

Note: This configuration will trigger the rebooting of all the nodes in the worker machine configuration pool. Rebooting the node is not preferred by iOCP project. A better solution is in the plan.

## dGPU Machine Configuration Setup Steps
### 1. Create dGPU machine config pool:

This will create a custom machine config pool of worker nodes that contain Intel dGPU card using NFD label `intel.feature.node.kubernetes.io/gpu: 'true'`. 

To create the dGPU machine config pool, run

```$ oc apply -f https://github.com/intel-sandbox/intel.dgpu.operator.prototype/blob/main/machine_configuration/intel-dgpu-machine-config-pool.yaml```

### 2. Disable in-tree ast driver 

```$ oc apply -f https://github.com/intel-sandbox/intel.dgpu.operator.prototype/blob/main/machine_configuration/100-intel-dgpu-machine-config-disable-ast.yaml```

This will blacklist (disable loading) of ast driver on worker nodes in Intel dGPU machine config pool. This is needed prior to loading Intel dGPU drivers and firmware. 

Note: ToDo: Add the link to the known dGPU driver issue.

Note: The driver that needs to be blacklisted will vary on the platform the user is using. As a result, please use this as a reference for your environment.

## QAT Machine Configuration Setup Steps

### 1. Create QAT machine config pool:

This will create a custom machine config pool of worker nodes that contain Intel QAT accelerator using NFD label `intel.feature.node.kubernetes.io/qat: 'true'`. 

To create the QAT machine config pool, run

```$ oc apply -f https://github.com/intel-sandbox/intel.dgpu.operator.prototype/blob/main/machine_configuration/intel-qat-machine-config-pool.yaml```

### 2. Turn on intel_iommu kernel parameter for QAT

For QAT, intel_iommu is a kernel parameter that needs to be turned on using the following command.

```$ oc apply -f https://github.com/intel-sandbox/intel.dgpu.operator.prototype/blob/main/machine_configuration/100-intel-qat-intel-iommu-on.yaml```

Note: This will reboot all the worker nodes in the Intel QAT machine config pool one by one to turn on intel_iommu kernel parameter. Rebooting the node is not preferred by iOCP project. A better solution is in the plan.

### 3. Deploy QAT OOT FW

```$ oc apply -f https://github.com/intel-sandbox/intel.dgpu.operator.prototype/blob/main/machine_configuration/100-intel-qat-firmware-blob.yaml```

This will deploy the QAT OOT Firmware on nodes in Intel QAT machine config pool.

Note: This will reboot all the worker nodes in the Intel QAT machine config pool one by one to deploy the QAT OOT firmware. Rebooting the node is not preferred by iOCP project. A better solution is in the plan.
