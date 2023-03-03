## KMMO
[Kernel Module Management (KMM) Operator](https://github.com/rh-ecosystem-edge/kernel-module-management) manages the deployment and lifecycle of out-of-tree kernel modules with OCP.

In this Project, KMM Operator is used to manage Intel dGPU drivers container images deployment on day 2.

Intel dGPU driver container images are released from [Intel Data Center GPU Driver for OpenShift Project](https://github.com/intel/intel-data-center-gpu-driver-for-openshift/tree/main/release#intel-data-center-gpu-driver-container-images-for-openshift-release)

### KMM Operator Working Mode 

* Pre-build Mode: This is the default and recommended mode. KMMO will use pre-built, certified, and released driver container images from Red Hat Ecosystem Catalog to deploy Intel dGPU drivers.

* On-premise Build Mode: With this mode, Users and build their own driver container image on-premise and then deploy it on the cluster.

* CI/CD mode: This mode is used to automatically build, package the driver container image, test, certificate and release it on the Red Hat ecyosytem Catalog. It is mainly used to construct the [OOT driver CI/CD pipeline](https://github.com/intel/intel-data-center-gpu-driver-for-openshift/tree/main/pipeline#deploy-oot-driver-cicd-pipeline-on-openshift-cluster) 

### Managing Intel dGPU driver with KMM Operator

Below operations are verified on OCP-4.12 bare metal cluster.

* Follow [KMMO operator installation guide](https://docs.openshift.com/container-platform/4.12/hardware_enablement/kmm-kernel-module-management.html#kmm-install-using-web-console_kernel-module-management-operator) to install the operator on OCP.

* Intel dGPU driver Canary Deployment on OpenShift

Canary deployment is used by default to deploy the driver only on the specific node(s). So Before depolying the driver on the nodes cluseter wide, user can get chance to verify the driver on these canary nodes. That can prevent the driver with some potential issues from damaging the cluster. 

label the nodes you want to run the canary deployment

```$ oc label node dGPU_node_name intel.feature.node.kubernetes.io/dgpu-canary=true```

Note: `intel.feature.node.kubernetes.io/gpu=true` is labled by NFD to show that Intel dGPU card is detected on the node. See [Lables Desctioption](/nfd/README.md#labels-description)

* Using pre-build Mode to deploy the driver

```$ oc apply -f https://github.com//intel/intel-technology-enabling-for-openshift/blob/main/kmmo/intel-dgpu.yaml```

* deploy the driver on all the nodes Cluster wide

if the driver is running properly on the canay nodes, you can deploy the driver cluster wide.

comments the line `intel.feature.node.kubernetes.io/dgpu-canary: 'true'` in the intel-dgpu.yamal file and run

```$ oc apply -f https://github.com//intel/intel-technology-enabling-for-openshift/blob/main/kmmo/intel-dgpu.yaml```

### using On-premise Build Mode

Prior to using this mode, run the following commands to create a `ConfigMap` and include the dockerfile to build the driver container image:

```$ git clone https://github.com/intel/intel-data-center-gpu-driver-for-openshift.git && cd intel-data-center-GPU-driver-for-openshift/docker```

```$oc create -n openshift-kmm configmap intel-dgpu-dockerfile-configmap --from-file=dockerfile=intel-dgpu-driver.Dockerfile```

To use this mode, run the following command:

```$ oc apply -f https://github.com/intel/intel-technology-enabling-for-openshift/blob/main/kmmo/intel-dgpu-on-premise-build-mode.yaml```
