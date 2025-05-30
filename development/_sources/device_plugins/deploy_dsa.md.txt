# Create Intel DSA Device Plugin CR

## Create a CR via web console
1.	Go to **Operator** -> **Installed Operators**.
2.	Open **Intel Device Plugins Operator**.
3.	Navigate to tab **Intel DSA Device Plugin**.
4.	Click **Create DSADevicePlugin** -> set correct parameters -> Click **Create** 
5.	Optional: If you want to make any customizations, select YAML view and edit the details. When you are done, click **Create**.

## Verify via web console
1.	Verify CR by checking the status of **Workloads** -> **DaemonSet** -> **intel-dsa-plugin**.
2.	Now `DsaDevicePlugin` is created.

## Create CR via CLI
Apply the CR yaml file:
```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/device_plugins/dsa_device_plugin.yaml
```

## Verify via CLI
Verify that the device plugin CR is ready: 
```
$ oc get DsaDevicePlugin
```
Output: 
```
NAME		        DESIRED		READY	NODE SELECTOR	                                    AGE
dsadeviceplugin-sample  3 	        3       {"intel.feature.node.kubernetes.io/dsa":"true"}     98m
```

# Verify DSA Device Plugin 
After the plugin is deployed, use below command to verify DSA resources: 
```
$ oc describe node srf-2 | grep dsa.intel.com
  dsa.intel.com/wq-user-dedicated:  0
  dsa.intel.com/wq-user-shared:     160
  dsa.intel.com/wq-user-dedicated:  0
  dsa.intel.com/wq-user-shared:     160
 ```

## DSA Resource Configuration
By default the DSA plugin uses [this configuration file](https://github.com/intel/intel-device-plugins-for-kubernetes/blob/main/demo/dsa.conf).
The dsa init container comes with a utility called `accel-config` which takes file as input and configures the DSA hardwares based on that.
The default configuration has creates dedicated WQs for each DSA device so that it's four groups per device where each groups is with 1 WQ linked to 1 engine.
Users can customise the config and can use the pre-customised config for their specific use case from [here](https://github.com/intel/idxd-config/tree/stable/contrib/configs)
There's also a possibility for a node specific configuration by passing a node specific profile via configMap volume mount.
Users can use the steps below to customize the DSA resource configuration:  
1. Create the configmap for DSA resource configuration 
    ```
    $ oc create configmap --namespace=openshift-operators intel-dsa-config --from-file=dsa[-$NODE_NAME].conf 
2. Create DSA device plugin CR with -provisioning-config set as the name of the ConfigMap (created in step 1) in the dsa_device_plugin.yaml file or set ConfigMap name in the provisioning-config option from web console. 

# Run Intel DSA based workloads on RHOCP
To run the Intel DSA based workloads as an unprivileged pod, you need to use a customised SCC. The customized `dsa-scc` Security Context Constraint (SCC) is provided to bind with service account and run the DSA based workload. 

See [Verify Intel DSA Provisioning](/tests/l2/dsa/README.md) for the detailed steps.  
