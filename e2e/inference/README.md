# Intel AI Inference End-to-End Solution

## Overview
Intel AI inference end-to-end solution with RHOCP is based on the Intel® Data Center GPU Flex Series provisioning, Intel® OpenVINO™, and [Red Hat OpenShift AI](https://www.redhat.com/en/technologies/cloud-computing/openshift/openshift-ai) (RHOAI) on RHOCP. There are two AI inference modes verified with Intel® Xeon® processors and Intel Data Center GPU Flex Series with RHOCP.
* Interactive mode – RHOAI provides OpenVINO based Jupyter Notebooks for users to interactively debug the inference applications or [optimize the models](https://docs.openvino.ai/2023.0/openvino_docs_MO_DG_Deep_Learning_Model_Optimizer_DevGuide.html) on RHOCP using data center GPU cards or Intel Xeon processors.
* Deployment mode – [OpenVINO Model Sever](https://github.com/openvinotoolkit/model_server) (OVMS) can be used to deploy the inference workloads in data center and edge computing environments on RHOCP.  

## Prerequisites
* Provisioned RHOCP cluster. Follow steps [here](/README.md#provisioning-rhocp-cluster)
* Provisioning Intel Data Center GPU Flex Series. Follow steps [here](/README.md#provisioning-intel-hardware-features-on-rhocp)
  * Setup node feature discovery (NFD). Follow the steps [here](/nfd/README.md)
  * Setup out of tree drivers for Intel GPU provisioning. Follow the steps [here](/kmmo/README.md)
  * Setup Intel device plugins operator and create Intel GPU device plugin. Follow the steps [here](/device_plugins/README.md)  

## Install RHOAI
The Red Hat certified RHOAI operator is published at [Red Hat Ecosystem Catalog](https://catalog.redhat.com/software/container-stacks/detail/63b85b573112fe5a95ee9a3a). You can use the command line interface (CLI) or web console to install it.
### Install using CLI (To be added)
### Install using Web Console
1.	On the RHOCP web console, click Operators → OperatorHub.
2.	Search RedHat OpenShift AI Operator and click Install. The operator is installed in the namespace `redhat-ods-operator`.
### Verification
1.	Navigate to Operators → Installed Operators page.
2.	Ensure that in the redhat-ods-operator namespace, RedHat OpenShift AI status is InstallSucceeded 
3.	Click on `Search` -> `Routes` -> `rhods-dashboard` from the web console and access the RHOAI UI link.  
**Note:** When installing the operator, the default `kfdef` Custom Resource (CR) is created. This CR enables the dashboard for users to browse and launch Jupyter Notebooks projects on an RHOCP cluster. Please refer to this [link](https://github.com/red-hat-data-services/odh-deployer) for more details about `kfdef`.
## Install OpenVINO Operator
The OpenVINO operator is published at [Red Hat Ecosystem Catalog](https://catalog.redhat.com/software/container-stacks/detail/60649a56209af65d24b7ca9e). You can use the CLI or web console to install it.
### Install using CLI (To be added)
### Install using Web Console
Follow this [link](https://github.com/openvinotoolkit/operator/blob/v1.1.0/docs/operator_installation.md#operator-instalation)  to install the operator via the web console. 

## Work with Interactive Mode
To enable the interactive mode, the OpenVINO notebook CR needs to be created and integrated with RHOAI.  
1.	Click on the `create Notebook` option from the web console and follow these [steps](https://github.com/openvinotoolkit/operator/blob/main/docs/notebook_in_rhods.md#integration-with-openshift-data-science-and-open-data-hub) to create the notebook CR.
2.	Enable Intel Data Center GPU on RHOAI Dashboard- **Technical Preview feature**

Create `AcceleratorProfile` in the `redhat-ods-applications` namespace 
```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/e2e/inference/accelerator_profile.yaml
```

1. Navigate to `openvino-notebooks` ImageStream and add the above created `AcceleratorProfile` key to the annotation field, as shown in the image below:

![Alt text](/docs/images/openvino-accelerator-field.png)

4. Navigate to `Search` -> `Networking` -> `Routes` from the web console and access `rhods-dashboard` route in the `redhat-ods-applications` namespace, as in the image below. Click on the location link to launch RHOAI dashboard. 
   
![Alt text](/docs/images/rhods-dashboard-route.png)
 
5. If step 2 is successful, ```Intel® Data Center GPU Flex Series 140 ``` is shown in the accelerator dropdown menu in ```rhods-dashboard```. Users can run OpenVINO notebook image with Intel® Data Center GPU Flex Series 140 card. 
   
![Alt text](/docs/images/accelerator-profile-dropdown.png)

Follow the [link](https://github.com/openvinotoolkit/operator/blob/main/docs/notebook_in_rhods.md#integration-with-openshift-data-science-and-open-data-hub) for more details on the available Jupyter Notebooks.

## Work with Deployment Mode
1.	From the web console, click on the ModelServer option in this [link](https://github.com/openvinotoolkit/operator/blob/v1.1.0/docs/operator_installation.md#operator-instalation) and follow the [steps](https://github.com/openvinotoolkit/operator/blob/v1.1.0/docs/modelserver.md#managing-model-servers-via-operator) to start the OVMS instance.  
2.	To enable the Intel Data Center GPU, make sure to modify the OVMS instance options according to the screenshot below.

* Below images show `gpu.intel.com/i915` resource requests and limits for OVMS

![Alt text](/docs/images/Ovms-Gpu-resource-limit.png)

![Alt text](/docs/images/Ovms-Gpu-resource-request.png)


## See Also 
[GPU accelerated demo with OpenVINO](https://www.youtube.com/watch?v=3fTz_k4JT2A)