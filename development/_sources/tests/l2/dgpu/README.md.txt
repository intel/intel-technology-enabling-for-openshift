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

### Verify Media feature provisioning for Intel® Data Center GPU
It is verified on Intel® Data Center GPU Flex Series.

#### VAInfo
This workload runs [vainfo](https://github.com/intel/libva-utils) utilizing the i915 resource from GPU provisioning and displays the Media features supported by the GPU.

*	Build the workload container image.

```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/tests/l2/dgpu/vainfo_build.yaml
```

*	Deploy and execute the workload.

```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/tests/l2/dgpu/vainfo_job.yaml
```

* Check the results.
```
  $ oc get pods -n intel-dgpu
  NAME                        READY   STATUS      RESTARTS   AGE
  intel-dgpu-vainfo-1-build   0/1     Completed   0          9m49s
  intel-dgpu-vainfo-rbdvz     0/1     Completed   0          45s
```
```
$ oc logs intel-dgpu-vainfo-rbdvz -n intel-dgpu
vainfo: VA-API version: 1.20 (libva 2.20.0)
vainfo: Driver version: Intel iHD driver for Intel(R) Gen Graphics - 23.4.3 (a9f272496)
vainfo: Supported profile and entrypoints
      VAProfileNone                   : VAEntrypointVideoProc
      VAProfileNone                   : VAEntrypointStats
      VAProfileMPEG2Simple            : VAEntrypointVLD
      VAProfileMPEG2Main              : VAEntrypointVLD
      VAProfileH264Main               : VAEntrypointVLD
      VAProfileH264Main               : VAEntrypointEncSliceLP
      VAProfileH264High               : VAEntrypointVLD
      VAProfileH264High               : VAEntrypointEncSliceLP
      VAProfileJPEGBaseline           : VAEntrypointVLD
      VAProfileJPEGBaseline           : VAEntrypointEncPicture
      VAProfileH264ConstrainedBaseline: VAEntrypointVLD
      VAProfileH264ConstrainedBaseline: VAEntrypointEncSliceLP
      VAProfileHEVCMain               : VAEntrypointVLD
      VAProfileHEVCMain               : VAEntrypointEncSliceLP
      VAProfileHEVCMain10             : VAEntrypointVLD
      VAProfileHEVCMain10             : VAEntrypointEncSliceLP
      VAProfileVP9Profile0            : VAEntrypointVLD
      VAProfileVP9Profile1            : VAEntrypointVLD
      VAProfileVP9Profile2            : VAEntrypointVLD
      VAProfileVP9Profile3            : VAEntrypointVLD
      VAProfileHEVCMain12             : VAEntrypointVLD
      VAProfileHEVCMain422_10         : VAEntrypointVLD
      VAProfileHEVCMain422_10         : VAEntrypointEncSliceLP
      VAProfileHEVCMain422_12         : VAEntrypointVLD
      VAProfileHEVCMain444            : VAEntrypointVLD
      VAProfileHEVCMain444            : VAEntrypointEncSliceLP
      VAProfileHEVCMain444_10         : VAEntrypointVLD
      VAProfileHEVCMain444_10         : VAEntrypointEncSliceLP
      VAProfileHEVCMain444_12         : VAEntrypointVLD
      VAProfileHEVCSccMain            : VAEntrypointVLD
      VAProfileHEVCSccMain            : VAEntrypointEncSliceLP
      VAProfileHEVCSccMain10          : VAEntrypointVLD
      VAProfileHEVCSccMain10          : VAEntrypointEncSliceLP
      VAProfileHEVCSccMain444         : VAEntrypointVLD
      VAProfileHEVCSccMain444         : VAEntrypointEncSliceLP
      VAProfileAV1Profile0            : VAEntrypointVLD
      VAProfileAV1Profile0            : VAEntrypointEncSliceLP
      VAProfileHEVCSccMain444_10      : VAEntrypointVLD
      VAProfileHEVCSccMain444_10      : VAEntrypointEncSliceLP
```
#### Using libvpl
This workload runs various test programs from [libvpl](https://github.com/intel/libvpl) utilizing the i915 resource from GPU provisioning and displays the Video Processing Library (VPL) fetures supported by the GPU.

*	Build the workload container image.

```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/tests/l2/dgpu/intelvpl_build.yaml
```

*	Deploy and execute the workload.

```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/tests/l2/dgpu/intelvpl_job.yaml
```

* Check the results.
```
$ oc get pods -n intel-dgpu
  NAME                          READY   STATUS      RESTARTS   AGE
  intel-dgpu-intelvpl-1-build   0/1     Completed   0          26m
  intel-dgpu-intelvpl-f68tb     0/1     Completed   0          15m
```
```
$ oc logs intel-dgpu-intelvpl-f68tb -n intel-dgpu

Implementation #0: mfxhw64
  Library path: /usr/lib64/libmfxhw64.so.1.35
  AccelerationMode: MFX_ACCEL_MODE_VIA_VAAPI
  ApiVersion: 1.35
  Impl: MFX_IMPL_TYPE_HARDWARE
  VendorImplID: 0x0000
  ImplName: mfxhw64
  License:
  Version: 1.2
  Keywords: MSDK,x64
  VendorID: 0x8086
  mfxAccelerationModeDescription:
    Version: 1.0
    Mode: MFX_ACCEL_MODE_VIA_VAAPI
  mfxPoolPolicyDescription:
    Version: 1.0
  mfxDeviceDescription:
    MediaAdapterType: MFX_MEDIA_UNKNOWN
    DeviceID: 56c1/0
    Version: 1.1
  mfxDecoderDescription:
    Version: 0.0
  mfxEncoderDescription:
    Version: 0.0
  mfxVPPDescription:
    Version: 0.0
  NumExtParam: 0
  Warning - MFX_IMPLCAPS_SURFACE_TYPES not supported

Implementation #1: mfxhw64
  Library path: /usr/lib64/libmfxhw64.so.1.35
  AccelerationMode: MFX_ACCEL_MODE_VIA_VAAPI
  ApiVersion: 1.35
  Impl: MFX_IMPL_TYPE_HARDWARE
  VendorImplID: 0x0001
  ImplName: mfxhw64
  License:
  Version: 1.2
  Keywords: MSDK,x64
  VendorID: 0x8086
  mfxAccelerationModeDescription:
    Version: 1.0
    Mode: MFX_ACCEL_MODE_VIA_VAAPI
  mfxPoolPolicyDescription:
    Version: 1.0
  mfxDeviceDescription:
    MediaAdapterType: MFX_MEDIA_UNKNOWN
    DeviceID: 56c1/1
    Version: 1.1
  mfxDecoderDescription:
    Version: 0.0
  mfxEncoderDescription:
    Version: 0.0
  mfxVPPDescription:
    Version: 0.0
  NumExtParam: 0
  Warning - MFX_IMPLCAPS_SURFACE_TYPES not supported

Total number of implementations found = 2
```

## See Also
For GPU demos on vanilla Kubernetes, refer to [link](https://github.com/intel/intel-device-plugins-for-kubernetes/tree/main/demo/intel-opencl-icd) 