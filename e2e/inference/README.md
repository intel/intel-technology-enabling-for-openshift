# Overview
Intel AI inference end-to-end solution with RHOCP is based on the Intel® Data Center GPU Flex Series provisioning, Intel® OpenVINO™, and [Red Hat OpenShift Data Science](https://www.redhat.com/en/technologies/cloud-computing/openshift/openshift-data-science) (RHODS) on RHOCP. There are two AI inference modes verified with Intel® Xeon® processors and Intel Data Center GPU Flex Series with RHOCP-4.12.
* Interactive mode – RHODS provides OpenVINO based Jupyter Notebooks for users to interactively debug the inference applications or [optimize the models](https://docs.openvino.ai/2023.0/openvino_docs_MO_DG_Deep_Learning_Model_Optimizer_DevGuide.html) on RHOCP using data center GPU cards or Intel Xeon processors.
* Deployment mode – [OpenVINO Model Sever](https://github.com/openvinotoolkit/model_server) (OVMS) can be used to deploy the inference workloads in data center and edge computing environments on RHOCP.  
## Prerequisites
* Provisioned RHOCP 4.12 cluster. Follow steps [here](https://github.com/intel/intel-technology-enabling-for-openshift/tree/main#provisioning-rhocp-cluster)
* Provisioning Intel Data Center GPU Flex Series. Follow steps [here](https://github.com/intel/intel-technology-enabling-for-openshift/tree/main#provisioning-intel-hardware-features-on-rhocp)
  * Setup node feature discovery (NFD). Follow the steps [here](https://github.com/intel/intel-technology-enabling-for-openshift/blob/main/nfd/README.md)
  * Setup machine configuration. Follow the steps [here](https://github.com/intel/intel-technology-enabling-for-openshift/blob/main/machine_configuration/README.md)
  * Setup out of tree drivers for Intel GPU provisioning. Follow the steps [here](https://github.com/intel/intel-technology-enabling-for-openshift/blob/main/machine_configuration/README.md)
  * Setup Intel device plugins operator and create Intel GPU device plugin. Follow the steps [here](https://github.com/intel/intel-technology-enabling-for-openshift/blob/main/device_plugins/README.md)  

## Install RHODS
The Red Hat certified RHODS operator is published at [Red Hat Ecosystem Catalog](https://catalog.redhat.com/software/container-stacks/detail/63b85b573112fe5a95ee9a3a). You can use the command line interface (CLI) or web console to install it.
### Install using CLI (To be added)
### Install using web console
1.	On the RHOCP web console, click Operators → OperatorHub.
2.	Search RedHat OpenShift Data Science Operator and click Install. The operator is installed in the namespace redhat-ods-operator.
### Verification
1.	Navigate to Operators → Installed Operators page.
2.	Ensure that in the redhat-ods-operator namespace, RedHat OpenShift Data Science Status is InstallSucceeded 
3.	Click on ```Search -> Routes -> rhods-dashboard``` from the web console and access the RHODS UI link.
Note: When installing the operator, the default ```kfdef``` Custom Resource (CR) is created. This CR enables the RHODS dashboard for users to browse and launch Jupyter Notebooks projects on an RHOCP cluster. Please refer to this [link](https://github.com/red-hat-data-services/odh-deployer) for more details about kfdef.
## Install OpenVINO operator
The OpenVINO operator is published at [Red Hat Ecosystem Catalog](https://catalog.redhat.com/software/container-stacks/detail/60649a56209af65d24b7ca9e). You can use the CLI or web console to install it.
### Install using CLI (To be added)
### Install using web console
Follow this [link](https://github.com/openvinotoolkit/operator/blob/v1.1.0/docs/operator_installation.md)  to install the operator via the web console. 
## Work with interactive mode
To enable the interactive mode, the OpenVINO notebook CR needs to be created and integrated with RHODS.  
1.	Click on the ```create Notebook``` option in this [link](https://github.com/red-hat-data-services/odh-deployer) from the web console and follow these [steps](https://github.com/openvinotoolkit/operator/blob/main/docs/notebook_in_rhods.md) to create the notebook CR.
2.	Enable Intel Data Center GPU 

**Note:** The  Intel Data Center GPU option is not visible in the RHODS UI at this time of release. For more details, please refer to this [issue](https://github.com/opendatahub-io/odh-dashboard/issues/956). Until this issue is resolved, please follow the steps below to enable the Intel Data Center GPU.

a.	Search for the OpenVINO Notebook Server from web console ```Search -> Notebook -> Jupyter-nb-<ocp-user>``` in the namespace ```rhods-notebooks```.

b.	Navigate to notebook yaml and modify the yaml file according to the example shown below.

 ```   
    containers:
       name: jupyter-nb-<ocp-user-name>
       resources:
         limits:
           cpu: '14'
           gpu.intel.com/i915: '1'
           memory: 56Gi
         requests: 
           cpu: '7'
           gpu.intel.com/i915: '1'
           memory: 56Gi
  ```  

c.	This operation respawns the notebook server to use the Intel Data Center GPU.

3.	Run sample Jupyter Notebooks.
Follow the [link](https://github.com/openvinotoolkit/operator/blob/main/docs/notebook_in_rhods.md) to execute the sample Jupyter Notebook There are 60+ sample notebooks 	available with this notebook image. For the details on the notebooks with Intel Data Center 	GPU, please check this [link](https://github.com/openvinotoolkit/openvino_notebooks).
## Work with deployment mode
1.	From the web console, click on the ModelServer option in this [link](https://github.com/openvinotoolkit/operator/blob/v1.1.0/docs/operator_installation.md) and follow the [steps](https://github.com/openvinotoolkit/operator/blob/v1.1.0/docs/modelserver.md) to start the OVMS instance.  
2.	To enable the Intel Data Center GPU, make sure to modify the OVMS instance options according to the screenshot below.

* Below images show gpu.intel.com/i915 resource requests and limits for OVMS

![Alt text](/docs/images/Ovms-Gpu-resource-limit.png)

![Alt text](/docs/images/Ovms-Gpu-resource-request.png)


## See Also 
[GPU accelerated demo with OpenVINO](https://www.youtube.com/watch?v=3fTz_k4JT2A)
