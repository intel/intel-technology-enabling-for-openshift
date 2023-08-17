# Create Intel QAT Device Plugin CR
Intel® QuickAssist Technology (Intel® QAT) provides cryptographic and compression acceleration capabilities used to improve performance and efficiency across the data center.
For more details see https://intel.github.io/quickassist

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

## verify QAT resource is available
QAT has two kinds of resource cryptography and compression. If the plugin deployment is a success, look for `qat.intel.com` resources for compression and crypto.
```
$ oc describe node <node name> | grep qat.intel.com/cy
 qat.intel.com/cy: 16
 qat.intel.com/cy: 16
```

## Advanced Configuration
By default, half of the virtual functions are used for compression and the other half is used for cryptography. This default behaviour can be customised via a config map which configures the hardware using sysfs. The config values could be `asym`, `sym` and `dc`, for more details see [cfg_services](https://github.com/torvalds/linux/blob/42e66b1cc3a070671001f8a1e933a80818a192bf/Documentation/ABI/testing/sysfs-driver-qat)

To create a configmap, run the following command:
```
oc create configmap --namespace=openshift-operators --from-literal "qat.conf=ServicesEnabled=<option>" qat-config
```
The option can be any combination of values described above.
When using the operator for deploying the plugin with provisioning config, use provisioningConfig field for the name of the ConfigMap, then the config is passed to initcontainer through the volume mount.

There's also a possibility for a node specific congfiguration through passing a nodename via NODE_NAME into initcontainer's environment and passing a node specific profile (qat-$NODE_NAME.conf) via ConfigMap volume mount.
