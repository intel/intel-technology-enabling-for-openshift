### Verify Intel® QuickAssist Technology provisioning
This workload runs [qatlib](https://github.com/intel/qatlib) sample tests using RedHat built and distributed Qatlib RPM packages from the codeready-builder-for-rhel-9-x86_64-rpms repo. Refer to the [qatlib readme](https://github.com/intel/qatlib/blob/main/INSTALL) for more details. 

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

* Deploy the qatlib workload job with intel-qat service account
  
```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/tests/l2/qat/qatlib_job.yaml
```

* Check the results.
``` 
  $ oc get pods
  intel-qat-workload-c6g9v   0/1     Completed   0          4m13s
```


* For all sample tests `cpa_sample_code` 

```
$ oc logs intel-qat-workload-c6g9v
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
