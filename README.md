# Intel® Enterprise AI Foundation for OpenShift*
## Overview
The project delivers a comprehensive full-stack solution for the Intel® Enterprise AI Foundation on the [OpenShift platform](https://www.redhat.com/en/technologies/cloud-computing/openshift/container-platform), applicable across data center, cloud, and edge environments. It utilizes innovative [General Operators technology](https://github.com/intel/intel-technology-enabling-for-openshift/wiki/Intel-Technology-Enabling-for-OpenShift-Architecture-and-Working-Scope#architecture-options) to provision AI accelerators, including the Intel Gaudi Processor, Flex and Max GPUs, and Xeon CPU accelerators such as QAT, SGX, and DSA. Additionally, the project introduces solutions for integrating [Gaudi Software](https://docs.habana.ai/en/latest/index.html) or [OneAPI-based](https://www.intel.com/content/www/us/en/developer/tools/oneapi/overview.html#gs.kgdasr) AI software into OpenShift AI. Key AI workload integrations, such as LLM inferencing, fine-tuning, and post-training for enterprise AI, are under development. The plans also include the GPU network provisioning and full-stack integration with OpenShift. 

![Alt text](/docs/images/Intel-Technology-Enabling-for-OpenShift-Architecture.png)

<div align="center">
  Figure-1 Intel Enterprise AI foundation for OpenShift Architecture 
</div>

## Infrastructure Foundation for Enterprise AI workloads
This project delivers reference infrastructures powered by Intel AI hardware and software technologies, tailored for the cutting-edge **enterprise GenAI** workloads and seamlessly integrated with the [**Red Hat AI platform**](https://www.redhat.com/en/products/ai).

The recommended **Infrastructure Cluster** is built with [**Intel® scalable Gaudi® Accelerator**](https://docs.habana.ai/en/latest/Gaudi_Overview/Gaudi_Architecture.html#gaudi-architecture) and standard servers. The [Intel® Xeon® processors](https://www.intel.com/content/www/us/en/products/details/processors/xeon/xeon6-product-brief.html ) are used in these Gaudi servers as worker nodes and in standard servers as highly available control plane nodes. This infrastructure is designed for **high availability**, **scalability**, and **efficiency** in **Retrieval-Augmented Generation (RAG) and other Large Language Model (LLM) inferencing** workloads.

The [**Gaudi embedded RDMA over Converged Ethernet (RoCE) network**](https://docs.habana.ai/en/latest/PyTorch/PyTorch_Scaling_Guide/Theory_of_Distributed_Training.html#theory-of-distributed-training), along with the [**Three Ply Gaudi RoCE Network topology**](https://docs.habana.ai/en/latest/Management_and_Monitoring/Network_Configuration/Configure_E2E_Test_in_L3.html#generating-a-gaudinet-json-example) supports high-throughput and low latency LLM Parallel Pre-training and Post-training workloads, such as Supervised Fine-Tuning (SFT) and Reinforcement Learning (RL). For more details, see: [LLM Post‐Training Solution with Intel Enterprise AI Foundation for OpenShift](https://github.com/intel/intel-technology-enabling-for-openshift/wiki/Fine-tunning-LLM-Models-with-Intel-Enterprise-AI-Foundation-on-OpenShift)

This highly efficient infrastructure has been validated with cutting-edge enterprise AI workloads on the production-ready OpenShift platform, enabling users to easily evaluate and integrate it into their own AI environments.

Additionally, Intel SGX, DSA, and QAT accelerators (available with Xeon processors) are supported to further enhance performance and security for AI workloads.

For more details, see: [Supported Red Hat OpenShift Container Platform (RHOCP) Infrastructure](https://github.com/intel/intel-technology-enabling-for-openshift/blob/main/docs/supported_platforms.md#supported-intel-hardware-features)

## AI Accelerators & Network Provisioning
Provisioning AI accelerators and networks on a scalable OpenShift/Kubernetes cluster while ensuring the manageability of AI infrastructure and platforms presents significant challenges. To address this, the [general Operator concept](https://github.com/intel/intel-technology-enabling-for-openshift/wiki/Intel-Technology-Enabling-for-OpenShift-Architecture-and-Working-Scope#architecture-options) has been proposed and implemented in this project. 


[OpenShift/Kubernetes Operators](https://www.redhat.com/en/technologies/cloud-computing/openshift/what-are-openshift-operators) automate the management of the software stack, streamlining AI infrastructure provisioning. Rather than relying on a single monolithic operator to handle the entire stack, the [operator best practice](https://sdk.operatorframework.io/docs/best-practices/best-practices/) - **"do one thing and do it well"** - is applied. This industrial-leading approach significantly simplifies both Operator development and the AI provisioning process.

* [**Intel® Network Operator**](https://github.com/intel/network-operator) allows automatic configuring and easier use of RDMA NICs with Intel AI accelerators.
* [**Intel® Device Plugins Operator**](https://catalog.redhat.com/software/container-stacks/detail/61e9f2d7b9cdd99018fc5736) handles the deployment and lifecycle of the device plugins to advertise Intel AI accelerators and other Hardware feature resources to OpenShift/Kubernetes.
*	[**Kernel module management (KMM) operator**](https://github.com/rh-ecosystem-edge/kernel-module-management) manages the deployment and lifecycle of out-of-tree kernel modules like Intel® Data Center GPU Driver for OpenShift*
*	[**Machine config operator (MCO)**](https://github.com/openshift/machine-config-operator) provides an unified interface for the other general operators to configure the Operating System running on the OpenShift nodes.
*	[**Node Feature Discovery (NFD)**](https://docs.redhat.com/en/documentation/openshift_container_platform/4.18/html/specialized_hardware_and_driver_enablement/psap-node-feature-discovery-operator) Operator detects and labels AI hardware features and system configurations. These labels are then used by other general operators.
*	[**Converged AI Operator**]() will be used in the future to simplify the usage of the general operators to provision Intel AI features as a stable and single-entry point.

The Other general Operators can be added in the future to extend the AI features.

## Red Hat AI Platform with Intel AI technologies
Intel and Red Hat have coordinated for years to deliver a production-quality open-source AI platform, built on the best provision of the Intel AI Accelerator computing and networking technologies. 

The **Red Hat AI portfolio**, powered by **Intel AI technologies**, now includes:
* [**Red Hat AI Inference Server**](https://www.redhat.com/en/about/press-releases/red-hat-unlocks-generative-ai-any-model-and-any-accelerator-across-hybrid-cloud-red-hat-ai-inference-server) leverages the [LLM-d](https://github.com/llm-d/llm-d) and [vLLM](https://github.com/vllm-project/vllm) projects, integrating with Llama Stack, Model Context Protocol (MCP), and the Open AI API to deliver standardized APIs for developing and deploying [OPEA-based](https://github.com/opea-project) and other production-grade GenAI applications scalable across edge, enterprise and cloud environments.
* **Red Hat OpenShift AI Distributed Training** provides pre-training, SFT and RL for major GenAI foundation models at scale. With seamless integration of the Kubeflow Training Operator, Intel Gaudi Computing and RoCE Networking technology, enterprises can unlock the full potential of cutting-edge GenAI technologies to drive innovation in their domains.  See [LLM Post‐Training Solution with Intel Enterprise AI Foundation for OpenShift](https://github.com/intel/intel-technology-enabling-for-openshift/wiki/Fine-tunning-LLM-Models-with-Intel-Enterprise-AI-Foundation-on-OpenShift).
* The operators to integrate [Intel Gaudi Software](https://docs.habana.ai/en/latest/index.html) or [OneAPI-based](https://www.intel.com/content/www/us/en/developer/tools/oneapi/overview.html#gs.kgdasr) AI software into OpenShift AI

## Releases and Supported Platforms 
Intel Enterprise AI foundation for OpenShift is released in alignment with the OpenShift release cadence. It is recommended to use the latest release. 

For details on supported features and components, refer to the links below: 
- [Supported Intel AI Accelerator and CPU features](/docs/supported_platforms.md#supported-intel-hardware-features) 
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
