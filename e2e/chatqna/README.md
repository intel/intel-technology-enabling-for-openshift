# OPEA ChatQnA workload on OCP

## Overview
This workload is based on the [OPEA ChatQnA](https://github.com/opea-project/GenAIExamples/tree/main/ChatQnA/deprecated) with [Intel® Gaudi 2 accelerator](https://habana.ai/products/gaudi2/) and [Red Hat OpenShift AI](https://www.redhat.com/en/technologies/cloud-computing/openshift/openshift-ai) (RHOAI) on RHOCP. It uses RAG with Redis vector-db and text generation inference server on RHOAI and Intel Gaudi2 accelerator. Refer to the OPEA readme for more details about the [workload architecture](https://github.com/opea-project/GenAIExamples/tree/main/ChatQnA#chatqna-application). 

## Prerequisites
* Provisioned RHOCP cluster. Follow steps [here](/README.md#provisioning-rhocp-cluster)
* Provisioning Intel Gaudi 2 accelerator:
  * Setup KMM and node feature discovery (NFD) for drivers. Follow the steps [here](https://docs.habana.ai/en/latest/Orchestration/HabanaAI_Operator/Environment_Setup.html)
  * Setup HabanaAI operator and create device config custom Resource. Follow the steps [here](https://docs.habana.ai/en/latest/Orchestration/HabanaAI_Operator/Deploying_HabanaAI_Operator.html)
* Install RHOAI. Follow steps [here](../inference/README.md/#install-rhoai)

## Deploy workload
* The workload uses docker-compose yaml's. They can be converted into OCP compatible yaml's using [Kompose tool](https://kubernetes.io/docs/tasks/configure-pod-container/translate-compose-kubernetes/). This creates the corresponding deployment and service yaml's.

### ChatQnA backend- langchain and redis vector database
* Please follow the [docker-compose file](https://github.com/opea-project/GenAIExamples/blob/main/ChatQnA/deprecated/langchain/docker/docker-compose.yml). To create the backend image, use the [dockerfile](https://github.com/opea-project/GenAIExamples/blob/main/ChatQnA/deprecated/langchain/docker/Dockerfile). Make sure to mount the langchain folder while building image and deploying the pod. The redis vector database image used is ``` redis/redis-stack:7.2.0-v9 ```. Please update the environment variables according to your setup and model variables in the yaml's.

### Minio storage for models
Minio/S3 storage format is supported on RHOAI. To setup minio on OCP, please follow the steps in [documentation](https://ai-on-openshift.io/tools-and-applications/minio/minio/#create-a-matching-data-connection-for-minio). Make sure to allocate enough storage based on the model weights. Once the minio-ui can be accessed, create the bucket and add your model files. Also create an access key and secret key.

### Text Generation Inference for Gaudi on RHOAI
1. Access rhods-dashboard link from routes on your web console and click on Data Science Projects. Now, create a project.
2. Click on the project, and choose ```data connection``` section. In the fields, add your access and secret keys from minio, along with minio-api route (from your minio deployment in previous step). This sets up the connection for your model to be downloaded and prepare for serving.
3. Now click on the Settings and choose ```ServingRuntime```. Copy or import the ```tgi-gaudi-servingruntime.yaml```. The [tgi-gaudi](https://github.com/huggingface/tgi-gaudi) serving runtime is used.
4. Go back to your project and choose ```Models```. Choose the servingruntime and model from the dropdown (it will populated based on the previous step). Choose the number of accelerators (based on servingruntime), size of container and the data connection- minio (populated from previous step). Add the appropriate folder in your bucket for the data connection.
5. Now the model server is in the creation state. Once ready, the status will be updated to green and the inference endpoint can be used to make inference requests.