### dGPU end to end solution
Intel dGPU e2e solution is composed of middleware and framework stack for intel dGPU based data center, cloud, or edge workloads on OCP. Some of the middleware and frameworks supported, include:

##### Intel oneAPI AI Analytics toolkit
The oneAPI AI Analytics toolkit provides tools and frameworks to develop data science and machine learning applications on different architectures and accelerators. On OCP, it is available as the Red Hat certified [oneAPI AI Analytics toolkit operator](https://catalog.redhat.com/software/operators/detail/60d2745bfea5217cbb0f319a). It is also backed by RHODS and OpenDataHub on OCP. Support for this toolkit with dGPU is on schedule for the upcoming releases with focus on AI based use-cases.

##### Intel OpenVINO operator
This toolkit is supported with various frameworks to deploy inference applications and optimize its performance. On OCP, it is available as the Red Hat certified [OpenVINO operator](https://catalog.redhat.com/software/operators/detail/60649a56209af65d24b7ca9e). It is also backed by RHODS and OpenDataHub on OCP. This operator is supported with dGPU for both notebook and OpenVINO Model Server (OVMS) CR'S.
* Once operator is installed, notebook CR is created. From rhods dashboard route, `OpenVINO Toolkit v2022.3` is selected as the jupyter notebook spawner. This starts a pod for notebook server. For dGPU support, please access the `Notebook` object in your cluster and update it with `gpu.intel.com/i915` for `request` and `limit` resource values in the `jupyter-nb-<user>` container. Jupyter notebook is now available to use.
  

##### RHODS (Red Hat OpenShift Data Science)
[RHODS operator](https://catalog.redhat.com/software/container-stacks/detail/63b85b573112fe5a95ee9a3a) is provided on OCP platforms to enable jupyter notebooks. These notebooks provide development environment to create and build machine learning/data science applications. Both oneAPI AI Analytics toolkit and OpenVINO operators are supported by RHODS. The upstream version ODH [OpenDataHub operator](https://github.com/opendatahub-io/opendatahub-operator) can also be used on OCP.  
* After installing RHODS, the kfdef CR is created by default. This CR creates the `rhods-dashboard` route, which is used to launch jupyter notebook server. OpenVINO and oneAPI are one of the options for jupyter notebook spawner.