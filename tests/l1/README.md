### L1 tests overview
Layer 1 tests are used to check for the Intel hardware featurs's devices and drivers (in-tree or out-of-tree) along with setting up the [machine configuration](https://github.com/intel/intel-technology-enabling-for-openshift/blob/main/machine_configuration/README.md) through MCO and OOT drivers through [KMMO](https://github.com/intel/intel-technology-enabling-for-openshift/blob/main/kmmo/README.md) and preparing for (before) the deployment of device plugins through [Intel-device-plugins operator](https://github.com/intel/intel-technology-enabling-for-openshift/blob/main/idpo/README.md). This gives user the information on whether Intel Hardware features are well supported in the OS running on the nodes. Currently, these tests are initial and will be refined in the future releases.

#### Deployment steps
The OCP buildconfig is leveraged to build the source code's container image and push it to the embedded repository through OCP imagestream. It is deployed as a daemonset with an initcontainer for pod completion. Information about Intel hardware features supported on every node in the cluster is the result.

The following deployment steps have been verified on OCP 4.12 baremetal cluster.
* Deploy the scc.yaml which utilizes the custom SCC "container-selinux". It uses the seLinuxOptions "container_device_t" from the [Intel user-container-selinux repo](https://github.com/intel/user-container-selinux). With this option, privileged container requirement is omitted.
 `oc adm policy add-scc-to-user container-scc -z intel-dgpu -n intel-dgpu`

* Deploy the buildconfig to push the image to internal registry
`oc apply -f l1_build.yaml`
* Deploy the job  
`oc apply -f l1_deploy.yaml`
* To check the results on a chosen node:
`oc logs <l1-test-pod>`
* Sample result on node with dGPU support only:
`E0223 20:04:11.864401 1 sgx.go:22] SGX enclave file not present: stat /dev/sgx_enclave: no such file or directory`
`E0223 20:04:11.864438 1 sgx.go:24] SGX provision file not present: stat /dev/sgx_provision: no such file or directory`
`I0223 20:04:11.864504 1 dgpu.go:44] Card and device nodes present for dGPU`