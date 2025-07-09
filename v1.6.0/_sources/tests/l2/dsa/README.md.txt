### Verify Intel® Data Streaming Accelerator (DSA) Technology provisioning
This workload runs [accel-config](https://github.com/intel/idxd-config) sample tests using RedHat built and distributed accel-config RPM packages from the rhel-9-for-x86_64-baseos-rpms repo. Refer to the [accel config readme](https://github.com/intel/idxd-config/blob/stable/README.md) for more details. 

*	Build the workload container image

Please replace the credentials in buildconfig yaml with your RedHat account login credentials. 

```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/v1.6.0/tests/l2/dsa/dsa_build.yaml 
```

* Create SCC intel-dsa-scc for Intel DSA based workload, if this SCC is not created   
  
```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/v1.6.0/security/dsa_scc.yaml
```
      
* Create the intel-dsa service account to use intel-dsa-scc
  
```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/v1.6.0/security/dsa_rbac.yaml
```

* Deploy the accel-config workload job with intel-dsa service account
  
```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/v1.6.0/tests/l2/dsa/dsa_job.yaml
```

* Check the results.
``` 
  $ oc get pods -n intel-dsa
  intel-dsa-workload-244xm   0/1     Completed   0          3m12s
```

* sample test logs
```
$ oc logs intel-dsa-workload-244xm -n intel-dsa
dsa0/wq0.1
dsa0
Testing with 'block on fault' flag ON
Performing dedicated WQ NOOP testing
Testing 1 bytes
[ info] alloc wq 1 dedicated size 16 addr 0x7f0cde00b000 batch sz 0x400 xfer sz 0x80000000
[ info] testnoop: tflags 0x1 num_desc 1
[ info] preparing descriptor for noop
[ info] Submitted all noop jobs
[ info] verifying task result for 0x2041620
[ info] test with op 0 passed
Testing 4096 bytes
[ info] alloc wq 1 dedicated size 16 addr 0x7fd4881da000 batch sz 0x400 xfer sz 0x80000000
[ info] testnoop: tflags 0x1 num_desc 1
[ info] preparing descriptor for noop
[ info] Submitted all noop jobs
[ info] verifying task result for 0x82f620
[ info] test with op 0 passed
Testing 65536 bytes
[ info] alloc wq 1 dedicated size 16 addr 0x7f462bbed000 batch sz 0x400 xfer sz 0x80000000
[ info] testnoop: tflags 0x1 num_desc 1
[ info] preparing descriptor for noop
[ info] Submitted all noop jobs
[ info] verifying task result for 0xe4e620
[ info] test with op 0 passed
Testing 1048576 bytes
[ info] alloc wq 1 dedicated size 16 addr 0x7fac2ac0c000 batch sz 0x400 xfer sz 0x80000000
[ info] testnoop: tflags 0x1 num_desc 1
[ info] preparing descriptor for noop
[ info] Submitted all noop jobs
[ info] verifying task result for 0xf21620
[ info] test with op 0 passed
Testing 2097152 bytes
[ info] alloc wq 1 dedicated size 16 addr 0x7f7426a5c000 batch sz 0x400 xfer sz 0x80000000
[ info] testnoop: tflags 0x1 num_desc 1
[ info] preparing descriptor for noop
[ info] Submitted all noop jobs
[ info] verifying task result for 0xeec620
[ info] test with op 0 passed
Performing shared WQ NOOP testing
```
