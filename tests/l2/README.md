# Verifying Intel Hardware Feature Provisioning
## Introduction
After provisioning Intel hardware features on RHOCP, the respective hardware resources are exposed to the RHOCP cluster. The workload containers can request these resources. The following sample workloads help verify if these resources can be used as expected. These sample workloads container images are built and packaged on-premises through [RHOCP BuildConfig](https://docs.openshift.com/container-platform/4.12/cicd/builds/understanding-buildconfigs.html) and pushed to the embedded repository through [RHOCP ImageStream](https://docs.openshift.com/container-platform/4.12/openshift_images/image-streams-manage.html).
## Prerequisites
•	Provisioned RHOCP 4.12 cluster. Follow steps [here](https://github.com/intel/intel-technology-enabling-for-openshift#provisioning-rhocp-cluster). 
•	Provisioning Intel HW features on RHOCP. Follow steps [here](https://github.com/intel/intel-technology-enabling-for-openshift#provisioning-intel-hardware-features-on-rhocp)
### Verify Intel® Software Guard Extensions (Intel® SGX) Provisioning
This [SampleEnclave](https://github.com/intel/linux-sgx/tree/master/SampleCode/SampleEnclave) application workload from the Intel SGX SDK runs an Intel SGX enclave utilizing the EPC resource from the Intel SGX provisioning.
* Build the container image. 
```$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/ main/tests/l2/sgx/sgx_build.yaml```
* Deploy and run the workload.
```$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/tests/l2/sgx/sgx_job.yaml```
*	Check the results.
```$ oc get pods
  intel-sgx-job-4tnh5          0/1     Completed   0          2m10s
  intel-sgx-workload-1-build   0/1     Completed   0          30s
```
```
$ oc logs intel-sgx-job-4tnh5
  Checksum(0x0x7fffac6f41e0, 100) = 0xfffd4143
  Info: executing thread synchronization, please wait...
  Info: SampleEnclave successfully returned.
  Enter a character before exit ...
```
### Verify Intel® Data Center GPU provisioning
This workload runs [clinfo](https://github.com/Oblomov/clinfo) utilizing the i915 resource from GPU provisioning and displays the related GPU information.
*	Build the workload container image. 
```$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/tests/l2/dgpu/clinfo_build.yaml ```
*	Deploy and execute the workload.
```$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/tests/l2/dgpu/clinfo_job.yaml```
* Check the results.
``` 
  $ oc get pods
  intel-dgpu-clinfo-1-build        0/1     Completed   0          3m20s
  intel-dgpu-clinfo-56mh2          0/1     Completed   0          35s
```
```
$ oc logs intel-dgpu-clinfo-56mh2  
  Platform Name                                   Intel(R) OpenCL HD Graphics
  Number of devices                                 1
  Device Name                                     Intel(R) Data Center GPU Flex Series 140 [0x56c1]
  Device Vendor                                   Intel(R) Corporation
  Device Vendor ID                                0x8086
  Device Version                                  OpenCL 3.0 NEO
  Device UUID                                     86800000-c156-0000-0000-000000000000
  Driver UUID                                     32322e34-332e-3234-3539-352e33350000
  Valid Device LUID                               No
  Device LUID                                     80c6-4e56fd7f0000
  Device Node Mask                                0
  Device Numeric Version                          0xc00000 (3.0.0)
  Driver Version                                  22.43.24595.35
  Device OpenCL C Version                         OpenCL C 1.2
  Device OpenCL C all versions                    OpenCL 
```                                               

## See Also
For Intel SGX demos on vanilla Kubernetes, refer to [link](https://github.com/intel/intel-device-plugins-for-kubernetes/tree/main/demo/sgx-sdk-demo) 
For GPU demos on vanilla Kubernetes, refer to [link](https://github.com/intel/intel-device-plugins-for-kubernetes/tree/main/demo/intel-opencl-icd) 