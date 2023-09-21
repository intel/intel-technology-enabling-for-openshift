# Setting up Machine Configuration

# Introduction
Machine configuration operation is used to configure [Red Hat Enterprise Linux CoreOS (RHCOS)](https://docs.openshift.com/container-platform/4.12/architecture/architecture-rhcos.html) on each node in a RHOCP cluster.

[Machine config operator](https://github.com/openshift/machine-config-operator) (MCO) is provided by Red Hat to manage the operating system and machine configuration. In this project through the MCO, cluster administrators can configure and update the kernel to provision Intel Hardware features on the worker nodes.

MCO is one of the technologies used in this project to manage the machine configuration. In current OCP-4.12, MCO might reboot the node to enable the machine configuration. Since rebooting the node is undesirable, alternative machine configuration technologies are under investigation. For more details, see this [issue](https://github.com/intel/intel-technology-enabling-for-openshift/issues/34).  

The best approach is to work with the RHCOS team to push the RHCOS configuration as the default configuration for a RHOCP cluster on [Day 0](https://www.ibm.com/cloud/architecture/content/course/red-hat-openshift-container-platform-day-2-ops/). 

For some general configuration, we recommend you set it up while provisioning the cluster on [Day 1](https://www.ibm.com/cloud/architecture/content/course/red-hat-openshift-container-platform-day-2-ops/).

If the configuration cannot be set as the default setting, we recommend using some operator to set the configuration on the fly without rebooting the node on [Day 2](https://www.ibm.com/cloud/architecture/content/course/red-hat-openshift-container-platform-day-2-ops/).

Any contribution in this area is welcome. 

# Prerequisites 
- Provisioned RHOCP 4.12 cluster. Follow steps [here](/README.md#provisioning-rhocp-cluster).
- Setup node feature discovery (NFD). Follow steps [here](/nfd/README.md).

# General configuration 

## Set up an alternative firmware path for the cluster 
The command below sets `/var/lib/firmware` as the alternative firmware path since the default firmware path `/lib/firmware` is read-only on RHCOS.
``` 
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/machine_configuration/100-alternative-fw-path-for-worker-nodes.yaml
```
**Note**: This command reboots all the worker nodes sequentially. To avoid reboot on Day 2, the cluster administrator can perform this step on Day 1.

## Verification
### From the web console
Navigate to the Compute -> MachineConfigPools section to check the status and make sure the MachineConfigPool update is complete. The status will update from `Updating` to `Up to date`.

### From CLI 
```
$ oc get mcp
```
Output:
```
NAME         CONFIG                                                 UPDATED   UPDATING   DEGRADED   MACHINECOUNT   
worker       rendered-worker-20c8785dee44f52d159fa1c04eeb8552       True      False      False      1              
```

## Verify the alternative firmware path
Navigate to the node terminal on the web console (Compute -> Nodes -> Select a node -> Terminal). Run the following command in the terminal:
```
$ cat /proc/cmdline 
```
Ensure `firmware_class.path=/var/lib/firmware` is present.

# Machine configuration for Intel® Data Center GPU
## Create `intel-dgpu` MachineConfigPool
The command shown below creates a custom `intel-dgpu` MachineConfigPool for worker nodes with an Intel Data Center GPU card, which is labeled with `intel.feature.node.kubernetes.io/gpu: 'true'` by [NFD](/nfd/README.md). 
```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/machine_configuration/intel-dgpu-machine-config-pool.yaml
```

## Verification
### From the web console
Navigate to the Compute -> MachineConfigPools section and ensure `intel-dgpu` MachineConfigPool is present.
### From the CLI 
```
$ oc get mcp
```
Output: 
```
NAME         CONFIG                                                 UPDATED   UPDATING   DEGRADED   MACHINECOUNT   
intel-dgpu   rendered-intel-dgpu-58fb5f4d72fe6041abb066880e112acd   True      False      False      1             
```
Ensure `intel-dgpu` MachineConfigPool is present.

# Disable conflicting driver
Run the command shown below to disable the loading of a potential conflicting driver, such as `ast` driver.

**Note**: The `i915` driver depends on a ported `drm` module. Some other drivers, such as ast that depends on in-tree drm module might have a compatibility issue. The known issue will be resolved on i915 driver for RHEL `9.x`, which will be used for RHOCP `4.13`. 
```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/machine_configuration/100-intel-dgpu-machine-config-disable-ast.yaml
```
**Note**: This command will reboot the worker nodes in the `intel-dgpu` MachineConfigPool sequentially.

## Verification
Navigate to the node terminal on the web console (Compute -> Nodes -> Select a node -> Terminal). Run the following commands in the terminal.
```
$ chroot /host
$ lsmod | grep ast
```
Ensure that ast driver is not loaded.

# Machine Configuration for Provisioning Intel® QAT

* Turn on `intel_iommu` kernel parameter and load `vfio-pci` at boot for QAT provisioning

```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/machine_configuration/100-intel-qat-intel-iommu-on.yaml
```

Note: This will reboot the worker nodes when changing the kernel parameter through MCO.

# See Also
- [Firmware Search Path](https://docs.kernel.org/driver-api/firmware/fw_search_path.html)
- [Red Hat OpenShift Container Platform Day-2 operations](https://www.ibm.com/cloud/architecture/content/course/red-hat-openshift-container-platform-day-2-ops/)
