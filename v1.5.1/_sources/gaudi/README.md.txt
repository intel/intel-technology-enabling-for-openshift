# Setting up Intel Gaudi Base Operator

## Overview
[Intel Gaudi Base Operator](https://catalog.redhat.com/software/container-stacks/detail/6683b2cce45daa25e36bddcb) is used to provision Intel Gaudi Accelerator with OpenShift. The steps and yaml files mentioned in this document to provision the Gaudi accelerator are based on [Intel Gaudi Base Operator for OpenShift](https://docs.habana.ai/en/latest/Orchestration/Intel_Gaudi_Base_Operator/index.html).

If you are familiar with the steps here to manually provision the accelerator, the Red Hat certified Operator and Ansible based [One-Click](/one_click/README.md#reference-playbook-–-habana-gaudi-provisioning) solution can be used as a reference to provision the accelerator automatically.

## Prerequisities
- To Provision RHOCP cluster, follow steps [here](/README.md#provisioning-rhocp-cluster).
- To Install NFD Operator, follow steps [here](/nfd/README.md#install-nfd-operator).
- To Install KMM Operator, follow steps [here](/kmmo/README.md#install-kmm-operator).

## Update Kernel Firmware Search Path with MCO
**Note:** This step will reboot the nodes, it is recommended to do this in the first step.

The default kernel firmware search path `/lib/firmware` in RHCOS is not writable. Command below can be used to add path `/var/lib/fimware` into the firmware search path list.
```
oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/v1.5.1/gaudi/gaudi_firmware_path.yaml
```

## Label Gaudi Accelerator Nodes With NFD
NFD operator can be used to configure NFD to automatically detect the Gaudi accelerators and label the nodes for the following provisioning steps.
```
oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/v1.5.1/gaudi/gaudi_nfd_instance_openshift.yaml
```
Verify NFD has labelled the node correctly:
```
oc get no -o json | jq '.items[].metadata.labels' | grep pci-1da3

"feature.node.kubernetes.io/pci-1da3.present": "true",
```
NFD detects underlying Gaudi Accelerator using its PCI device class and the vendor ID.

## Install Intel Gaudi Base Operator on Red Hat OpenShift
### Installation via web console
Follow the steps below to install Intel Gaudi Base Operator using OpenShift web console:
1.	In the OpenShift web console, navigate to **Operator** -> **OperatorHub**.
2.	Search for **Intel Gaudi Base Operator** in all items field -> Click **Install**.
### Verify Installation via web console
1.	Go to **Operator** -> **Installed Operators**.
2.	Verify that the status of the operator is **Succeeded**.

### Installation via Command Line Interface (CLI)
```
oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/v1.5.1/gaudi/gaudi_install_operator.yaml
```

### Verify Installation via CLI
Verify that the operator controller manager pod is up and running:
```
oc get pods -n habana-ai-operator

NAME                                  READY   STATUS    RESTARTS   AGE
controller-manager-6c8459d9cb-fqs8h   2/2     Running   0          25m
```

## Creating Intel Gaudi Base Operator DeviceConfig Instance
To create a Habana Gaudi device plugin CR, follow the steps below.

### Create CR via web console
1.	Go to **Operator** -> **Installed Operators**.
2.	Open **Intel Gaudi Base Operator**.
3.	Navigate to tab **Device Config**.
4.	Click **Create DeviceConfig** -> set correct parameters -> Click **Create**. To set correct parameters please refer [Using RedHat OpenShift Console](https://docs.habana.ai/en/latest/Installation_Guide/Additional_Installation/Intel_Gaudi_Base_Operator/Deploying_Intel_Gaudi_Base_Operator.html?highlight=openshift#id2).

### Verify via web console
1.	Verify CR by checking the status of **Workloads** -> **DaemonSet** -> **habana-ai-module-device-plugin-xxxxx**.
2.	Now `DeviceConfig` is created.

### Create CR via CLI
Apply the CR yaml file:
```
oc apply -f  https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/v1.5.1/gaudi/gaudi_device_config.yaml
```

### Verify the DeviceConfig CR is created
You can use command below to verify that the `DeviceConfig` CR has been created:
```
oc get pod -n habana-ai-operator

NAME                                         READY   STATUS    RESTARTS        AGE
controller-manager-6586758d54-qw644          2/2     Running   0            5d5h
habana-ai-habana-runtime-bqpvp               1/1     Running   0            5d6h
habana-ai-module-device-plugin-pljkf-kxgdj   1/1     Running   0            5d6h
habana-ai-node-metrics-rghlr                 1/1     Running   0            5d6h
```
Alternatively, you can also check the status of the `DeviceConfig` CR like below: 
```
oc describe deviceconfig habana-ai -n habana-ai-operator

Name:         habana-ai
Namespace:    habana-ai-operator
.
.
Status:
  Conditions:
    Last Transition Time:  2024-07-24T14:05:11Z
    Message:               All resources have been successfully reconciled
    Reason:                Reconciled
    Status:                True
```
## Verify Gaudi Provisioning
After the `DeviceConfig` instance CR is created, it will take some time for the operator to download the Gaudi OOT driver source code and build it on-premise with the help of the KMM operator. The OOT driver module binaries will be loaded into the RHCOS kernel on each node with Gaudi cards labelled by NFD. Then, the Gaudi device plugin can advertise the Gaudi resources listed in the table for the pods on OpenShit to use. Run the command below to check the availability of Gaudi resources:
```
oc describe node | grep habana.ai/gaudi

  habana.ai/gaudi:    8 -> Gaudi cards number on the cluster
  habana.ai/gaudi:    8 -> Gaudi cards number allocatble on the cluster
  habana.ai/gaudi    4       4 -> number of Gaudi cards allocated and number of Gardi cards available
```

To view the metrics on a node with Gaudi card, refer [Collecting Metrics](https://docs.habana.ai/en/latest/Orchestration/Prometheus_Metric_Exporter.html?highlight=metrics#collecting-metrics).

## Resources Provided by Habana Gaudi Device Plugin
The resources provided are the user interface for customers to claim and consume the hardware features from the user pods. See below table for the details:

| Feature | Resources | Description |
| ------- | --------- | ----------- |
| Habana Gaudi | `habana.ai/gaudi` | Number of Habana Gaudi Card resources ready to claim | 

## Upgrade Intel Gaudi SPI Firmware
Refer [Upgrade Intel Gaudi SPI Firmware](/gaudi/Gaudi-SPI-Firmware-Upgrade.md) to upgrade the SPI Firmware on Intel Gaudi.
