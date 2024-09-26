# Verify Intel® Gaudi® AI Accelerator Provisioning

## HCCL
HCCL (Habana Collective Communication Library) demo is a program that demonstrates HCCL usage and supports communication via Gaudi based scale-out or Host NIC scale-out. Refer [HCCL Demo](https://github.com/HabanaAI/hccl_demo/tree/main?tab=readme-ov-file#hccl-demo) for more details.

Build the workload container image:
```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/tests/gaudi/l2/hccl_build.yaml
```
Deploy and execute the workload:
```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/tests/gaudi/l2/hccl_job.yaml
```

Verify Output:
``` 
$ oc get pods
NAME                         READY   STATUS      RESTARTS   AGE
hccl-demo-workload-1-build   0/1     Completed   0          23m
hccl-demo-workload-wq8mx     0/1     Completed   0          10m
```
```
$ oc logs hccl-demo-workload-wq8mx
Affinity: Numa mapping directory: /tmp/affinity_topology_output
Affinity: Script has not been executed before, going to execute...
Affinity: Script has finished successfully
Welcome to HCCL demo
.
.
.
####################################################################################################
[BENCHMARK] hcclAllReduce(src!=dst, data_size=33554432, count=8388608, dtype=float, iterations=1000)
[BENCHMARK]     NW Bandwidth   : 258.209121 GB/s
[BENCHMARK]     Algo Bandwidth : 147.548069 GB/s
####################################################################################################
```