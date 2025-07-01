# Intel® Technology Enabling for OpenShift*
## Overview
The project delivers a comprehensive full-stack solution for the Intel® Enterprise AI Foundation on the [OpenShift platform](https://www.redhat.com/en/technologies/cloud-computing/openshift/container-platform), applicable across data center, cloud, and edge environments. It utilizes innovative [General Operators technology](https://github.com/intel/intel-technology-enabling-for-openshift/wiki/Intel-Technology-Enabling-for-OpenShift-Architecture-and-Working-Scope#architecture-options) to provision AI accelerators, including the Intel Gaudi Processor, Flex and Max GPUs, and Xeon CPU accelerators such as QAT, SGX, and DSA. Additionally, the project introduces solutions for integrating [Gaudi Software](https://docs.habana.ai/en/latest/index.html) or [OneAPI-based](https://www.intel.com/content/www/us/en/developer/tools/oneapi/overview.html#gs.kgdasr) AI software into OpenShift AI. Key AI workload integrations, such as LLM inferencing, fine-tuning, and post-training for enterprise AI, are under development. The plans also include the GPU network provisioning and full-stack integration with OpenShift. 

![Alt text](/docs/images/Intel-Technology-Enabling-for-OpenShift-Architecture.png)

<div align="center">
  Figure-1 Intel Enterprise AI foundation for OpenShift Architecture 
</div>

## Releases and Supported Platforms 
Intel Technology Enabling for OpenShift is released in alignment with the OpenShift release cadence. It is recommended to use the latest release. 

For details on supported features and components, refer to the links below: 
- [Supported Intel hardware features](/docs/supported_platforms.md#supported-intel-hardware-features) 
- [Component Matrix](/docs/supported_platforms.md#component-matrix) 

To review the release history, please visit the following link: 
- [Intel Technology Enabling for OpenShift Release details](/docs/releases.rst)

## Getting started
See reference [BIOS Configuration](/docs/supported_platforms.md#bios-configuration) required for each feature.

### Provisioning RHOCP cluster   
Use one of these two options to provision an RHOCP cluster: 
- Use the methods introduced in [RHOCP documentation](https://docs.redhat.com/en/documentation/openshift_container_platform/4.18/html/installation_overview/ocp-installation-overview). 
- Use [Distributed CI](https://doc.distributed-ci.io/) as we do in this project.  

In this project, we provisioned RHOCP 4.18 on a bare-metal multi-node cluster. For details about the supported RHOCP infrastructure, see the [Supported Platforms](/docs/supported_platforms.md) page.

### Provisioning Intel hardware features on RHOCP
If you are familiar with the steps mentioned below to provision the accelerators, you can use [One-Click](/one_click/README.md) solution as a reference to provision the accelerator automatically.

Follow [Setting up HabanaAI Operator](/gaudi/README.md) to provision Intel Gaudi AI accelerator.  

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
Here are the reference workloads built on the end-to-end solution and Intel hardware feature provisioning in this project. 
- [OPEA Workloads](workloads/opea/chatqna/README.md)

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
