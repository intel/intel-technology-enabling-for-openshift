# Upgrade Intel Gaudi SPI Firmware
To upgrade Intel Gaudi SPI Firmware, follow below steps: 

**NOTE:** Currently this is only supported on Single Node OpenShift cluster. Multi node cluster support will be added in the future.

## Prerequisites
- Make sure Gaudi drivers are unloaded.
  - On Red Hat OpenShift, delete existing ClusterPolicy Custom Resource. Verify output on the node using below command: 
    ```
    lsmod | grep habana
    ```
  - Check the firmware version following the [firmware version check](https://github.com/intel/intel-technology-enabling-for-openshift/tree/main/tests/gaudi/l2#check-firmware-version-with-hl-smi).

## SPI Firmware Upgrade
Build the container image with `habanalabs-firmware-odm` tool:
```
oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/gaudi/gaudi_spi_fw_upgrade_build.yaml
```
Create service account with required permissions: 
```
oc create sa gaudi-fw-upgrade-sa -n gaudi-spi-fw-upgrade
oc adm policy add-scc-to-user privileged -z gaudi-fw-upgrade-sa -n gaudi-spi-fw-upgrade
```
Deploy and execute the SPI firmware upgrade tool:
```
oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/gaudi/gaudi_spi_fw_upgrade_job.yaml
```

Verify Output:
``` 
oc get pods

NAME                               READY   STATUS      RESTARTS   AGE
gaudi-spi-firmware-upgrade-ndmjp   0/1     Completed   0          10m
```
```
oc logs gaudi-spi-firmware-upgrade-ndmjp
.
.
####
#### Finished sending firmware: OK
```
Verify by following the [firmware version check](https://github.com/intel/intel-technology-enabling-for-openshift/tree/main/tests/gaudi/l2#check-firmware-version-with-hl-smi).