# Setting up Intel Device Plugins Operator 

# Overview
Intel Device Plugins are utilized to advertise Intel hardware features (resources) to a Red Hat OpenShift Container Platform (RHOCP) Cluster. This allows workloads running on pods deployed within the clusters to leverage these features. To handle the deployment and lifecycle of these device plugins, the [Intel Device Plugins Operator](https://catalog.redhat.com/software/container-stacks/detail/61e9f2d7b9cdd99018fc5736) is used. The Intel Device Plugins container images and operator have been officially certified and published on the [Red Hat Ecosystem Catalog](https://catalog.redhat.com/software/container-stacks/detail/61e9f2d7b9cdd99018fc5736). For more details on the upstream project, please refer to [Intel Device Plugins for Kubernetes](https://github.com/intel/intel-device-plugins-for-kubernetes).  

# Prerequisities
- Provisioned RHOCP 4.12 cluster. Follow steps [here](/README.md).
- Setup Node Feature Discovery (NFD). Follow steps [here](/nfd/README.md).
- Setup Machine Configuration. Follow steps [here](/machine_configuration/README.md).
- Optional: Setup out-of-tree drivers for Intel Data Center GPU provisioning. Follow the steps listed [here](/kmmo/README.md).

# Install Intel Device Plugins Operator on Red Hat OpenShift
## Installation via web console
Follow the steps below to install Intel Device Plugins Operator using OpenShift web console:
1.	In the OpenShift web console, navigate to **Operator** -> **OperatorHub**.
2.	Search for **Intel Device Plugins Operator** in all items field -> Click **Install**.
## Verify Installation via web console
1.	Go to **Operator** -> **Installed Operators**.
2.	Verify that the status of the operator is **Succeeded**.

## Installation via command line interface (CLI)
Apply the [install_operator.yaml](/device_plugins/install_operator.yaml) file:
```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/device_plugins/install_operator.yaml
```

## Verify Installation via CLI
Verify that the operator controller manager pod is up and running:
```
$ oc get pod | grep inteldeviceplugins-controller-manager

inteldeviceplugins-controller-manager-6b8c76c867-hftqm   2/2     Running   0          17m
```

# Creating Intel Device Plugin custom resource (CR)
- To create an Intel SGX device plugin CR, follow this [link](/device_plugins/deploy_sgx.md):
- To create an Intel GPU device plugin CR, follow this [link](/device_plugins/deploy_gpu.md):
