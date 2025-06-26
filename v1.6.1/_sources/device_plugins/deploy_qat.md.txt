# Create Intel QAT Device Plugin CR

## Create a CR via web console
1.	Go to **Operator** -> **Installed Operators**.
2.	Open **Intel Device Plugins Operator**.
3.	Navigate to tab **Intel QuickAssist Technology Device Plugin**.
4.	Click **Create QatDevicePlugin** -> set correct parameters -> Click **Create** 
5.	Optional: If you want to make any customizations, select YAML view and edit the details. When you are done, click **Create**.

## Verify via web console
1.	Verify CR by checking the status of **Workloads** -> **DaemonSet** -> **intel-qat-plugin**.
2.	Now `QatDevicePlugin` is created.

## Create CR via CLI
Apply the CR yaml file:
```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/v1.6.1/device_plugins/qat_device_plugin.yaml
```

## Verify via CLI
Verify that the device plugin CR is ready: 
```
$ oc get QatDevicePlugin
```
Output: 
```
NAME		        DESIRED		READY	NODE SELECTOR	                                    AGE
qatdeviceplugin-sample  1 	        1       {"intel.feature.node.kubernetes.io/qat":"true"}     3h27m
```

# Verify QAT Device Plugin 
After the plugin is deployed, use below command to verify QAT resources: 
```
$ oc describe node <node name> | grep qat.intel.com  
 qat.intel.com/cy: 32 
 qat.intel.com/cy: 32 
 qat.intel.com/dc: 32 
 qat.intel.com/dc: 32 
 ```
**Note**: By default the device plugin registers half resources each for `qat.intel.com/cy` and `qat.intel.com/dc` respectively. For more details about the QAT resources configuration, please refer to the QAT Device Plugin Configuration section below.

# QAT Device Plugin Configuration
> **Note**: The QAT device plugin can be configured with the flags. In this release, only the configurations in the table below are verified and supported on RHOCP. 

For more details about the QAT device plugin configuration flags, see [Modes and Configurations Options](https://github.com/intel/intel-device-plugins-for-kubernetes/blob/main/cmd/qat_plugin/README.md#modes-and-configuration-options).

| Flag | Configuration | Description |
| ---- | ---- | ---- |
| `-dpdk-driver` | vfio-pci | Using vfio-pci driver to manage QAT VFIO device. See details [here](https://doc.dpdk.org/guides/linux_gsg/linux_drivers.html) |
| `-kernel-vf-drivers` | 4xxxvf | Supporting 4xxx QAT device </br> **Note**: Verified on 4th Gen Intel® Xeon® Scalable processors. See details [here](https://github.com/intel/qatlib/blob/main/INSTALL#L72) |
| `-max-num-devices ` | 128 | It is the maximum VF device it can support for 4xxx QAT device. If the number exceeds the maximum number the QAT device supports, then the maximum number will be enabled. |
| `-provisioning-config ` | Name of ConfigMap | See section [QAT resource configuration](/device_plugins/deploy_qat.md#qat-resource-configuration-experimental)  |

## QAT Resource Configuration (experimental)

**NOTE**: In this release, this is an experimental feature. The efforts to [enhance this feature](https://github.com/intel/intel-device-plugins-for-kubernetes/issues/1529) and [make it more stable](https://github.com/intel/intel-device-plugins-for-kubernetes/issues/1542) are on going.

Users can use the steps below to customize the QAT resource configuration:  
1. Create the configmap for qat resource configuration 
    ```
    $ oc create configmap --namespace=openshift-operators --from-literal "qat.conf=ServicesEnabled=<option>" <name-of-configmap> 
    ```
    Options- (refer to [link](https://www.kernel.org/doc/Documentation/ABI/testing/sysfs-driver-qat) for more details):  
    `dc` : Configure all the QAT VF devices managed by the device plugin CR for compression services. The resource created is `qat.intel.com/dc`.

    `sym;asym` and `asym;sym`: Configure all the QAT VF devices managed by the device plugin CR for crypto services. The resource created is `qat.intel.com/cy`.

    `sym;dc` and `dc;sym` : Configure all the QAT VF devices managed by the device plugin CR for symmetric crypto and compression services. The resource created is `qat.intel.com/sym-dc`.

    `asym;dc` and `dc;asym`: Configure all the QAT VF devices managed by the device plugin CR for asymmetric crypto and compression services. The resource created is `qat.intel.com/asym-dc`.

    `sym`: Configure all the QAT VF devices managed by the device plugin CR for running symmetric crypto services. The resource created is `qat.intel.com/sym`. 

    `asym`: Configure all the QAT VF devices managed by the device plugin CR for running asymmetric crypto services. The resource created is `qat.intel.com/asym`. 


2. Create QAT device plugin CR with -provisioning-config set as the name of the ConfigMap (created in step 1) in the qat_device_plugin.yaml file or set ConfigMap name in the provisioning-config option from web console. 

# Multiple Custom Resources

The [feature](https://github.com/intel/intel-device-plugins-for-kubernetes/tree/main/cmd/operator#multiple-custom-resources) can be used with a `nodeSelector` label representing the capabilities supported on the node. Multiple custom resources on nodes can be supported with different QAT capabilities.

* For example, to assign the capabilities `sym` and `asym` for two nodes: create the labels on the nodes
```
oc label node <node1_name> qat.mode=sym
oc label node <node2_name> qat.mode=asym
```
* Create a separate `QatDevicePlugin` by adding each of the above labels in `nodeSelector` and choose the configmaps created with `sym` and `asym` capabilities respectively. Please refer to [QAT Resource Configuration](#qat-resource-configuration-experimental).

`qatdeviceplugin-sym`
```
nodeSelector:
    qat.mode: sym
```
`qatdeviceplugin-asym`
```
nodeSelector:
    qat.mode: asym
```

* Verify that the device plugin CRs are ready
```
$ oc get QatDevicePlugin
```
Output: 
```
NAME		        DESIRED		READY	NODE SELECTOR	                                    AGE
qatdeviceplugin-sym  1 	        1       {"intel.feature.node.kubernetes.io/qat":"true","qat.mode":"sym"}     72m
qatdeviceplugin-asym  1 	        1       {"intel.feature.node.kubernetes.io/qat":"true","qat.mode":"asym"}     71m
```
For more information about the CR and the node it is using:
```
$ oc get QatDevicePlugin -o=jsonpath='{range .items[*]}{.metadata.name}{"\t"}{.status.nodeNames}{"\n"}{end}'
```
Output:
```
qatdeviceplugin-sym       ["node1"]
qatdeviceplugin-asym      ["node2"]
```
* Check the resources on the nodes:
```
oc describe <node1_name> | grep qat.intel.com
    qat.intel.com/sym:                128
    qat.intel.com/sym:                128
    qat.intel.com/sym                0            0

oc describe <node2_name> | grep qat.intel.com
    qat.intel.com/asym:                128
    qat.intel.com/asym:                128
    qat.intel.com/asym                0            0
```

# Run Intel QAT based workloads on RHOCP
To run the Intel QAT based workloads as an unprivileged pod (see [issue](https://github.com/intel/intel-technology-enabling-for-openshift/issues/122)). The customized `qat-scc` Security Context Constraint (SCC) is provided to bind with service account and run the QAT based workload. 

See [Verify Intel QuickAssist Technology Provisioning](/tests/l2/README.md#verify-intel-quickassist-technology-provisioning) for the detailed steps.  
