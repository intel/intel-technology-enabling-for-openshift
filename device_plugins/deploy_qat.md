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
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/device_plugins/qat_device_plugin.yaml
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

In this release, if the user does not configure the QAT resources through the device plugin `-provisioning-config` flag. The device plugin will configure half of the QAT VFIO VF devices for compression/decompression and the other half for cryptography.

Users can use the steps below to customize the QAT resource configuration:  
1. Create the configmap for qat resource configuration 
    ```
    $ oc create configmap --namespace=openshift-operators --from-literal "qat.conf=ServicesEnabled=<option>" Name of ConfigMap 
    ```
    Options:  
    `dc`: Configure all the QAT VF devices managed by the device plugin CR for compression/decompression.  
    `sym;asym`: Configure all the QAT VF devices managed by the device plugin CR for cryptography 
2. Create QAT device plugin CR with -provisioning-config set as the name of the ConfigMap (created in step 1) in the qat_device_plugin.yaml file or set ConfigMap name in the provisioning-config option from web console. 

# Run Intel QAT based workloads on RHOCP
To run the Intel QAT based workloads as an unprivileged pod (see [issue](https://github.com/intel/intel-technology-enabling-for-openshift/issues/122)). The customized `qat-scc` Security Context Constraint (SCC) is provided to bind with service account and run the QAT based workload. 

See [Verify Intel QuickAssist Technology Provisioning](/tests/l2/README.md#verify-intel-quickassist-technology-provisioning) for the detailed steps.  
