# Create Intel GPU Device Plugin CR

## Create CR via web console
1.	Go to **Operator** -> **Installed Operators**.
2.	Open **Intel Device Plugins Operator**.
3.	Navigate to tab **Intel GPU Device Plugin**.
4.	Click **Create GpuDevicePlugin** -> set correct parameters -> Click **Create**.
5.	Optional: If you want to  make any customizations, select YAML view and edit the details. Once you are done, click **Create**.

## Verify via web console
1.	Verify CR by checking the status of **Workloads** -> **DaemonSet** -> **intel-gpu-plugin**.
2.	Now `GpuDevicePlugin` is created.

## Create CR via CLI
Apply the CR yaml file:
```
$ oc apply -f  https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/device_plugins/gpu_device_plugin.yaml
```

## Verify via CLI
Verify that the device plugin CR is ready:
```
$ oc get GpuDevicePlugin 
```
Output: 
```
NAME		        DESIRED	  READY	  NODE SELECTOR	                                    AGE
gpudeviceplugin-sample  1         1       {"intel.feature.node.kubernetes.io/gpu":"true"}   3m12s
```

# Using Intel Data Center GPU resource exclusively
In this release, we only verified and support the single Intel GPU `i915` resource dedicated to the single workload pod. To achieve this, we set `sharedDevNum: 1` and `preferredAllocationPolicy: none` as default options.   
As the cluster administrator, use the [gpu_device_plugin.yaml](/device_plugins/gpu_device_plugin.yaml) provided from the previous section Create CR via CLI or use the default options from Create CR via web Console.  
As the application owner, when claiming the i915 resource, make sure the resource limits and requests are set as shown below:
```
spec:
  containers:
    - name: gpu-pod
      resources:
        limits:
          gpu.intel.com/i915: 1
        requests:
          gpu.intel.com/i915: 1
```
For more details, please refer to this [issue](https://github.com/intel/intel-device-plugins-for-kubernetes/issues/1408).  