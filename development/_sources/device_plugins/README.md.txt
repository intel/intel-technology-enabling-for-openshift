# Setting up Intel Device Plugins Operator 

## Overview
Intel Device Plugins are utilized to advertise Intel hardware features (resources) to a Red Hat OpenShift Container Platform (RHOCP) Cluster. This allows workloads running on pods deployed within the clusters to leverage these features. To handle the deployment and lifecycle of these device plugins, the [Intel Device Plugins Operator](https://catalog.redhat.com/software/container-stacks/detail/61e9f2d7b9cdd99018fc5736) is used. The Intel Device Plugins container images and operator have been officially certified and published on the [Red Hat Ecosystem Catalog](https://catalog.redhat.com/software/container-stacks/detail/61e9f2d7b9cdd99018fc5736). For more details on the upstream project, please refer to [Intel Device Plugins for Kubernetes](https://github.com/intel/intel-device-plugins-for-kubernetes).  

## Prerequisities
- Provisioned RHOCP cluster. Follow steps [here](/README.md).
- Setup Node Feature Discovery (NFD). Follow steps [here](/nfd/README.md).
- Follow the additional prerequisites for provisioning Intel® Data Center GPU: 
    - Setup out of tree drivers for Intel Data Center GPU provisioning. Follow the steps listed [here](/kmmo/README.md). 
- Follow the additional prerequisites for provisioning Intel® QuickAssist Technology: 
    - Configure MCO for provisioning Intel QAT. Follow steps [here](/machine_configuration/README.md#machine-configuration-for-provisioning-intel-qat).

## Install Intel Device Plugins Operator on Red Hat OpenShift
### Installation via web console
Follow the steps below to install Intel Device Plugins Operator using OpenShift web console:
1.	In the OpenShift web console, navigate to **Operator** -> **OperatorHub**.
2.	Search for **Intel Device Plugins Operator** in all items field -> Click **Install**.
### Verify Installation via web console
1.	Go to **Operator** -> **Installed Operators**.
2.	Verify that the status of the operator is **Succeeded**.

### Installation via command line interface (CLI)
Apply the [install_operator.yaml](/device_plugins/install_operator.yaml) file:
```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/device_plugins/install_operator.yaml
```

### Verify Installation via CLI
Verify that the operator controller manager pod is up and running:
```
$ oc get pod | grep inteldeviceplugins-controller-manager

inteldeviceplugins-controller-manager-6b8c76c867-hftqm   2/2     Running   0          17m
```

## Resources Provided by Intel Device Plugins
The resources are the user interface for customers to claim and consume the hardware features provided by Intel Device Plugins from the user pods. See below table for the details:

| Feature | Resources | Description | Usage |
| ------- | --------- | ----------- | ----- |
| Intel® SGX | `sgx.intel.com/epc` | Intel SGX EPC memory for user pod to claim | [Link](https://github.com/intel/intel-technology-enabling-for-openshift/blob/64a6c86f3be25459c14ea988e892f9f5d873a8ca/tests/l2/sgx/sgx_job.yaml#L21) |
| Intel® Data Center GPU Flex Series </br> Intel® Data Center GPU Max Series | `gpu.intel.com/i915 ` | Intel Data Center GPU Card for user pod to claim | [Link](https://github.com/intel/intel-technology-enabling-for-openshift/blob/main/device_plugins/deploy_gpu.md#using-intel-data-center-gpu-resource-exclusively) |
| Intel® QAT | `qat.intel.com/cy` </br> `qat.intel.com/dc` | `cy`: Intel QAT VFIO Virtual Function device configured for cryptography for user pod to claim </br> `dc`: Intel QAT VFIO Virtual Function device configured for cryptography for user pod to claim | [Link](https://github.com/intel/intel-technology-enabling-for-openshift/blob/main/tests/l2/qat/qatlib_job.yaml#L24) </br> [Link](https://github.com/intel/intel-technology-enabling-for-openshift/blob/main/tests/l2/qat/qatlib_job.yaml#L28) |


## Creating Intel Device Plugin custom resource (CR)
- To create an Intel SGX device plugin CR, follow this [link](/device_plugins/deploy_sgx.md).
- To create an Intel GPU device plugin CR, follow this [link](/device_plugins/deploy_gpu.md).
- To create an Intel QAT device plugin CR, follow this [link](/device_plugins/deploy_qat.md).
