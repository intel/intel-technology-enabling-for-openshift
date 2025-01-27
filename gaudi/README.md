# Setting up Intel Gaudi AI Accelerator Operator

## Overview
[Intel Gaudi AI Accelerator Operator](https://catalog.redhat.com/software/container-stacks/detail/6683b2cce45daa25e36bddcb) is used to provision Intel Gaudi Accelerator with OpenShift. The steps and yaml files mentioned in this document to provision the Gaudi accelerator are based on [Intel Gaudi AI Accelerator Operator for OpenShift](https://docs.habana.ai/en/latest/Orchestration/Intel_Gaudi_Base_Operator/index.html).

If you are familiar with the steps here to manually provision the accelerator, the Red Hat certified Operator and Ansible based [One-Click](/one_click/README.md#reference-playbook-–-habana-gaudi-provisioning) solution can be used as a reference to provision the accelerator automatically.

## Prerequisities
- To Provision RHOCP cluster, follow steps [here](/README.md#provisioning-rhocp-cluster).

## Install Intel Gaudi AI Accelerator Operator on Red Hat OpenShift
### Installation via web console
Follow the steps below to install Intel Gaudi AI Accelerator Operator using OpenShift web console:
1.	In the OpenShift web console, navigate to **Operator** -> **OperatorHub**.
2.	Search for **Intel Gaudi AI Accelerator Operator** in all items field -> Click **Install**.
### Verify Installation via web console
1.	Go to **Operator** -> **Installed Operators**.
2.	Verify that the status of the operator is **Succeeded**.

### Installation via Command Line Interface (CLI)
```
oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/gaudi/gaudi_install_operator.yaml
```

### Verify Installation via CLI
Verify that the operator controller manager pod is up and running:
```
oc get pods -n habana-ai-operator

NAME                                  READY   STATUS    RESTARTS   AGE
controller-manager-6c8459d9cb-fqs8h   2/2     Running   0          25m
```

## Creating Intel Gaudi AI Accelerator Operator ClusterPolicy Instance
To create a Habana Gaudi device plugin CR, follow the steps below.

### Create CR via web console
1.	Go to **Operator** -> **Installed Operators**.
2.	Open **Intel Gaudi AI Accelerator Operator**.
3.	Navigate to tab **Cluster Policy**.
4.	Click **Create ClusterPolicy** -> set correct parameters -> Click **Create**. To set correct parameters please refer [Using RedHat OpenShift Container Platform Console](https://docs.habana.ai/en/latest/Installation_Guide/Additional_Installation/Kubernetes_Installation/Kubernetes_Operator.html#id1).

### Verify via web console
1.	Verify CR by checking the status of **Workloads** -> **DaemonSet** -> **habana-ai-device-plugin-ds**,  **habana-ai-driver-rhel-9-4-xxxxx**, **habana-ai-feature-discovery-ds**, **habana-ai-metric-exporter-ds**, **habana-ai-runtime-ds**.
2.	Now `ClusterPolicy` is created.

### Create CR via CLI
Apply the CR yaml file:
```
oc apply -f  https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/gaudi/gaudi_cluster_policy.yaml
```

### Verify the ClusterPolicy CR is created
You can use command below to verify that the `ClusterPolicy` CR has been created:
```
oc get pod -n habana-ai-operator

NAME                                                       READY   STATUS    RESTARTS      AGE
habana-ai-device-plugin-ds-thj7b                           1/1     Running   0             10d
habana-ai-driver-rhel-9-4-416-94-202412170927-0-ds-vqhzb   1/1     Running   2             10d
habana-ai-feature-discovery-ds-ztl2j                       1/1     Running   5             10d
habana-ai-metric-exporter-ds-g5qqh                         1/1     Running   0             10d
habana-ai-operator-controller-manager-6c995b5646-wl7cp     2/2     Running   0             10d
habana-ai-runtime-ds-x49lf                                 1/1     Running   0             10d
```
Alternatively, you can also check the status of the `ClusterPolicy` CR like below: 
```
oc describe ClusterPolicy habana-ai -n habana-ai-operator

Name:         habana-ai
Namespace:    habana-ai-operator
.
.
Status:
  Conditions:
    Last Transition Time:  2025-01-21T18:50:46Z
    Message:               All resources have been successfully reconciled
    Reason:                Reconciled
    Status:                True
```
## Verify Gaudi Provisioning
After the `ClusterPolicy` instance CR is created, it will take some time for the operator to download the Gaudi OOT driver source code and build it on-premise with the help of the KMM operator. The OOT driver module binaries will be loaded into the RHCOS kernel on each node with Gaudi cards labelled by feature discovery. Then, the Gaudi device plugin can advertise the Gaudi resources listed in the table for the pods on OpenShit to use. Run the command below to check the availability of Gaudi resources:
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