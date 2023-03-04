## Intel Device Plugins for OpenShift 

### overview

Intel Device Plugins are used to advertise Intel hardware features (resources) to OpenShift Cluster, so the workload running on pod deployed on the Clusters can use these features. To manage deployment and lifecycle of these device plugins, Intel device plugin operator is used. Intel Device Plugin Container images and operator has been certificated and released on Red Hat Ecosystem Catalog since OCP-4.10 release. Please refer to [Intel Device Plugins for Kubernetes project](https://github.com/intel/intel-device-plugins-for-kubernetes) for detail

### Deploy Intel Device Plugins on OpenShift

### Deploy Intel Device Plugins on OpenShift thourgh web Console
Below operations have been verified on OCP-4.11 and OCP-4.12

* Install Intel Device Plugin operator 

```
$ operator-sdk run bundle quay.io/ocpeng/intel-device-plugins-operator-bundle:0.26.0
```

* Verify Operator installation
1.  On the OpenShift console, go to **Operator** -> **Installed Operators**
2.  Verify the status of Intel Device Plugins Operator as **Succeeded**

* Deploy Intel SGX Device Plugin
Follow the steps below to deploy Intel SGX Device Plugin Custom Resource
1.	Go to **Operator** -> **Installed Operators**
2.  Open **Intel Device Plugins Operator**
3.  Navigate to tab **Intel Software Guard Extensions Device Plugin**
4.  Click **Create SgxDevicePlugin ->** set correct parameters -> Click **Create** 
    OR for any customizations, please select `YAML view` and edit details. Once done, click **Create**  
5.  Verify CR by checking the status of DaemonSet **`intel-sgx-plugin`**
6.  Now `SgxDevicePlugin` is ready to deploy any workloads

* Deploy dGPU Device Plugin
Follow the steps below to deploy Intel GPU Device Plugin Custom Resource
1.	Go to **Operator** -> **Installed Operators**
2.  Open **Intel Device Plugins Operator**
3.  Navigate to tab **Intel GPU Device Plugin**
4.  Click **Create GpuDevicePlugin ->** set correct parameters -> Click **Create** 
    OR for any customizations, please select `YAML view` and edit details. Once done, click **Create**  
5.  Verify CR by checking the status of DaemonSet **`intel-gpu-plugin`**
6.  Now `GpuDevicePlugin` is ready to deploy any workloads

* Deploy Intel QAT Device Plugin
Follow the steps below to deploy Intel QAT Device Plugin Custom Resource
1.	Go to **Operator** -> **Installed Operators**
2.  Open **Intel Device Plugins Operator**
3.  Navigate to tab **Intel QuickAssist Technology Device Plugin**
4.  Click **Create QatDevicePlugin ->** set correct parameters -> Click **Create** 
    OR for any customizations, please select `YAML view` and edit details. Once done, click **Create**  
5.  Verify CR by checking the status of DaemonSet **`intel-qat-plugin`**
6.  Now `QatDevicePlugin` is ready to deploy any workloads
