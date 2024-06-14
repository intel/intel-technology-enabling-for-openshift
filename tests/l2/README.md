# Verifying Intel Hardware Feature Provisioning
## Introduction
After provisioning Intel hardware features on RHOCP, the respective hardware resources are exposed to the RHOCP cluster. The workload containers can request these resources. The following sample workloads help verify if these resources can be used as expected. These sample workloads container images are built and packaged on-premises through [RHOCP BuildConfig](https://docs.openshift.com/container-platform/4.14/cicd/builds/understanding-buildconfigs.html) and pushed to the embedded repository through [RHOCP ImageStream](https://docs.openshift.com/container-platform/4.14/openshift_images/image-streams-manage.html).

## Prerequisites
•	Provisioned RHOCP cluster. Follow steps [here](/README.md#provisioning-rhocp-cluster). 

•	Provisioning Intel HW features on RHOCP. Follow steps [here](/README.md#provisioning-intel-hardware-features-on-rhocp)

### Verify Intel® Data Center GPU provisioning
Please refer to Intel DGPU provisioning validation tests [here](dgpu/README.md)

### Verify Intel® Software Guard Extensions (Intel® SGX) Provisioning
Please refer to Intel SGX provisioning validation tests [here](sgx/README.md)

### Verify Intel® QuickAssist Technology provisioning
Please refer to Intel QAT provisioning validation tests [here](qat/README.md)
