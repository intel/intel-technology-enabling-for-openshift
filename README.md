# Disclaimer 
This project is currently under active development, and as such, all source code may not be included in any release. This means that the code is subject to change without notice, and that any information contained within the code should be considered as work in progress. 

# Intel Technology Enabling for OpenShift
## General
Intel Technology Enabling for OpenShift project provides Intel data center and edge hardware feature provisioning, to enable related end to end (E2E) solutions and the reference workloads for these Intel features on Red Hat OpenShift platform.   

The goal of the project is to develop and release open, easy-to-use, integrated, and more secure solutions to provision and manage Intel technologies on [Red Hat OpenShift Container Platform (OCP)](https://www.redhat.com/en/technologies/cloud-computing/openshift/container-platform). These Intel technologies mainly include Intel data center and edge hardware features and related software stacks for E2E solutions and reference workloads.

To achieve this goal, following OCP software stack development and management life cycle should be followed:
- The related software stacks must be designed and implemented for Cloud Native and OCP environment from Day 0
-	The solutions and instructions must allow users to configure and provision their OCP cluster with technologies from intel and other vendors on Day 1
-	The solutions and instructions to provision intel hardware features and manage the life cycle of containerized software stack must be present in Day 2

## Intel Hardware features Provisioning for OpenShift

To provision Intel Hardware features on OCP, following open-source projects are used: 
* **[Node Feature Discovery (NFD)](https://github.com/kubernetes-sigs/node-feature-discovery), [NFD Operator](https://github.com/openshift/cluster-nfd-operator)** are used to automatically label the nodes for Hardware provisioning operation.
* **[Machine Config Operator (MCO)](https://github.com/openshift/machine-config-operator)** is used to configure the Red Hat Enterprise Linux Core OS (RHCOS) on the nodes.
* **[Kernel Module Management (KMM)](https://github.com/kubernetes-sigs/kernel-module-management), [KMM Operator](https://github.com/rh-ecosystem-edge/kernel-module-management)** are used to manage deployment and lifecycle of Intel Data Center GPU Driver.
* **[Intel Data Center GPU Driver For OpenShift](https://github.com/intel/intel-data-center-gpu-driver-for-openshift)** use **[Intel GPU Drivers](https://github.com/intel-gpu)** build, package certify and release Intel data center GPU driver container images for OCP.
* **[Intel Device Plugins for Kubernetes project](https://github.com/intel/intel-device-plugins-for-kubernetes)** Provides Intel GPU/SGX/QAT device plugins images and the operator to deploy and manage the life cycle of these device plugins 

Intel Hardware features include:

* **Intel® Software Guard Extensions (Intel® SGX)**
* **Intel® Data Center GPU Flex Series**
*	**Intel® QuickAssist Technology (Intel® QAT)**

Below features are under consideration to be included in the future releases 

*	Intel® Data Center GPU Max Series
*	Intel® Data Streaming Accelerator (Intel® DSA)
*	Intel® In-Memory Analytics Accelerator (Intel® IAA)
*	Intel® Dynamic Load Balancer (Intel® DLB)

## [Intel AI Inference E2E Solution for OpenShift](e2e/inference/README.md)

## Hardware Requirements

### Intel SGX supported platform

* Third Generation Intel® Xeon® Scalable Processors (or later version) are used by the cluster.
* Contact your server or BIOS vendor for the BIOS setting to enable the feature.

### Intel Data Center GPU card supported platform

* The Intel® Data Center GPU Flex Series 140 or Intel® Data Center GPU Flex Series 170 Card is enabled on the nodes.
* Contact your server or BIOS vendor for the BIOS setting to enable the cards.

### Intel QAT supported platform
* 4th Gen Intel® Xeon® Scalable Processors (or later versions) are used by the cluster.
* Contact your server or BIOS vendor for the BIOS setting to enable the feature.

## Get Started 
To properly provision Intel hardware features, deploy and manage the related e2e solutions as well as the reference workload, below OCP software stack development and life cycle management flow is followed by this project

### Day 0 - Define the requirements of the OCP platform and design it.
Red Hat [OpenShift Operator](https://www.redhat.com/en/technologies/cloud-computing/openshift/what-are-openshift-operators) automates the creation, configuration, and management of instances of Kubernetes-native applications. It is based on [Kubernetes operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/). So the software stack running on OCP needs to be containerized and managed by a specific Operator. As an essential part of OCP, operators need to be well-designed from Day 0. Good examples are [Intel Device Plugins Operator](https://github.com/intel/intel-device-plugins-for-kubernetes) and [KMM Operator](https://github.com/rh-ecosystem-edge/kernel-module-management).

### Day 1 - provision the OCP platform and configure it to a working state.
This project mainly focuses on bare metal OCP cluster.  [Distributed CI (DCI)](https://doc.distributed-ci.io/dci-openshift-agent/) is used to provision the bare metal OCP cluster. Users can also refer to [bare metal OCP cluster installation instructions](https://docs.openshift.com/container-platform/4.11/installing/installing_bare_metal_ipi/ipi-install-overview.html) to install the bare metal OCP cluster.

To avoid rebooting the nodes and some other issues on Day 2, Some Machine Configurations operations can be enforced on day 1 when provisioning the cluster. The related discussion is ongoing.

### Day 2.0 - Zero O’clock of Day 2
The Day 2.0 concept is introduced for users to provision intel hardware features right after provisioning an OCP cluster and before any user workloads are deployed. Refer to the steps below to provision Intel hardware features:

* **[Deploy Node Feature Descovery on OpenShift](nfd/README.md#steps-to-install-and-configure-nfd-operator-on-ocp-cluster)**
* **[Setup Machine Configuration on OpenShift](machine_configuration/README.md#general-configuration-for-provisioning-intel-hardware-features)**  

`Note: Running the above steps on Day 2.0 is recommended. However, if you want to provision the features above with the existing cluster on day 2, please be advised that some machine configuration operations might trigger the pods to drain and reboot the nodes. Some of the ongoing efforts in the MCO upstream are to set the machine configurations without rebooting.`

### Day 2 - OCP platform is installed and ready to begin providing services.
Multiple operators are used to provision Intel hardware features and deploy, manage the e2e solutions, and reference workloads.

**Provisioning Intel Hardware features on OpenShift**
* **[Deploy Intel Data Center GPU Driver on OpenShift](kmmo/README.md#managing-intel-dgpu-driver-with-kmm-operator)**
* **[Deploy Intel Device Plugins on OpenShift](device_plugins/README.md#deploy-intel-device-plugins-on-openshift)**
 
**Deploy E2E Solution**
* **[Deploy Intel AI Reference E2E Solution](e2e/inference/README.md#deploy-intel-ai-inference-e2e-solution)**
