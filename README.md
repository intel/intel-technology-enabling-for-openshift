# Intel® Technology Enabling for OpenShift*
# Overview
The Intel Technology Enabling for OpenShift project provides Intel Data Center hardware feature-provisioning technologies with the [Red Hat OpenShift Container Platform (RHOCP)](https://www.redhat.com/en/technologies/cloud-computing/openshift/container-platform). The technology to deploy and manage the [End-to-End (E2E)](/e2e) solutions as well as the related reference workloads for these features are also included in the project.  

These Intel Data Center hardware features currently include: 
- Intel® Software Guard Extensions (Intel® SGX) 
- Intel® Data Center GPU Flex Series 

The following features will be included in future releases.  
- Intel® QuickAssist Technology (Intel® QAT) 
- Intel® Data Center GPU Max Series 
- Intel® Data Streaming Accelerator (Intel® DSA) 
- Intel® In-Memory Analytics Accelerator (Intel® IAA) 
- Intel® FPGA N6000 

See details about [Supported Intel Hardware features]().  

Figure-1 is the [Architecture and Working Scope](https://github.com/intel/intel-technology-enabling-for-openshift/wiki/Intel-Technology-Enabling-for-OpenShift-Architecture-and-Working-Scope) of the project  

![Alt text](/docs/images/Intel-Technology-Enabling-for-OpenShift-Architecture.png)

Figure-1 Intel Technology Enabling for OpenShift Architecture 

# Supported platforms 

This [section]() describes the RHOCP infrastructure and Intel hardware features supported by this project. The project lifecycle and support channels can also be found [here](). 

# Getting started  
## Provisioning RHOCP cluster   
Use one of these two options to provision an RHOCP cluster: 
- Use the methods introduced in [RHOCP documentation](https://docs.openshift.com/container-platform/4.12/installing/index.html). 
- Use [Distributed CI](https://doc.distributed-ci.io/) as we do in this project.  

In this project, we provisioned RHOCP 4.12 on a bare-metal multi-node cluster. For details about the supported RHOCP infrastructure, see the [Supported Platforms]() page.

## Provisioning Intel hardware features on RHOCP
Please follow the steps below to provision the hardware features 
1. Setting up [Node Feature Discovery](/nfd/README.md) 
2. Setting up [Machine Configuration](/machine_configuration/README.md) 
3. Setting up [Out of Tree Drivers](/kmmo/README.md) 
4. Setting up [Device Plugins](/device_plugins/README.md) 

## Verifying hardware feature provisioning 
You can use the instructions in the [link]() to verify the hardware features provisioning. 

# Upgrade (To be added) 

# Reference end-to-end solution 
The reference end-to-end solution is based on Intel hardware feature provisioning provided by this project. 

[Intel AI Inferencing Solution](/e2e/inference/README.md)

# Reference workloads 
Here are some reference workloads built on the end-to-end solution and Intel hardware feature provisioning in this project. 
- Large Language Model (To be added) 
- Open Federated Learning (To be added) 

# Advanced Guide 
This section discusses architecture and other technical details that go beyond getting started. 
- [Architecture and Working Scope](https://github.com/intel/intel-technology-enabling-for-openshift/wiki/Intel-Technology-Enabling-for-OpenShift-Architecture-and-Working-Scope) 

# Release Notes
Check the [link](https://github.com/intel/intel-technology-enabling-for-openshift/releases/) for the Release Notes.  

# Contribute
See [CONTRIBUTING](CONTRIBUTING.md) for more information.

# Security
To report a potential security vulnerability, please refer to [security.md](/security.md) file. 

# License
Distributed under the open source license. See [LICENSE](/LICENSE.txt) for more information.

# Code of Conduct
Intel has adopted the Contributor Covenant as the Code of Conduct for all of its open source projects. See [CODE_OF_CONDUCT](/CODE_OF_CONDUCT.md) file.
