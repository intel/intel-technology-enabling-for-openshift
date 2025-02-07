### Verify Intel® QuickAssist Technology provisioning
This workload runs [qatlib](https://github.com/intel/qatlib) sample tests using RedHat built and distributed Qatlib RPM packages from the codeready-builder-for-rhel-9-x86_64-rpms repo. Refer to the [qatlib readme](https://github.com/intel/qatlib/blob/main/INSTALL) for more details. 
### As a prerequisite - configure the OCP 4.16 QAT nodes to be do sym-dc :
The ConfigMap to enable sym-dc since there is no way to have cy and dc in two different endPoints 
https://github.com/intel/intel-device-plugins-for-kubernetes/blob/5f305b4a96314c5486256baa330434ad48d147e3/cmd/qat_plugin/dpdkdrv/dpdkdrv.go#L418
or
https://github.com/intel/intel-device-plugins-for-kubernetes/blob/b19c91f8f1b935edade27e46848686025f3d96db/cmd/qat_plugin/dpdkdrv/dpdkdrv.go#L449

:
```
apiVersion: v1
data:
  qat.conf: ServicesEnabled=sym;dc
kind: ConfigMap
metadata:
  name: qat-services-cm
  namespace: openshift-operators

```
*	Build the workload container image

Please replace the credentials in buildconfig yaml with your RedHat account login credentials. 

```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/tests/l2/qat/qatlib_build.yaml 
```

* Create SCC intel-qat-scc for Intel QAT based workload, if this SCC is not created   
  
```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/security/qatlib_scc.yaml
```
      
* Create the intel-qat service account to use intel-qat-scc
  
```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/security/qatlib_rbac.yaml
```
Please replace above qatlib_rbac.yaml with this (namespace: intel-qat goes into metadata:)
```
# Copyright (c) 2023 Intel Corporation
# SPDX-License-Identifier: Apache-2.0
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: intel-qat
  namespace: intel-qat
---
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: intel-qat
  namespace: intel-qat
rules:
- apiGroups:
  - security.openshift.io
  resources:
  - securitycontextconstraints
  resourceNames:
  - intel-qat-scc
  verbs:
  - use
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: intel-qat
  namespace: intel-qat
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: intel-qat
subjects:
- kind: ServiceAccount
  name: intel-qat
  #namespace: intel-qat

```

* Deploy the qatlib workload job with intel-qat service account

in the bellow job replace the request and limits for cy: 1 and dc: 1 by one single line in each:
qat.intel.com/sym-dc: '2' #Or maybe 1 . can work?
  
```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/tests/l2/qat/qatlib_job.yaml
```

* Check the results.
``` 
  $ oc get pods -n intel-qat
  intel-qat-workload-c6g9v   0/1     Completed   0          4m13s
```


* For all sample tests `cpa_sample_code` 

```
$ oc logs intel-qat-workload-c6g9v -n intel-qat
qaeMemInit started
icp_sal_userStartMultiProcess("SSL") started
There are no crypto instances
*** QA version information ***
device ID               = 0
software                = 23.2.0
*** END QA version information ***
Inst 0, Affin: 0, Dev: 0, Accel 0, EE 0, BDF ED:00:01
Inst 1, Affin: 1, Dev: 0, Accel 0, EE 0, BDF ED:00:01
Inst 2, Affin: 2, Dev: 0, Accel 0, EE 0, BDF ED:00:01
Inst 3, Affin: 3, Dev: 0, Accel 0, EE 0, BDF ED:00:01
---------------------------------------
API                    Traditional
Session State          STATELESS
Algorithm              DEFLATE
Huffman Type           STATIC
Mode                   ASYNCHRONOUS
CNV Enabled            YES
Direction              COMPRESS
Packet Size            8192
Compression Level      1
Corpus                 CALGARY_CORPUS
Corpus Filename        calgary
CNV Recovery Enabled   YES
Number of threads      4
Total Responses        158400
Total Retries          2242671
Clock Cycles Start     126150916653843
Clock Cycles End       126151409143747
Total Cycles           492489904
CPU Frequency(kHz)     1700160
Throughput(Mbps)       35920
Compression Ratio      0.4897
---------------------------------------

Inst 0, Affin: 0, Dev: 0, Accel 0, EE 0, BDF ED:00:01
Inst 1, Affin: 1, Dev: 0, Accel 0, EE 0, BDF ED:00:01
Inst 2, Affin: 2, Dev: 0, Accel 0, EE 0, BDF ED:00:01
Inst 3, Affin: 3, Dev: 0, Accel 0, EE 0, BDF ED:00:01
---------------------------------------
```
