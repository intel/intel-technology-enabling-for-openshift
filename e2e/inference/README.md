### Intel AI Inference E2E test for OpenShift

#### Overview
Intel AI inference e2e test utilizes Intel® Data Center GPU Flex Series provisioning on OpenShift. For provisioning details, refer to [KMM](../../kmmo/README.md#managing-intel-dgpu-driver-with-kmm-operator). The following AI inference modes support the Intel Data Center GPU Flex Series:
* **Interactive Mode**
[Red Hat OpenShift Data Science (RHODS)](https://www.redhat.com/en/technologies/cloud-computing/openshift/openshift-data-science) helps spawn the Intel® OpenVINO based [Jupyter Notebook](https://jupyter.org/). It can be used to interactively debug the inference applications or optimize the models on OCP.
*	**Deployment Mode**
[Intel OpenVINO™ Toolkit](https://www.intel.com/content/www/us/en/developer/tools/openvino-toolkit/overview.html) and [Operator](https://github.com/openvinotoolkit/operator) provide the [OpenVINO Model Server (OVMS)](https://github.com/openvinotoolkit/model_server) to deploy and infer workloads on OCP.
#### Prerequisites to deploy Intel AI Inference E2E test

##### Install RHODS on OpenShift
* Red Hat OpenShift Data Science (RHODS) enables Jupyter notebooks on OCP. These notebooks provide development environments to create machine learning and data science applications. For installation steps, refer to [RHODS installation](https://catalog.redhat.com/software/container-stacks/detail/63b85b573112fe5a95ee9a3a). The upstream community operator [OpenDataHub operator (ODH)](https://github.com/opendatahub-io/opendatahub-operator) also provides notebooks and is supported on OCP.

##### Install Intel OpenVINO Toolkit Operator
* [Intel OpenVINO toolkit](https://github.com/openvinotoolkit/openvino) includes tools and runtimes to deploy inference applications and optimize performance. For more details, refer to [OpenVINO operator](https://catalog.redhat.com/software/operators/detail/60649a56209af65d24b7ca9e). It uses RHODS operator to spawn OpenVINO jupyter notebooks. Intel GPU is supported for both `Notebook` and `OpenVINO Model Server (OVMS)` Custom Resources.

#### Run Interactive Mode Demo
Once the OpenVINO Jupyter notebook server is spawned with GPU, sample [OpenVINO notebooks](https://github.com/openvinotoolkit/openvino_notebooks) can be used to verify. The following notebooks have been verified to work with Intel Data Center GPU Flex Series on OCP 4.12:
  * [004-hello-detection](https://github.com/openvinotoolkit/openvino_notebooks/tree/main/notebooks/004-hello-detection)    
  * [101-tensorflow-to-openvino](https://github.com/openvinotoolkit/openvino_notebooks/tree/main/notebooks/101-tensorflow-to-openvino)
  * [104-model-tools](https://github.com/openvinotoolkit/openvino_notebooks/tree/main/notebooks/104-model-tools)
  * [108-gpu-device](https://github.com/openvinotoolkit/openvino_notebooks/tree/main/notebooks/108-gpu-device) 
  * [110-ct-segmentation-quantize](https://github.com/openvinotoolkit/openvino_notebooks/blob/main/notebooks/110-ct-segmentation-quantize/110-ct-scan-live-inference.ipynb)     
  * [115-async-api](https://github.com/openvinotoolkit/openvino_notebooks/tree/main/notebooks/115-async-api) 
  * [213-question-answering](https://github.com/openvinotoolkit/openvino_notebooks/tree/main/notebooks/213-question-answering) 
  * [215-image-inpainting](https://github.com/openvinotoolkit/openvino_notebooks/tree/main/notebooks/215-image-inpainting)  
  * [216-license-plate-recognition](https://github.com/openvinotoolkit/openvino_notebooks/tree/main/notebooks/216-license-plate-recognition)
  * [218-vehicle-detection-and-recognition](https://github.com/openvinotoolkit/openvino_notebooks/tree/main/notebooks/218-vehicle-detection-and-recognition)
  * [401-object-detection-webcam](https://github.com/openvinotoolkit/openvino_notebooks/tree/main/notebooks/401-object-detection-webcam)
  

#### Run Deployment Mode Demo 
* **Note**: This mode's verification is ongoing