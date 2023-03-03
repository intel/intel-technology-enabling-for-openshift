# IDPO - Intel Device Plugins Operator 
## overview
To simplify the deployment of the Intel device plugins, a unified device plugins operator is implemented.

Currently the operator has support for Intel SGX, Intel dGPU, and Intel QAT device plugins. Each device plugin has its own custom resource definition (CRD) and the corresponding controller that watches CRUD operations to those custom resources.

This README gives the installation and usage details of the Intel Device Plugins Operator available for Red Hat OpenShift Container Platform.


## Installation

### Prerequisites 
- Make sure Red Hat OpenShift Cluster is ready to use and the developing machine is RHEL and `oc` command is installed and configured properly. Please note that the following operation is verified on Red Hat OpenShift Cluster 4.11 and working machine RHEL-8.6
- Install the `oc` command to your development machine
- Development machine configurations tested on -  

    | Tool        | Version  |
    | ------------- |-------------|
    |go | go1.19 |
    |kubectl | v1.24.3 |
    |operator-sdk | v1.22.2 |
- Follow the [link](https://docs.openshift.com/container-platform/4.10/hardware_enablement/psap-node-feature-discovery-operator.html) to install **NFD operator** (if it's not already installed).   
    **Note:** Please only install the NFD operator and use steps below to create the NodeFeatureDiscovery instance.  
    - Create the NodeFeatureDiscovery instance  
    ```
    $ oc apply -f https://raw.githubusercontent.com/intel/intel-device-plugins-for-kubernetes/v0.24.0/deployments/nfd/overlays/node-feature-discovery/node-feature-discovery-openshift.yaml
    ```
    - Create the NodeFeatureRule instance  
    **Note:** By default, QAT device pci id supported are `["37c8", "4940"]`. If the pci id of the underlying QAT card is different, please add the entry to QAT device value in the file `/deployments/nfd/overlays/node-feature-rules/node-feature-rules-openshift.yaml`.
    ``` 
    $ oc apply -f https://raw.githubusercontent.com/intel/intel-device-plugins-for-kubernetes/v0.24.0/deployments/nfd/overlays/node-feature-rules/node-feature-rules-openshift.yaml
    ```

## Intel Device Plugins Operator Installation
Install the operator using below command - 
```
$ operator-sdk run bundle quay.io/ocpeng/intel-device-plugins-operator-bundle:0.26.0
```

## Verify Operator installation
1.  On the OpenShift console, go to **Operator** -> **Installed Operators**
2.  Verify the status of Intel Device Plugins Operator as **Succeeded**

## Deploying Intel Device Plugins

### Intel SGX Device Plugin
Follow the steps below to deploy Intel SGX Device Plugin Custom Resource
1.	Go to **Operator** -> **Installed Operators**
2.  Open **Intel Device Plugins Operator**
3.  Navigate to tab **Intel Software Guard Extensions Device Plugin**
4.  Click **Create SgxDevicePlugin ->** set correct parameters -> Click **Create** 
    OR for any customizations, please select `YAML view` and edit details. Once done, click **Create**  
5.  Verify CR by checking the status of DaemonSet **`intel-sgx-plugin`**
6.  Now `SgxDevicePlugin` is ready to deploy any workloads

### Intel dGPU Device Plugin
Follow the steps below to deploy Intel GPU Device Plugin Custom Resource
1.	Go to **Operator** -> **Installed Operators**
2.  Open **Intel Device Plugins Operator**
3.  Navigate to tab **Intel GPU Device Plugin**
4.  Click **Create GpuDevicePlugin ->** set correct parameters -> Click **Create** 
    OR for any customizations, please select `YAML view` and edit details. Once done, click **Create**  
5.  Verify CR by checking the status of DaemonSet **`intel-gpu-plugin`**
6.  Now `GpuDevicePlugin` is ready to deploy any workloads

### Intel QAT Device Plugin
Follow the steps below to deploy Intel QAT Device Plugin Custom Resource
1.	Go to **Operator** -> **Installed Operators**
2.  Open **Intel Device Plugins Operator**
3.  Navigate to tab **Intel QuickAssist Technology Device Plugin**
4.  Click **Create QatDevicePlugin ->** set correct parameters -> Click **Create** 
    OR for any customizations, please select `YAML view` and edit details. Once done, click **Create**  
5.  Verify CR by checking the status of DaemonSet **`intel-qat-plugin`**
6.  Now `QatDevicePlugin` is ready to deploy any workloads