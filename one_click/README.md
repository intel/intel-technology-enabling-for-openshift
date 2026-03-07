# Deploy Intel Technology Enabling Solutions with Red Hat OpenShift using “One-Click”

## Overview
Red Hat [Ansible](https://www.ansible.com/) and Operator technologies are used for “One-Click Deployment” of Intel technology enabling solutions with Red Hat OpenShift Container Platform (RHOCP). Ansible technology automates the operator installation and configuration steps using a playbook, making deployment as simple as a single click.

The referenced Ansible playbooks here can be used by the cluster administrators to customize their own playbooks.

**Note:** It is recommended to start from [Get started](/README.md#getting-started) to get familiar with the installation and configuration of the general operator before composing the first playbook.

## Reference Playbook – Intel Data Center GPU Provisioning

This playbook demonstrates the one-click provisioning of Intel Data Center GPU  on an RHOCP cluster. The steps involved are installation and configuration of general Operators including Node Feature Discovery (NFD) operator, Kernel Module Management (KMM) operator, and the Intel Device Plugins Operator. 

### Prerequisite
Before running the playbook, ensure the following prerequisites are met:
- Provisioned RHOCP Cluster
- Red Hat Enterprise Linux (RHEL) system with [Ansible](https://docs.ansible.com/ansible/2.9/installation_guide/intro_installation.html#installing-ansible-on-rhel-centos-or-fedora) installed and configured with a `kubeconfig` to connect to your RHOCP cluster.

### Run the Playbook
To run the ansible playbook, clone this repository to your RHEL system. Navigate to the directory containing the playbook. 
```
$ git clone https://github.com/intel/intel-technology-enabling-for-openshift.git
$ cd intel-technology-enabling-for-openshift/one_click 
```
Execute below single command to provision Intel Data Center GPU:
```
$ ansible-playbook gpu_provisioning_playbook.yaml
```

## Reference Playbook – Intel Gaudi Provisioning
This playbook demonstrates the one-click provisioning of Intel Gaudi AI Accelerator on an RHOCP cluster. The steps involved are installation and configuration of general Operators including Node Feature Discovery (NFD) operator, Kernel Module Management (KMM) operator, and the Intel Gaudi Base Operator. The playbook also creates the Gaudi `DeviceConfig` CR which deploys the Gaudi Out-of-Tree drivers, Gaudi device plugins, Habana container runtime and Habana node metrics.

### Prerequisite
Before running the playbook, ensure the following prerequisites are met:
- Provisioned RHOCP Cluster
- Red Hat Enterprise Linux (RHEL) system with [Ansible](https://docs.ansible.com/ansible/2.9/installation_guide/intro_installation.html#installing-ansible-on-rhel-centos-or-fedora) installed and configured with a `kubeconfig` to connect to your RHOCP cluster.
- Set Firmware search path using MCO, follow [Update Kernel Firmware Search Path with MCO](/gaudi/README.md#update-kernel-firmware-search-path-with-mco).

### Run the Playbook
To run the ansible playbook, clone this repository to your RHEL system. Navigate to the directory containing the playbook. 
```
$ git clone https://github.com/intel/intel-technology-enabling-for-openshift.git
$ cd intel-technology-enabling-for-openshift/one_click 
```
Execute below single command to provision Intel Gaudi Accelerator:
```
$ ansible-playbook gaudi_provisioning_playbook.yaml
```

## Reference Playbook – Intel Gaudi Provisioning and SW Stack Validation
This playbook demonstrates the one-click solution to validate Gaudi provisioning and software stack on RHOCP. This playbook validates L1, L2 and L3 level test cases in [Verify Intel® Gaudi® AI Accelerator Provisioning](/tests/gaudi/l2/README.md).

### Prerequisite
Before running the playbook, ensure the following prerequisites are met:
- Provisioned RHOCP cluster
- RHOCP cluster with provisioned with Intel Gaudi Base Operator. Refer [Setting up Intel Gaudi Base Operator](/gaudi/README.md#setting-up-intel-gaudi-base-operator)

### Run the Playbook
To run the ansible playbook, clone this repository to your RHEL system. Navigate to the directory containing the playbook. 
```
$ git clone https://github.com/intel/intel-technology-enabling-for-openshift.git
$ cd intel-technology-enabling-for-openshift/one_click 
```
Execute below single command to Validate Intel Gaudi accelerator provisioning and SW stack:
```
$ ansible-playbook gaudi_validation_playbook.yaml
```