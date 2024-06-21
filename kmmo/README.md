# Setting up Out of Tree Drivers

## Introduction
[Kernel module management (KMM) operator](https://github.com/rh-ecosystem-edge/kernel-module-management) manages the deployment and lifecycle of out-of-tree kernel modules on RHOCP.

In this release, KMM operator is used to manage and deploy the Intel® Data Center GPU driver container image on the RHOCP cluster.

Intel data center GPU driver container images are released from [Intel Data Center GPU Driver for OpenShift Project](https://github.com/intel/intel-data-center-gpu-driver-for-openshift/tree/main/release#intel-data-center-gpu-driver-container-images-for-openshift-release).

## KMM operator working mode
- **Pre-build mode** - This is the default and recommended mode. KMM Operator uses [this pre-built and certified Intel Data Center GPU driver container image](https://catalog.redhat.com/software/containers/intel/intel-data-center-gpu-driver-container/6495ee55c8b2461e35fb8264), which is published on the Red Hat Ecosystem Catalog to provision Intel Data Center GPUs on a RHOCP cluster.
- **On-premises build mode** - Users can optionally build and deploy their own driver container images on-premises through the KMM operator.

## Prerequisites
- Provisioned RHOCP cluster. Follow steps [here](/README.md#provisioning-rhocp-cluster).
- Setup node feature discovery. Follow steps [here](/nfd/README.md).

## Install KMM operator
Follow the installation guide below to install the KMM operator via CLI or web console. 
- [Install from CLI](https://docs.openshift.com/container-platform/4.14/hardware_enablement/kmm-kernel-module-management.html#kmm-install-using-cli_kernel-module-management-operator)
- [Install from web console](https://docs.openshift.com/container-platform/4.14/hardware_enablement/kmm-kernel-module-management.html#kmm-install-using-web-console_kernel-module-management-operator)

## Canary deployment with KMM
Canary deployment is enabled by default to deploy the driver container image only on specific node(s) to ensure the initial deployment succeeds prior to rollout to all the eligible nodes in the cluster. This safety mechanism can reduce risk and prevent a deployment from adversely affecting the entire cluster.

## Set alternative firmware path at runtime with KMM
Follow the steps below to set the alternative firmware path at runtime.

1. Update KMM operator `ConfigMap` to set `worker.setFirmwareClassPath` to `/var/lib/firmware`
``` 
$ oc patch configmap kmm-operator-manager-config -n openshift-kmm --type='json' -p='[{"op": "add", "path": "/data/controller_config.yaml", "value": "healthProbeBindAddress: :8081\nmetricsBindAddress: 127.0.0.1:8080\nleaderElection:\n  enabled: true\n  resourceID: kmm.sigs.x-k8s.io\nwebhook:\n  disableHTTP2: true\n  port: 9443\nworker:\n  runAsUser: 0\n  seLinuxType: spc_t\n  setFirmwareClassPath: /var/lib/firmware"}]'
```

2. Delete the KMM operator controller pod for `ConfigMap` changes to take effect.
``` 
$ oc get pods -n openshift-kmm | grep -i "kmm-operator-controller-" | awk '{print $1}' | xargs oc delete pod -n openshift-kmm
```

For more details, see [link.](https://openshift-kmm.netlify.app/documentation/firmwares/#setting-the-kernels-firmware-search-path)

## Deploy Intel Data Center GPU Driver with pre-build mode
Follow the steps below to deploy the driver container image with pre-build mode.
1.	Find all nodes with an Intel Data Center GPU card using the following command:
``` 
$ oc get nodes -l intel.feature.node.kubernetes.io/gpu=true
```
Example output: 
```
NAME         STATUS   ROLES    AGE   VERSION
icx-dgpu-1   Ready    worker   30d   v1.25.4+18eadca
```

2.	Label the node(s) in the cluster using the command shown below for the initial canary deployment.
```
$ oc label node <node_name> intel.feature.node.kubernetes.io/dgpu-canary=true
```

3.	Use pre-build mode to deploy the driver container.
```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/kmmo/intel-dgpu.yaml   
```

4.	After the driver is verified on the cluster through the canary deployment, simply remove the line shown below from the [`intel-dgpu.yaml`](/kmmo/intel-dgpu.yaml) file and reapply the yaml file to deploy the driver to the entire cluster. As a cluster administrator, you can also select another deployment policy.
```
intel.feature.node.kubernetes.io/dgpu-canary: 'true'
```

## Verification
To verify that the drivers have been loaded, follow the steps below:
1.	List the nodes labeled with `kmm.node.kubernetes.io/openshift-kmm.intel-dgpu.ready` using the command shown below:
```
$ oc get nodes -l kmm.node.kubernetes.io/openshift-kmm.intel-dgpu.ready
```
Example output: 
```
NAME         STATUS   ROLES    AGE   VERSION
icx-dgpu-1   Ready    worker   30d   v1.25.4+18eadca
```
The label shown above indicates that the KMM operator has successfully deployed the drivers and firmware on the node.

2.	If you want to further debug the driver on the node, follow these steps:  
    a. Navigate to the web console (Compute -> Nodes -> Select a node that has the GPU card -> Terminal).  
    b. Run the commands shown below in the web console terminal:  
    ```
    $ chroot /host 
    $ lsmod | grep i915
    ```
    Ensure `i915` and `intel_vsec` are loaded in the kernel, as shown in the output below:
    ```
    i915                   3633152 0
    i915_compat            16384 1 i915
    intel_vsec             16384  1 i915
    intel_gtt              20480  1 i915
    video                  49152  1 i915
    i2c_algo_bit           16384  1 i915
    drm_kms_helper        290816  1 i915
    drm                   589824  3 drm_kms_helper,i915
    dmabuf                 77824  4 drm_kms_helper,i915,i915_compat,dr
    ```
    c. Run dmesg to ensure there are no errors in the kernel message log.

## See Also
