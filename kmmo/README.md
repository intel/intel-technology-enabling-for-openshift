## Kernel Module Management (KMM)
[Kernel Module Management (KMM) Operator](https://github.com/rh-ecosystem-edge/kernel-module-management) manages the deployment and lifecycle of out-of-tree kernel modules on Red Hat OpenShift Container Platform (OCP).

In this project, KMM Operator is used to manage and deploy the Intel data center GPU driver container image on day 2.

Intel data center GPU driver container images are released from [Intel Data Center GPU Driver for OpenShift Project](https://github.com/intel/intel-data-center-gpu-driver-for-openshift/tree/main/release#intel-data-center-gpu-driver-container-images-for-openshift-release).

### KMM Operator Working Mode 

* [Pre-build mode](/kmmo/intel-dgpu.yaml): This is the **default and recommended** mode. KMM will use the pre-built, certified Intel data center GPU driver container image from Red Hat Ecosystem Catalog to provision Intel data center GPU on OCP.

* [On-premise build mode](/kmmo/intel-dgpu-on-premise-build-mode.yaml): Users can optionally build and deploy their own driver container image on-premise.

### Prerequisites

* Please visit [NFD README](/nfd/README.md) first and then [Machine Configuration README](/machine_configuration/README.md).

### Canary Deployment on OpenShift

* Canary deployment is enabled by default to deploy the driver container image only on specific node(s) to ensure the initial deployment succeeds prior to rollout to all the eligible nodes in the cluster. This safety mechanism can reduce risk and prevent a deployment from affecting the entire cluster.

Note: If you would like to skip canary deployment, simply remove `intel.feature.node.kubernetes.io/dgpu-canary: 'true'` in the [Pre-build mode](/kmmo/intel-dgpu.yaml) or [On-premise build mode](/kmmo/intel-dgpu-on-premise-build-mode.yaml) yaml file.

### Deploying Pre-build Mode with KMM Operator

1. Follow [KMM operator installation guide](https://docs.openshift.com/container-platform/4.12/hardware_enablement/kmm-kernel-module-management.html#kmm-install-using-web-console_kernel-module-management-operator) to install the operator

2. Label the node(s) you want to use for the initial canary deployment. Select nodes that contain the following NFD label `intel.feature.node.kubernetes.io/gpu=true` which indicates the node contains an Intel data center GPU card

Note: If you would like to use another label name, replace `intel.feature.node.kubernetes.io/dgpu-canary=true` in the below command with the label of your choice. 

Note: If you create a label on a node, the label is ephemeral and will not be present on the node if it is rebooted. Another option is to use default node labels which are persistent after reboot. Make sure to replace `intel.feature.node.kubernetes.io/dgpu-canary: 'true'` with the label of your choice in the [intel-dgpu.yaml](/kmmo/intel-dgpu.yaml) file.

```
$ oc label node <node_name> intel.feature.node.kubernetes.io/dgpu-canary=true
```

3. Use Pre-build mode to deploy the driver container

```
$ oc apply -f https://github.com/intel/intel-technology-enabling-for-openshift/blob/main/kmmo/intel-dgpu.yaml
```

To verify the drivers have been loaded and the GPU card has been successfully provisioned, run `chroot /host` and then `lsmod | grep i915` on the node terminal via OCP web console (Compute -> Nodes -> Select a node that has the GPU card -> Terminal) or CLI. Ensure `i915` and `intel_vsec` is present in the command output.

Optionally, run `dmesg` to ensure there are no errors in the kernel message log during driver loading and firmware initialization.

For Pre-build mode, KMM will add `kmm.node.kubernetes.io/intel-dgpu.ready` label to the node which indicates the drivers and firmware have been loaded successfully.

For On-premise build mode, the label name that KMM will add is `kmm.node.kubernetes.io/intel-dgpu-on-premise-build-mode.ready`.

4. If the initial canary deployment succeeds, deploy the driver container on all the eligible nodes in the cluster. Comment the line `intel.feature.node.kubernetes.io/dgpu-canary: 'true'` in the [intel-dgpu.yaml](/kmmo/intel-dgpu.yaml) file and run the above command in step 3

Note: Skip this step if you skipped canary deployment.

### Optional: Using On-premise Build Mode

1. Prior to using this mode, review the `Canary Deployment on OpenShift` section above. Please complete step 1 and 2 above. Next, run the following commands to create a `ConfigMap` from the specified dockerfile to build the driver container image

```
$ git clone https://github.com/intel/intel-data-center-gpu-driver-for-openshift.git && cd intel-data-center-GPU-driver-for-openshift/docker
$ oc create -n openshift-kmm configmap intel-dgpu-dockerfile-configmap --from-file=dockerfile=intel-dgpu-driver.Dockerfile
```

2. Deploy On-premise Build Mode

```
$ oc apply -f https://github.com/intel/intel-technology-enabling-for-openshift/blob/main/kmmo/intel-dgpu-on-premise-build-mode.yaml
```

Review step 3 above for steps to verify the drivers have been loaded and the GPU card has been successfully provisioned.

3. If the initial canary deployment succeeds, deploy the driver container on all the eligible nodes in the cluster. Comment the line `intel.feature.node.kubernetes.io/dgpu-canary: 'true'` in the [intel-dgpu-on-premise-build-mode.yaml](/kmmo/intel-dgpu-on-premise-build-mode.yaml) file and run the above command in step 2

Note: Skip this step if you skipped canary deployment.
