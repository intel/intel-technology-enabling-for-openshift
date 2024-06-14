### Verify Intel® Data Center GPU provisioning
It is supported for Intel® Data Center GPU Flex and Max Series. 
#### clinfo
This workload runs [clinfo](https://github.com/Oblomov/clinfo) utilizing the i915 resource from GPU provisioning and displays the related GPU information.

*	Build the workload container image. 

```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/tests/l2/dgpu/clinfo_build.yaml 
```

*	Deploy and execute the workload.

```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/tests/l2/dgpu/clinfo_job.yaml
```

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

```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/tests/l2/dgpu/hwinfo_build.yaml
```

*	Deploy and execute the workload.

```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/tests/l2/dgpu/hwinfo_job.yaml
```

* Check the results
``` 
  $ oc get pods
  intel-dgpu-hwinfo-1-build   0/1     Completed   0          2m23s
  intel-dgpu-hwinfo-44k4d     0/1     Completed   0          106s
```
```
$ oc logs intel-dgpu-hwinfo-44k4d  
 282: PCI aa00.0: 0380 Display controller
  [Created at pci.386]
  Unique ID: YxOB.+ER_Ec9Ujm4
  Parent ID: xBFW.xbjkZcxCQYD
  SysFS ID: /devices/pci0000:a7/0000:a7:01.0/0000:a8:00.0/0000:a9:01.0/0000:aa:00.0
  SysFS BusID: 0000:aa:00.0
  Hardware Class: graphics card
  Model: "Intel Display controller"
  Vendor: pci 0x8086 "Intel Corporation"
  Device: pci 0x0bda
  SubVendor: pci 0x8086 "Intel Corporation"
  SubDevice: pci 0x0000
  Revision: 0x2f
  Driver: "i915"
  Driver Modules: "i915"
  Memory Range: 0x44fe3f000000-0x44fe3fffffff (ro,non-prefetchable)
  Memory Range: 0x447000000000-0x447fffffffff (ro,non-prefetchable)
  IRQ: 511 (140 events)
  Module Alias: "pci:v00008086d00000BDAsv00008086sd00000000bc03sc80i00"
  Config Status: cfg=new, avail=yes, need=no, active=unknown
  Attached to: #89 (PCI bridge)
```                        

## See Also
For GPU demos on vanilla Kubernetes, refer to [link](https://github.com/intel/intel-device-plugins-for-kubernetes/tree/main/demo/intel-opencl-icd) 