# Supported Red Hat OpenShift Container Platform (RHOCP) Infrastructure 

Before provisioning the Intel hardware and software technology with RHOCP, users need to decide the infrastructure configuration they want to use to install their RHOCP cluster. Red Hat provides support for multiple infrastructure combinations with RHOCP, and this document provides additional details for the specific infrastructure features included in the Intel® Technology Enabling for OpenShift* project. 

## Multi-node bare metal RHOCP cluster provisioning 

The software and hardware technology included in this project are developed, integrated, and tested on the multi-mode bare metal RHOCP 4.14 cluster. Refer to Red Hat’s documentation for instructions to [install a user-provisioned cluster on bare metal](https://docs.openshift.com/container-platform/4.14/installing/installing_bare_metal/installing-bare-metal.html#installation-machine-requirements_installing-bare-metal) and reach out to Red Hat for support in provisioning your bare metal cluster.  

The related container images and operators required to enable the Intel technology included in this project have been certified with Red Hat and published on the [Red Hat Ecosystem Catalog](https://catalog.redhat.com/software). Users identifying issues or feature requests related to the Intel Technology Enabling for OpenShift project are encouraged to submit [issues](https://github.com/intel/intel-technology-enabling-for-openshift/issues) to engage and collaborate with the open source community.  

## Single Node OpenShift (SNO) cluster 

The SNO cluster is supported by RHOCP; however, SNO resource restrictions for some configurations could negatively impact the functionality of the operators included in this project or user workloads. The user should verify that the resource requirements necessary for successful provisioning and execution of the workloads and features are available. Users can run the included images and operators for multi-node bare metal cluster on the RHOCP SNO cluster, but this implementation is not fully tested by this project. Please submit any identified [issues](https://github.com/intel/intel-technology-enabling-for-openshift/issues) and feature requests to the Intel Technology Enabling for OpenShift project to engage and collaborate with the open source community. 

## Cloud Service Provider (CSP) managed RHOCP cluster 

The container images, operators, and documents included in the Intel Technology Enabling for OpenShift project are intended for multi-node bare metal RHOCP cluster provisioning. While RHOCP may be installed with the major CSP environments, refer to Red Hat’s documentation for [selecting an installation method and preparing a cluster](https://docs.openshift.com/container-platform/4.14/installing/installing-preparing.html#installing-preparing-install-manage) for CSP details and consult Red Hat or your CSP for specific support. Any contributions in this area are welcomed.  

## Virtual Machine (VM) based RHOCP cluster 

Multi-node bare metal RHOCP cluster provisioning is the primary cluster use case supported by the project. Virtual Machine technology-based infrastructures support will be included in a future release. For feature requests and related discussions, please submit a project [issue](https://github.com/intel/intel-technology-enabling-for-openshift/issues). For RHOCP infrastructure support, please contact Red Hat. 

## Supported Intel Hardware Features 

The following Intel feature technologies are supported in the 1.2.1 release.  

| Intel Feature Technology                      | Intel Hardware Platform                                        |
|-----------------------------------------------|----------------------------------------------------------------|
| Intel® Software Guard Extensions (Intel® SGX) | [5th Gen Intel® Xeon® Scalable processors](https://www.intel.com/content/www/us/en/products/docs/processors/xeon/5th-gen-xeon-scalable-processors.html) <br/>  [4th Gen Intel® Xeon® Scalable processors](https://www.intel.com/content/www/us/en/products/docs/processors/xeon-accelerated/4th-gen-xeon-scalable-processors.html) <br/>  [3rd Gen Intel® Xeon® Scalable processors   ](https://www.intel.com/content/www/us/en/products/docs/processors/xeon-accelerated/3rd-gen-xeon-scalable-processors.html)                                                   |                               
| Intel® Data Center GPU Flex Series            | [Intel Data Center GPU Flex 140](https://www.intel.com/content/www/us/en/products/sku/230020/intel-data-center-gpu-flex-140/specifications.html) <br/>  [Intel Data Center GPU Flex 170](https://www.intel.com/content/www/us/en/products/sku/230019/intel-data-center-gpu-flex-170/specifications.html)  |
| Intel® Data Center GPU Max Series            | [Intel Data Center GPU Max 1100](https://www.intel.com/content/www/us/en/products/sku/232876/intel-data-center-gpu-max-1100/specifications.html) <br/>  [Intel Data Center GPU Max 1550](https://www.intel.com/content/www/us/en/products/sku/232873/intel-data-center-gpu-max-1550/specifications.html)  |
| Intel® QuickAssist Technology (Intel® QAT) | [5th Gen Intel® Xeon® Scalable processors](https://www.intel.com/content/www/us/en/products/docs/processors/xeon/5th-gen-xeon-scalable-processors.html) <br/>  [4th Gen Intel® Xeon® Scalable processors](https://www.intel.com/content/www/us/en/products/docs/processors/xeon-accelerated/4th-gen-xeon-scalable-processors.html) |                               

## Setting up Intel Hardware Features

### BIOS Configuration
Note: Please refer to your BIOS vendor for specific instructions. This is only a reference for BIOS configuration. 
| Feature | BIOS Configuration | 
| ----- | ---------------------- |
| Intel SGX | [Link](https://www.intel.com/content/www/us/en/support/articles/000087972/server-products/single-node-servers.html) |
| Intel Data Center GPU Flex Series | [Link](https://www.intel.com/content/www/us/en/content-details/774119/virtualization-guide-for-intel-data-center-gpu-flex-series.html?wapkw=gpu%20flex%20series%20setup%20guide) |
| Intel Data Center GPU Max Series | [Link](https://dgpu-docs.intel.com/system-user-guides/DNP-Max-1100-userguide/DNP-Max-1100-userguide.html#bios-setup) |
| Intel QAT | [Link](https://github.com/intel/qatlib/blob/7429ee2b7c837137ed11959a3c2cc3729dc15739/INSTALL#L104) |

## Supported RHOCP Versions

| Feature | Supported RHOCP Version | 
| ----- | ---------------------- |
| Intel SGX |  RHOCP 4.14.11 and above (support is only for z streams in 4.14.z where z >= 11) |
| Intel Data Center GPU Flex Series |  RHOCP 4.14.11 and above (support is only for z streams in 4.14.z where z >= 11) |
| Intel Data Center GPU Max Series |  RHOCP 4.14.11 and above (support is only for z streams in 4.14.z where z >= 11) |
| Intel QAT |  RHOCP 4.14.11 and above (support is only for z streams in 4.14.z where z >= 11) |

## Support 

### Project life cycle introduction 

The Intel Technology Enabling for OpenShift project is based on the [RHOCP life cycle](https://access.redhat.com/support/policy/updates/openshift) and implements the “X.Y.Z” version naming for this project. X.Y version name for this project matches the RHOCP X.Y version. The Z release is intended for hot bug fixes and minor features. 

For example: 

* 1.0.0 release for RHOCP 4.12.0 and 1.0.z is the potential hot bug fixes or minor feature release for RHOCP 4.12.z  

* 1.1.0 release for RHOCP 4.13.0 and 1.1.z is the potential hot bug fixes or minor feature release for RHOCP 4.13.z  

* 1.2.0 release for RHOCP 4.14 and 1.2.z is the potential hot bug fixes or minor feature release for RHOCP 4.14.z  

The Intel Technology Enabling for OpenShift project typically begins a new development and release cycle after the GA of each RHOCP Y release is complete. Expect some delay for the latest RHOCP release to complete component dependency integration with this project and for the required testing and verification necessary for certification with Red Hat. 

### Commercial support from Red Hat 

This project relies on features developed and released with the latest RHOCP release. Commercial RHOCP release support is outlined in the [Red Hat OpenShift Container Platform Life Cycle Policy](https://access.redhat.com/support/policy/updates/openshift) and Intel will collaborate with Red Hat to address specific requirements from our users.  

### Support from the open source community 

Intel Technology Enabling for OpenShift is run as an open source project on GitHub. Project GitHub [issues](https://github.com/intel/intel-technology-enabling-for-openshift/issues) can be used as the primary support interface for users to submit feature requests and report issues to the community when using Intel technology provided by this project. 
