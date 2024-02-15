# Verifying Intel Hardware Feature Provisioning
## Introduction
After provisioning Intel hardware features on RHOCP, the respective hardware resources are exposed to the RHOCP cluster. The workload containers can request these resources. The following sample workloads help verify if these resources can be used as expected. These sample workloads container images are built and packaged on-premises through [RHOCP BuildConfig](https://docs.openshift.com/container-platform/4.14/cicd/builds/understanding-buildconfigs.html) and pushed to the embedded repository through [RHOCP ImageStream](https://docs.openshift.com/container-platform/4.14/openshift_images/image-streams-manage.html).

## Prerequisites
•	Provisioned RHOCP cluster. Follow steps [here](/README.md#provisioning-rhocp-cluster). 

•	Provisioning Intel HW features on RHOCP. Follow steps [here](/README.md#provisioning-intel-hardware-features-on-rhocp)

### Verify Intel® Software Guard Extensions (Intel® SGX) Provisioning
This [SampleEnclave](https://github.com/intel/linux-sgx/tree/master/SampleCode/SampleEnclave) application workload from the Intel SGX SDK runs an Intel SGX enclave utilizing the EPC resource from the Intel SGX provisioning.
* Build the container image. 
  
```$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/tests/l2/sgx/sgx_build.yaml```

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
#### clinfo
This workload runs [clinfo](https://github.com/Oblomov/clinfo) utilizing the i915 resource from GPU provisioning and displays the related GPU information.
* To work around [issue](https://github.com/intel/intel-technology-enabling-for-openshift/issues/107), please run the below command on the node with the GPU where your workload will run:
  
```$ setsebool container_use_devices on```

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
#### hwinfo

This workload runs ```hwinfo``` utilizing the i915 resource from GPU provisioning and displays the related GPU information. Refer to [link](https://dgpu-docs.intel.com/driver/installation.html#verify-install)


*	Build the workload container image. 

```$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/tests/l2/dgpu/hwinfo_build.yaml ```

*	Deploy and execute the workload.

```$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/tests/l2/dgpu/hwinfo_job.yaml```

* Check the results
``` 
  $ oc get pods
  intel-dgpu-hwinfo-1-build   0/1     Completed   0          111m
  intel-dgpu-hwinfo-l4572     0/1     Completed   0          110m
```
```
$ oc logs intel-dgpu-hwinfo-l4572  
  88: PCI 3700.0: 0300 VGA compatible controller (VGA)
  [Created at pci.386]
  Unique ID: ug9x.K5dXxUXtHX2
  Parent ID: QGIh.mr2N3fBJq5F
  SysFS ID: /devices/pci0000:30/0000:30:02.0/0000:31:00.0/0000:32:00.0/0000:33:00.0/0000:34:08.0/0000:35:00.0/0000:36:01.0/0000:37:00.0
  SysFS BusID: 0000:37:00.0
  Hardware Class: graphics card
  Model: "Intel VGA compatible controller"
  Vendor: pci 0x8086 "Intel Corporation"
  Device: pci 0x56c1
  SubVendor: pci 0x8086 "Intel Corporation"
  SubDevice: pci 0x4905
  Revision: 0x04
  Driver: "i915"
  Driver Modules: "i915"
  Memory Range: 0xd0000000-0xd0ffffff (rw,non-prefetchable)
  Memory Range: 0x3afc00000000-0x3afdffffffff (ro,non-prefetchable)
  IRQ: 928 (86 events)
  Module Alias: "pci:v00008086d000056C1sv00008086sd00004905bc03sc00i00"
  Config Status: cfg=new, avail=yes, need=no, active=unknown
  Attached to: #29 (PCI bridge)
```                        

### Verify Intel® QuickAssist Technology provisioning
This workload runs [qatlib](https://github.com/intel/qatlib) sample tests using RedHat built and distributed Qatlib RPM packages from the codeready-builder-for-rhel-9-x86_64-rpms repo. Refer to the [qatlib readme](https://github.com/intel/qatlib/blob/main/INSTALL) for more details. 

*	Build the workload container image

Please replace the credentials in buildconfig yaml with your RedHat account login credentials. 

```$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/tests/l2/qat/qatlib_build.yaml ```

* Create SCC intel-qat-scc for Intel QAT based workload, if this SCC is not created   
  
```$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/security/qatlib_scc.yaml```
      
* Create the intel-qat service account to use intel-qat-scc
  
```$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/security/qatlib_rbac.yaml```

* Deploy the qatlib workload job with intel-qat service account
  
```$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/tests/l2/qat/qatlib_job.yaml```

* Check the results.
``` 
  $ oc get pods
  intel-qat-workload-c6g9v   0/1     Completed   0          4m13s
```


* For all sample tests `cpa_sample_code` 

```
$ oc logs intel-qat-workload-c6g9v
qaeMemInit started
icp_sal_userStartMultiProcess("SSL") started
There are no crypto instances
*** QA version information ***
device ID               = 0
software                = 23.2.0
*** END QA version information ***
Inst 0, Affin: 0, Dev: 0, Accel 0, EE 0, BDF ED:00:01
Inst 1, Affin: 1, Dev: 0, Accel 0, EE 0, BDF ED:00:01
Inst 2, Affin: 2, Dev: 0, Accel 0, EE 0, BDF ED:00:01
Inst 3, Affin: 3, Dev: 0, Accel 0, EE 0, BDF ED:00:01
---------------------------------------
API                    Traditional
Session State          STATELESS
Algorithm              DEFLATE
Huffman Type           STATIC
Mode                   ASYNCHRONOUS
CNV Enabled            YES
Direction              COMPRESS
Packet Size            8192
Compression Level      1
Corpus                 CALGARY_CORPUS
Corpus Filename        calgary
CNV Recovery Enabled   YES
Number of threads      4
Total Responses        158400
Total Retries          2242671
Clock Cycles Start     126150916653843
Clock Cycles End       126151409143747
Total Cycles           492489904
CPU Frequency(kHz)     1700160
Throughput(Mbps)       35920
Compression Ratio      0.4897
---------------------------------------

Inst 0, Affin: 0, Dev: 0, Accel 0, EE 0, BDF ED:00:01
Inst 1, Affin: 1, Dev: 0, Accel 0, EE 0, BDF ED:00:01
Inst 2, Affin: 2, Dev: 0, Accel 0, EE 0, BDF ED:00:01
Inst 3, Affin: 3, Dev: 0, Accel 0, EE 0, BDF ED:00:01
---------------------------------------
```


## See Also
For Intel SGX demos on vanilla Kubernetes, refer to [link](https://github.com/intel/intel-device-plugins-for-kubernetes/tree/main/demo/sgx-sdk-demo) 

For GPU demos on vanilla Kubernetes, refer to [link](https://github.com/intel/intel-device-plugins-for-kubernetes/tree/main/demo/intel-opencl-icd) 