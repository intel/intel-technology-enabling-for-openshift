## Test plan and architecture overview
To ensure the whole stack is working as expected and track bugs, a layer based test architecture is needed on OCP. This plan consists of four layers. The first and second layers mentioned below, would be a part of the future automation testing framework on each OCP (x,y,z) release.

#### L1 First Layer: Host OS and Driver Interfaces 
This layer's test cases will be executed after the cluster is setup, NFD operator is installed and right before the Intel Device Plugins Operator installation. 
* Test cases for in-tree and OOT tree drivers
* Any test cases for SELinux and host OS security
* go-based tests cases for devices on all nodes
* Dependency- rely on driver teams for any other testing result 
* Future plan to include OOT driver test cases with KMMO CI/CD pipeline- sanity testing.
* Currently execute as a job. Plan to find an efficient mechanism to deploy and automate on all nodes.

#### L2 Second Layer: Device plugin resource provisioning  
This layer will be executed after the Intel Device Plugins Operator installation and Custom Resource creation. It includes 
* Pod's resource allocation and scheduling 
* Simple workloads 
* Boundary testing for the resources
* Integration and stress testing 
* Future plan to include any OOT driver test cases in KMMO CI/CD pipeline- canary testing 
* Any failure analysis needed during automation, in the future. 

#### L3 Third Layer: E2E solution
This layer consists of the reference workload's/fourth layer's prerequisites and dependencies- after the cluster, operators and resources are accurately deployed. It is custom for each Intel Hero feature's final use case. For example, dGPU smart city workload would have RHODS (Red Hat OpenShift Data Science) and OpenVINO operator as prerequisites included here.

#### L4 Fourth Layer: Reference Workloads
This layer includes the reference final application/usecase for the user. It integrates the whole stack and is custom for each Intel Hero feature.
