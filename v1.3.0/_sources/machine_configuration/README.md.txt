# Setting up Machine Configuration

## Introduction
Machine configuration operation is used to configure [Red Hat Enterprise Linux CoreOS (RHCOS)](https://docs.openshift.com/container-platform/4.14/architecture/architecture-rhcos.html) on each node in a RHOCP cluster.

[Machine config operator](https://github.com/openshift/machine-config-operator) (MCO) is provided by Red Hat to manage the operating system and machine configuration. In this project through the MCO, cluster administrators can configure and update the kernel to provision Intel Hardware features on the worker nodes.

MCO is one of the technologies used in this project to manage the machine configuration. In current OCP, MCO might reboot the node to enable the machine configuration. Since rebooting the node is undesirable, alternative machine configuration technologies are under investigation. For more details, see this [issue](https://github.com/intel/intel-technology-enabling-for-openshift/issues/34).  

The best approach is to work with the RHCOS team to push the RHCOS configuration as the default configuration for a RHOCP cluster on [Day 0](https://www.ibm.com/cloud/architecture/content/course/red-hat-openshift-container-platform-day-2-ops/). 

For some general configuration, we recommend you set it up while provisioning the cluster on [Day 1](https://www.ibm.com/cloud/architecture/content/course/red-hat-openshift-container-platform-day-2-ops/).

If the configuration cannot be set as the default setting, we recommend using some operator to set the configuration on the fly without rebooting the node on [Day 2](https://www.ibm.com/cloud/architecture/content/course/red-hat-openshift-container-platform-day-2-ops/).

Any contribution in this area is welcome. 

## Prerequisites 
- Provisioned RHOCP cluster. Follow steps [here](/README.md#provisioning-rhocp-cluster).
- Setup node feature discovery (NFD). Follow steps [here](/nfd/README.md).

## Machine Configuration for Provisioning Intel® QAT

* Turn on `intel_iommu` kernel parameter and load `vfio_pci` at boot for QAT provisioning

```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/v1.3.0/machine_configuration/100-intel-qat-intel-iommu-on.yaml
```

Note: This will reboot the worker nodes when changing the kernel parameter through MCO.

## Verification
Navigate to the node terminal on the web console (Compute -> Nodes -> Select a node -> Terminal). Run the following commands in the terminal.
```
$ cat /proc/cmdline
```
Ensure that `intel_iommu=on` is present.

```
$ chroot /host
$ lsmod | grep vfio_pci
```
Ensure that `vfio_pci` driver is present.

## See Also
- [Red Hat OpenShift Container Platform Day-2 operations](https://www.ibm.com/cloud/architecture/content/course/red-hat-openshift-container-platform-day-2-ops/)
