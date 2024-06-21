# Intel® Technology Enabling for OpenShift*
## Overview
The Intel Technology Enabling for OpenShift project provides Intel Data Center hardware feature-provisioning technologies with the [Red Hat OpenShift Container Platform (RHOCP)](https://www.redhat.com/en/technologies/cloud-computing/openshift/container-platform). The technology to deploy and manage Intel Enterprise AI End-to-End (E2E) solutions and the related reference workloads for these features are also included in the project.  

These Intel Data Center hardware features currently include: 
- **Intel® Software Guard Extensions (Intel® SGX)**
- **Intel® Data Center GPU Flex Series**
- **Intel® Data Center GPU Max Series** 
- **Intel® QuickAssist Technology (Intel® QAT)**

The following features will be included in future releases:  
- Intel® Data Streaming Accelerator (Intel® DSA) 

Intel AI hardware and optimized software solutions are integrated into Red Hat OpenShift AI for ease of provisioning and configuration. The [Habana AI Operator](https://catalog.redhat.com/software/container-stacks/detail/64342b3bcbfbb9a6588ce8dd?gs&q=habana) is used to provision Intel® Gaudi® accelerators and released on the Red Hat Ecosystem Catalog.

Red Hat Distributed CI* (DCI) based CI/CD pipeline is leveraged to enable and test this E2E solution with each RHOCP release to ensure new features and improvements can be promptly available.

The [Open Platform for Enterprise AI (OPEA)](https://github.com/opea-project) RAG workloads are used to validate and optimize Intel enterprise AI E2E solutions.

![Alt text](/docs/images/Intel-Technology-Enabling-for-OpenShift-Architecture.png)

<div align="center">
  Figure-1 Intel Technology Enabling for OpenShift Architecture 
</div>

## Releases and Supported Platforms 
- [Supported Intel hardware features](/docs/supported_platforms.md#supported-intel-hardware-features) and [supported RHOCP versions](/docs/supported_platforms.md#supported-rhocp-versions)  
- [Release details](/docs/releases.rst)

## Getting started
See reference [BIOS Configuration](/docs/supported_platforms.md#bios-configuration) required for each feature.

### Provisioning RHOCP cluster   
Use one of these two options to provision an RHOCP cluster: 
- Use the methods introduced in [RHOCP documentation](https://docs.openshift.com/container-platform/4.14/installing/index.html). 
- Use [Distributed CI](https://doc.distributed-ci.io/) as we do in this project.  

In this project, we provisioned RHOCP 4.14 on a bare-metal multi-node cluster. For details about the supported RHOCP infrastructure, see the [Supported Platforms](/docs/supported_platforms.md) page.

### Provisioning Intel hardware features on RHOCP
Please follow the steps below to provision the hardware features 
1. Setting up [Node Feature Discovery](/nfd/README.md) 
2. Setting up [Machine Configuration](/machine_configuration/README.md) 
3. Setting up [Out of Tree Drivers](/kmmo/README.md) 
4. Setting up [Device Plugins](/device_plugins/README.md) 

### Verifying hardware feature provisioning 
You can use the instructions in the [link](/tests/l2/README.md) to verify the hardware features provisioning. 

## Upgrade (To be added) 

## Reference end-to-end solution 
The reference end-to-end solution is based on Intel hardware feature provisioning provided by this project. 

[Intel AI Inferencing Solution](/e2e/inference/README.md) with [OpenVINO](https://github.com/openvinotoolkit/openvino) and [RHOAI](https://www.redhat.com/en/technologies/cloud-computing/openshift/openshift-data-science) 

## Reference workloads 
Here are some reference workloads built on the end-to-end solution and Intel hardware feature provisioning in this project. 
- Large Language Model (To be added) 
- Open Federated Learning (To be added) 

## Advanced Guide 
This section discusses architecture and other technical details that go beyond getting started. 
- [Architecture and Working Scope](https://github.com/intel/intel-technology-enabling-for-openshift/wiki/Intel-Technology-Enabling-for-OpenShift-Architecture-and-Working-Scope) 

## Release Notes
Check the [link](https://github.com/intel/intel-technology-enabling-for-openshift/releases/) for the Release Notes.  

## Support
If users encounter any issues or have questions regarding Intel Technology Enabling for OpenShift, we recommend them to seek support through the following channels:
### Commercial support from Red Hat 
This project relies on features developed and released with the latest RHOCP release. Commercial RHOCP release support is outlined in the [Red Hat OpenShift Container Platform Life Cycle Policy](https://access.redhat.com/support/policy/updates/openshift) and Intel collaborates with Red Hat to address specific requirements from our users.  

#### Open-Source Community Support
Intel Technology Enabling for OpenShift is run as an open-source project on GitHub. Project GitHub [issues](https://github.com/intel/intel-technology-enabling-for-openshift/issues) can be used as the primary support interface for users to submit feature requests and report issues to the community when using Intel technology provided by this project. Please provide detailed information about your issue and steps to reproduce it, if possible.

## Contribute
See [CONTRIBUTING](CONTRIBUTING.md) for more information.

## Security
To report a potential security vulnerability, please refer to [security.md](/security.md) file. 

## License
Distributed under the open source license. See [LICENSE](/LICENSE.txt) for more information.

## Code of Conduct
Intel has adopted the Contributor Covenant as the Code of Conduct for all of its open source projects. See [CODE_OF_CONDUCT](/CODE_OF_CONDUCT.md) file.
