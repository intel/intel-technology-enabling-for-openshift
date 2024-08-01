# Test Plan

```{admonition} Disclaimer  
Please note that this module is currently under development and may contain partially implemented features, therefore it is not supported in the current release.
```

## Overview
To ensure the whole stack is working as expected and track bugs, a layer based test architecture is needed on OCP. This plan consists of four layers. The first and second layers mentioned below, would be a part of the future automation testing framework on each OCP (x,y,z) release.

### L1 First Layer: Host OS and Driver Interfaces 
Layer 1 test cases should be executed before deploying [Intel Device Plugins Operator](/device_plugins/README.md) and after deploying OOT drivers like [Intel Data Center GPU Driver on OpenShift](/kmmo/README.md). It includes test cases :
* to check existence of in-tree and OOT tree drivers
* for SELinux and host OS security
* check for devices on all nodes
  
### L2 Second Layer: Device Plugin Resource Provisioning  
L2 test cases are executed after deploying the [Intel Device Plugins Operator](/device_plugins/README.md). Refer to [readme](l2/README.md). It includes: 
* Pod's resource allocation and scheduling 
* Simple workloads 
* Boundary testing for the resources
* Future plan for any failure analysis needed during automation. 

### L3 Third Layer: E2E solution
L3 test cases are executed after the specific device plugin related [e2e solution](/e2e) has been deployed. Please refer to [L3 test cases](l3/README.md) for detail.

### L4 Fourth Layer: Reference Workloads
This layer includes the reference final application/usecase for the user. It integrates the whole stack and is custom for each Intel hardware feature and device plugin. This layer is yet to be added in upcoming releases.