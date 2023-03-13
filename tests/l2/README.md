### L2 overview
This layer consists of workloads for resource provisioning after Intel Device Plugins Operator is installed and custom resources are created.

#### Intel® Data Center GPU Flex Series testcase
The testcase used is [clinfo](https://github.com/Oblomov/clinfo), which displays the related information of Intel Data Center GPU Flex Series. The OCP buildconfig is leveraged to build clinfo container image and push it to the embedded repository through OCP imagestream.
The Job pod is scheduled on a node with Intel Data Center GPU Flex Series card and the resource ```gpu.intel.com/i915``` is registered by the GPU device plugin.
Below operations are verified on OCP-4.12 bare-metal cluster.
To build the testcase:
```
oc apply -f clinfo_build.yaml
```
To deploy the job:
```
oc apply -f clinfo_job.yaml
```
To check the clinfo pod logs:
```
oc get pods | grep clinfo
oc logs <clinfo_pod_name>
```

A sample result for clinfo detecting Intel Data Center GPU Flex Series card :
```
  Platform Name                                   Intel(R) OpenCL HD Graphics
Number of devices                                 1
  Device Name                                     Intel(R) Graphics [0x56c1]
  Device Vendor                                   Intel(R) Corporation
  Device Vendor ID                                0x8086
  Device Version                                  OpenCL 3.0 NEO
  Driver Version                                  22.23.23405
  Device OpenCL C Version                         OpenCL C 1.2
  Device Type                                     GPU
  Device Profile                                  FULL_PROFILE
  Device Available                                Yes
  Compiler Available                              Yes
  Linker Available                                Yes
  Max compute units                               128
  Max clock frequency                             2100MHz
  Device Partition                                (core)
    Max number of sub-devices                     0
    Supported partition types                     None
    Supported affinity domains                    (n/a)
  Max work item dimensions                        3
  Max work item sizes                             1024x1024x1024
  Max work group size                             1024
  Preferred work group size multiple              64
  Max sub-groups per work group                   128
  Sub-group sizes (Intel)                         8, 16, 32
  Preferred / native vector sizes
```

#### Intel® SGX test case
The test case used is Intel SGX SDK [Sample Enclave App](https://github.com/intel/linux-sgx/tree/master/SampleCode/SampleEnclave), which launches a simple Intel SGX enclave. OCP buildconfig and imagestream are leveraged for the container image.
The job pod is scheduled on a node enabled with Intel SGX requesting enclave memory resource ```sgx.intel.com/epc```. The resource is created by the SGX device plugin. 
Below operations are verified on OCP 4.12 bare-metal cluster.
To build the test case:
```
oc apply -f sgx_build.yaml
```
To deploy the job:
```
oc apply -f sgx_job.yaml
```
To check the pod logs:
```
oc get pods | grep sgx
oc logs <sgx_pod_name>
```
Sample pod result:
```
Checksum(0x0x7fffac6f41e0, 100) = 0xfffd4143
Info: executing thread synchronization, please wait...
Info: SampleEnclave successfully returned.
Enter a character before exit ...
```
On the node, the updated resources are:
```
oc describe <node name> | grep sgx.intel.com
sgx.intel.com/enclave    1             1
sgx.intel.com/epc        5Mi          5Mi
sgx.intel.com/provision  0            0
```