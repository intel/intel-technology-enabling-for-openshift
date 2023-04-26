## Machine Configuration

Machine Configuration is used to configure the [Red Hat Enterprise Linux CoreOS](https://docs.openshift.com/container-platform/4.12/architecture/architecture-rhcos.html) (RHCOS) on each node in a OCP cluster.

[Machine Config Operator](https://github.com/openshift/machine-config-operator) (MCO) is provided by Red Hat to manage updates and configuration changes.

### Prerequisites:

Please visit [NFD README](/nfd/README.md) first to ensure NFD operator has been installed and configured properly. 

### General Configuration for Provisioning Intel Hardware Features
* Set up an alternative firmware path for the cluster 

The command below sets `/var/lib/firmware` as an alternative firmware path for RHCOS kernel to load OOT firmware. This is needed since the default firmware directory `/lib/firmware` is read-only on OCP cluster nodes.

Note: This command will reboot all the worker nodes sequentially. Rebooting is not preferred by this project. A better solution is being explored.

To reboot only the worker nodes with Intel data center GPU card, add the `selector` field under `spec` and the label `intel.feature.node.kubernetes.io/gpu: 'true'` to this [yaml file](/machine_configuration/100-alternative-fw-path-for-worker-nodes.yaml). Use the below command as a reference to deploy the modified yaml file.

```
$ oc apply -f https://github.com/intel/intel-technology-enabling-for-openshift/blob/main/machine_configuration/100-alternative-fw-path-for-worker-nodes.yaml
```

Once the `worker` MachineConfigPool update is complete, the status will update from `Updating` to `Up to date`. To check the status on the OCP web console, navigate to the Compute -> MachineConfigPools section or run `oc get mcp` via CLI. 

To verify the alternative firmware path has been set correctly on worker nodes, run `cat /proc/cmdline` on the node terminal via OCP web console (Compute -> Nodes -> Select a node -> Terminal) or CLI. Ensure `firmware_class.path=/var/lib/firmware` is present in the command output.

### Machine Configuration for Provisioning Intel Data Center GPU
* Create `intel-dgpu` MachineConfigPool:

The command below creates a custom `intel-dgpu` MachineConfigPool for worker nodes with Intel data center GPU card. These nodes are labeled `intel.feature.node.kubernetes.io/gpu: 'true'` by NFD. 

```
$ oc apply -f https://github.com/intel/intel-technology-enabling-for-openshift/blob/main/machine_configuration/intel-dgpu-machine-config-pool.yaml
```

To verify the MachineConfigPool has been successfully created, navigate to the Compute -> MachineConfigPools section in the web console or run `oc get mcp` via CLI. Ensure `intel-dgpu` MachineConfigPool is present.

* Disable in-tree ast driver 

Due to a known incompatible issue with the Intel i915 driver, the below command is used to blacklist (disable loading of) the ast driver on worker nodes in the `intel-dgpu` MachineConfigPool. This is needed prior to loading Intel data center GPU drivers and firmware. 

Note: The driver that needs to be blacklisted will vary on the hardware platform the user is using. As a result, please use this [yaml file](/machine_configuration/100-intel-dgpu-machine-config-disable-ast.yaml) and below command as a reference for your environment.

Note: This command will reboot all the worker nodes in the `intel-dgpu` MachineConfigPool sequentially. This known incompatible issue will be addressed when Intel data center GPU drivers arrive for RHEL 9.x.

```
$ oc apply -f https://github.com/intel/intel-technology-enabling-for-openshift/blob/main/machine_configuration/100-intel-dgpu-machine-config-disable-ast.yaml
```

To verify the in-tree ast driver has been blacklisted on nodes with Intel data center GPU card, run `chroot /host` and then `lsmod | grep ast` on the node terminal via OCP web console (Compute -> Nodes -> Select a node that has the GPU card -> Terminal) or CLI. Ensure no output is returned by the command. If ast is still loaded, the command output will contain ast.

### Machine Configuration for Provisioning Intel QAT

* Create `intel-qat` machine config pool:

The command below creates a custom `intel-qat` MachineConfigPool for worker nodes with Intel QAT feature. These nodes are labeled `intel.feature.node.kubernetes.io/qat: 'true'` by NFD. 

```
$ oc apply -f https://github.com/intel/intel-technology-enabling-for-openshift/blob/main/machine_configuration/intel-qat-machine-config-pool.yaml
```

To verify the MachineConfigPool has been successfully created, navigate to the Compute -> MachineConfigPools section in the web console or run `oc get mcp` via CLI. Ensure `intel-qat` MachineConfigPool is present.

* Turn on `intel_iommu` kernel boot parameter for QAT

For QAT, `intel_iommu` is a kernel boot parameter that needs to be turned on using the following command.

Note: This command will reboot the worker nodes in the `intel-qat` MachineConfigPool sequentially to turn on the `intel_iommu` kernel parameter. Rebooting is not preferred by this project. A better solution is being explored.

```
$ oc apply -f https://github.com/intel/intel-technology-enabling-for-openshift/blob/main/machine_configuration/100-intel-qat-intel-iommu-on.yaml
```

To verify the `intel_iommu` kernel parameter is turned on in nodes with QAT support, run `cat /proc/cmdline` on the node terminal via OCP web console (Compute -> Nodes -> Select a node -> Terminal) or CLI. Ensure `intel_iommu=on` is present in the command output.
