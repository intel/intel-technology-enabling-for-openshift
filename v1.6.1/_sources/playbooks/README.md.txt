# Intel Technology Enabling Ansible Playbooks

## Overview
This directory contains Ansible playbooks designed to automate the deployment and configuration of Intel technologies on Red Hat OpenShift clusters. These playbooks streamline the Intel feature provisioning and validation process on OpenShift environments.

## Prerequisites
Before running the playbook, ensure the following prerequisites are met:
- Provisioned RHOCP Cluster
- Red Hat Enterprise Linux (RHEL) system with [Ansible](https://docs.ansible.com/ansible/2.9/installation_guide/intro_installation.html#installing-ansible-on-rhel-centos-or-fedora) installed and configured with a `kubeconfig` to connect to your RHOCP cluster.

## Run the Playbook

To run the ansible playbook, clone this repository to your RHEL system. Navigate to the directory containing the playbook.

```bash
git clone https://github.com/intel/intel-technology-enabling-for-openshift.git

cd intel-technology-enabling-for-openshift/

ansible-playbook playbooks/intel_ocp_provisioning.yaml
```