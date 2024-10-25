# Upgrade Intel Gaudi SPI Firmware
To upgrade Intel Gaudi SPI Firmware, follow below steps: 

## Prerequisites
- Make sure Gaudi drivers are unloaded

## SPI Firmware Upgrade
Build the container image with `habanalabs-firmware-odm` tool:
```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/gaudi/gaudi_spi_fw_upgrade_build.yaml
```
Create service account with required permissions: 
```
$ oc create sa gaudi-fw-upgrade-sa -n gaudi-validation
$ oc adm policy add-scc-to-user privileged -z gaudi-fw-upgrade-sa -n gaudi-validation
```
Deploy and execute the SPI firmware upgrade tool:
```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/gaudi/gaudi_spi_fw_upgrade_job.yaml
```

Verify Output:
``` 
$ oc get pods
NAME                               READY   STATUS      RESTARTS   AGE
gaudi-spi-firmware-upgrade-ndmjp   0/1     Completed   0          10m
```
```
$ oc logs gaudi-spi-firmware-upgrade-ndmjp
.
.
####
#### Finished sending firmware: OK
```
Verify using `hl-smi`: 
```
sh-5.1$ hl-smi -L | grep SPI
    Firmware [SPI] Version          : Preboot version hl-gaudi2-1.18.0-fw-53.1.1-sec-9 (Oct 02 2024 - 11:52:39)
    Firmware [SPI] Version          : Preboot version hl-gaudi2-1.18.0-fw-53.1.1-sec-9 (Oct 02 2024 - 11:52:39)
    Firmware [SPI] Version          : Preboot version hl-gaudi2-1.18.0-fw-53.1.1-sec-9 (Oct 02 2024 - 11:52:39)
    Firmware [SPI] Version          : Preboot version hl-gaudi2-1.18.0-fw-53.1.1-sec-9 (Oct 02 2024 - 11:52:39)
    Firmware [SPI] Version          : Preboot version hl-gaudi2-1.18.0-fw-53.1.1-sec-9 (Oct 02 2024 - 11:52:39)
    Firmware [SPI] Version          : Preboot version hl-gaudi2-1.18.0-fw-53.1.1-sec-9 (Oct 02 2024 - 11:52:39)
    Firmware [SPI] Version          : Preboot version hl-gaudi2-1.18.0-fw-53.1.1-sec-9 (Oct 02 2024 - 11:52:39)
    Firmware [SPI] Version          : Preboot version hl-gaudi2-1.18.0-fw-53.1.1-sec-9 (Oct 02 2024 - 11:52:39)
```