### Verify Intel® Software Guard Extensions (Intel® SGX) Provisioning
This [SampleEnclave](https://github.com/intel/linux-sgx/tree/master/SampleCode/SampleEnclave) application workload from the Intel SGX SDK runs an Intel SGX enclave utilizing the EPC resource from the Intel SGX provisioning.
* Build the container image. 
```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/tests/l2/sgx/sgx_build.yaml
```

* Deploy and run the workload.
  
```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/tests/l2/sgx/sgx_job.yaml
```

*	Check the results.
```
  $ oc get pods
  intel-sgx-job-4tnh5          0/1     Completed   0          2m10s
  intel-sgx-workload-1-build   0/1     Completed   0          30s
```
```
$ oc logs intel-sgx-job-4tnh5
  Checksum(0x0x7fffac6f41e0, 100) = 0xfffd4143
  Info: executing thread synchronization, please wait...
  Info: SampleEnclave successfully returned.
  Enter a character before exit ...
```
## See Also
For Intel SGX demos on vanilla Kubernetes, refer to [link](https://github.com/intel/intel-device-plugins-for-kubernetes/tree/main/demo/sgx-sdk-demo) 
