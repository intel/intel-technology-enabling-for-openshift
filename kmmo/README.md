# KMMO
The [Kernel Module Management Operator](https://github.com/rh-ecosystem-edge/kernel-module-management) manages the deployment and lifecycle of out-of-tree kernel modules.

For iOCP Project, KMMO is used on Day 2 to manage Intel OOT (Out-Of-Tree) dGPU drivers and firmware via Module CRD API.

Note: KMMO project is still under development. Welcome to report issues in [iOCP](https://github.com/intel-sandbox/intel.dgpu.operator.prototype/issues) project or [KMMO project](https://github.com/rh-ecosystem-edge/kernel-module-management).

## Canary Deployment
Canary deployment is the concept of deploying something to a subset of nodes (instead of all nodes) to ensure it does not break the entire cluster.

Canary deployment is used by default to prevent KMMO from deploying Intel dGPU drivers and firmware to all nodes to ensure that the entire cluster is not affected due to a potential issue from driver insmod into kernel or loaded firmware.

As a result, prior to deplying the KMM module CRD, the cluster administrator needs to add label `intel.feature.node.kubernetes.io/dgpu-canary=true` to one or more Intel dGPU nodes for initial deployment.

Note: An Intel dGPU node is a node that contains label `intel.feature.node.kubernetes.io/gpu=true`. It is labeled automatically by NFD. See [NFD README.](https://github.com/intel-sandbox/intel.dgpu.operator.prototype/tree/main/nfd/README.md)

Below is a command for reference to add a label to a node:

```$ oc label node dGPU_node_name intel.feature.node.kubernetes.io/dgpu-canary=true```

## Install KMMO

Follow the [KMMO operator installation steps](https://docs.openshift.com/container-platform/4.12/hardware_enablement/kmm-kernel-module-management.html#kmm-install-using-web-console_kernel-module-management-operator) to install the operator.

## 3 Modes for iOCP Project:

User can select from one of the modes listed below.

### 1. Pre-build Mode 

This is the default and recommended mode. KMMO will use pre-built driver container images from quay registry to deploy Intel dGPU drivers and firmware.

To use this mode, run the following command:

```$ oc apply -f https://github.com/intel-sandbox/intel.dgpu.operator.prototype/blob/main/kmmo/intel-dgpu.yaml```

### 2. On-premise Build Mode

KMMO will build driver container image directly on cluster using in-cluster registry. KMMO uses OCP `Build` and `ImageStream` to support this mode.

Prior to using this mode, run the following commands to create a `ConfigMap` which contains the dockerfile to build the driver container image:

```$ git clone https://github.com/intel-sandbox/Intel-data-center-GPU-driver-for-openshift.git && cd Intel-data-center-GPU-driver-for-openshift```

```$oc create -n openshift-kmm configmap intel-dgpu-dockerfile-configmap --from-file=intel-dgpu-driver.Dockerfile```

To use this mode, run the following command:

```$ oc apply -f https://github.com/intel-sandbox/intel.dgpu.operator.prototype/blob/main/kmmo/intel-dgpu-on-premise-build-mode.yaml```

### 3. CI/CD mode

This mode is used to automatically build driver container image, test, and push image to quay registry to support Pre-build mode.

This mode is still under development.


