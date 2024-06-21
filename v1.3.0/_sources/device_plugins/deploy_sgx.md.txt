# Create Intel SGX Device Plugin CR

## Create a CR via web console
1.	Go to **Operator** -> **Installed Operators**.
2.	Open **Intel Device Plugins Operator**.
3.	Navigate to tab **Intel Software Guard Extensions Device Plugin**.
4.	Click **Create SgxDevicePlugin** -> set correct parameters -> Click **Create** 
5.	Optional: If you want to make any customizations, select YAML view and edit the details. When you are done, click **Create**.

## Verify via web console
1.	Verify CR by checking the status of **Workloads** -> **DaemonSet** -> **intel-sgx-plugin**.
2.	Now `SgxDevicePlugin` is created.

## Create CR via CLI
Apply the CR yaml file:
```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/v1.3.0/device_plugins/sgx_device_plugin.yaml
```

## Verify via CLI
Verify that the device plugin CR is ready: 
```
$ oc get SgxDevicePlugin
```
Output: 
```
NAME		        DESIRED		READY	NODE SELECTOR	                                    AGE
sgxdeviceplugin-sample  1 	        1       {"intel.feature.node.kubernetes.io/sgx":"true"}     2m
```