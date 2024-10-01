# Verify Intel® Gaudi® AI Accelerator Provisioning

## hl-smi 
System Management Interface Tool (hl-smi) utility tool obtains information and monitors data of the Intel Gaudi AI accelerators.
`hl-smi` tool is packaged with the Gaudi base image. Run below command to deploy and execute the tool:
```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/tests/gaudi/l2/hl-smi_job.yaml
```

Verify Output:
```
$ oc get pods
NAME                         READY   STATUS      RESTARTS   AGE
hl-smi-workload-2-f5qgs      0/1     Completed   0          27m
```
```
$ oc logs hl-smi-workload-2-f5qgs
+-----------------------------------------------------------------------------+
| HL-SMI Version:                                hl-1.17.1-fw-51.5.0          |
| Driver Version:                                     1.17.1-78932ae          |
|-------------------------------+----------------------+----------------------+
| AIP  Name        Persistence-M| Bus-Id        Disp.A | Volatile Uncorr. ECC |
| Fan  Temp  Perf  Pwr:Usage/Cap|         Memory-Usage | AIP-Util  Compute M. |
|===============================+======================+======================|
|   0  HL-225              N/A  | 0000:4d:00.0     N/A |                   0  |
| N/A   31C   N/A    95W / 600W |    768MiB / 98304MiB |     0%           N/A |
|-------------------------------+----------------------+----------------------+
|   1  HL-225              N/A  | 0000:b4:00.0     N/A |                   0  |
| N/A   28C   N/A    85W / 600W |    768MiB / 98304MiB |     0%           N/A |
|-------------------------------+----------------------+----------------------+
| Compute Processes:                                               AIP Memory |
|  AIP       PID   Type   Process name                             Usage      |
|=============================================================================|
|   0        N/A   N/A    N/A                                      N/A        |
|   1        N/A   N/A    N/A                                      N/A        |
+=============================================================================+
```

## HCCL
HCCL (Habana Collective Communication Library) demo is a program that demonstrates HCCL usage and supports communication via Gaudi based scale-out or Host NIC scale-out. Refer [HCCL Demo](https://github.com/HabanaAI/hccl_demo/tree/main?tab=readme-ov-file#hccl-demo) for more details.

Build the workload container image:
```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/tests/gaudi/l2/hccl_build.yaml
```
Create service account with required permissions: 
```
$ oc create sa hccl-demo-anyuid-sa -n hccl-demo
$ oc adm policy add-scc-to-user anyuid -z hccl-demo-anyuid-sa -n hccl-demo
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