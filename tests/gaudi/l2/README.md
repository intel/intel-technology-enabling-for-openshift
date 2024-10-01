# Verify Intel® Gaudi® AI Accelerator Provisioning

## hl-smi 
System Management Interface Tool (hl-smi) utility tool obtains information and monitors data of the Intel Gaudi AI accelerators.
`hl-smi` tool is packaged with the Gaudi base image. Run below command to deploy and execute the tool:
```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/v1.4.0/tests/gaudi/l2/hl-smi_job.yaml
```

Verify Output:
```
$ oc get pods -n gaudi-validation
NAME                         READY   STATUS      RESTARTS   AGE
hl-smi-workload-2-f5qgs      0/1     Completed   0          27m
```
```
$ oc logs hl-smi-workload-2-f5qgs -n gaudi-validation
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
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/v1.4.0/tests/gaudi/l2/hccl_build.yaml
```
Create service account with required permissions: 
```
$ oc create sa hccl-demo-anyuid-sa -n gaudi-validation
$ oc adm policy add-scc-to-user anyuid -z hccl-demo-anyuid-sa -n gaudi-validation
```
Deploy and execute the workload:
```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/v1.4.0/tests/gaudi/l2/hccl_job.yaml
```

Verify Output:
``` 
$ oc get pods -n gaudi-validation
NAME                         READY   STATUS      RESTARTS   AGE
hccl-demo-workload-1-build   0/1     Completed   0          23m
hccl-demo-workload-wq8mx     0/1     Completed   0          10m
```
```
$ oc logs hccl-demo-workload-wq8mx -n gaudi-validation
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

## vLLM 
vLLM is a serving engine for LLM's. The following workloads deploys a VLLM server with an LLM using Intel Gaudi. Refer to [Intel Gaudi vLLM fork](https://github.com/HabanaAI/vllm-fork.git) for more details.

Use the gaudi-validation project
```
$ oc project gaudi-validation
```
Build the workload container image:
```
git clone https://github.com/HabanaAI/vllm-fork.git --branch v1.18.0

cd vllm-fork/

$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/tests/gaudi/l2/vllm_buildconfig.yaml

$ oc start-build vllm-workload --from-dir=./ --follow
```
Check if the build has completed
```
$ oc get builds
NAMESPACE           NAME               TYPE     FROM         STATUS     STARTED       DURATION
gaudi-validation   vllm-workload-1    Docker   Dockerfile   Complete   7 minutes ago   4m58s

```

Deploy the workload:
* Update the hugging face token in the ```vllm_hf_secret.yaml``` file, refer to [link](https://huggingface.co/docs/hub/en/security-tokens) for more details. 
```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/tests/gaudi/l2/vllm_hf_secret.yaml
```
meta-llama/Llama-3.1-8B model is used in this deployment and the hugging face token is used to access such gated models.
* For the PV setup with NFS, refer to [documentation](https://docs.openshift.com/container-platform/4.17/storage/persistent_storage/persistent-storage-nfs.html).
* The vLLM pod needs to access the host's shared memory for tensor parallel inference, which is mounted as a volume.
```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/tests/gaudi/l2/vllm_deployment.yaml
```
Create the vllm service
```
oc expose deploy/vllm-workload
```
Verify Output:
``` 
$ oc get pods
NAME                             READY   STATUS      RESTARTS   AGE
vllm-workload-1-build            0/1     Completed   0          19m
vllm-workload-55f7c6cb7b-cwj2b   1/1     Running     0          8m36s

$ oc get svc
NAME            TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)    AGE
vllm-workload   ClusterIP   1.2.3.4          <none>        8000/TCP   114s
```
```
$ oc logs vllm-workload-55f7c6cb7b-cwj2b 

INFO 10-30 19:35:53 habana_model_runner.py:95] VLLM_DECODE_BS_BUCKET_MIN=32 (default:min)
INFO 10-30 19:35:53 habana_model_runner.py:95] VLLM_DECODE_BS_BUCKET_STEP=32 (default:step)
INFO 10-30 19:35:53 habana_model_runner.py:95] VLLM_DECODE_BS_BUCKET_MAX=256 (default:max)
INFO 10-30 19:35:53 habana_model_runner.py:95] VLLM_PROMPT_SEQ_BUCKET_MIN=128 (default:min)
INFO 10-30 19:35:53 habana_model_runner.py:95] VLLM_PROMPT_SEQ_BUCKET_STEP=128 (default:step)
INFO 10-30 19:35:53 habana_model_runner.py:95] VLLM_PROMPT_SEQ_BUCKET_MAX=1024 (default:max)
INFO 10-30 19:35:53 habana_model_runner.py:95] VLLM_DECODE_BLOCK_BUCKET_MIN=128 (default:min)
INFO 10-30 19:35:53 habana_model_runner.py:95] VLLM_DECODE_BLOCK_BUCKET_STEP=128 (default:step)
INFO 10-30 19:35:53 habana_model_runner.py:95] VLLM_DECODE_BLOCK_BUCKET_MAX=4096 (default:max)
INFO 10-30 19:35:53 habana_model_runner.py:691] Prompt bucket config (min, step, max_warmup) bs:[1, 32, 64], seq:[128, 128, 1024]
INFO 10-30 19:35:53 habana_model_runner.py:696] Decode bucket config (min, step, max_warmup) bs:[32, 32, 256], block:[128, 128, 4096]
============================= HABANA PT BRIDGE CONFIGURATION ===========================
PT_HPU_LAZY_MODE = 1
PT_RECIPE_CACHE_PATH =
PT_CACHE_FOLDER_DELETE = 0
PT_HPU_RECIPE_CACHE_CONFIG =
PT_HPU_MAX_COMPOUND_OP_SIZE = 9223372036854775807
PT_HPU_LAZY_ACC_PAR_MODE = 1
PT_HPU_ENABLE_REFINE_DYNAMIC_SHAPES = 0
PT_HPU_EAGER_PIPELINE_ENABLE = 1
PT_HPU_EAGER_COLLECTIVE_PIPELINE_ENABLE = 1
---------------------------: System Configuration :---------------------------
Num CPU Cores : 160
CPU RAM : 1056371848 KB
------------------------------------------------------------------------------
INFO 10-30 19:35:56 selector.py:85] Using HabanaAttention backend.
INFO 10-30 19:35:56 loader.py:284] Loading weights on hpu ...
INFO 10-30 19:35:56 weight_utils.py:224] Using model weights format ['*.safetensors', '*.bin', '*.pt']
Loading safetensors checkpoint shards: 0% Completed | 0/4 [00:00<?, ?it/s]
Loading safetensors checkpoint shards: 25% Completed | 1/4 [00:03<00:11, 3.87s/it]
Loading safetensors checkpoint shards: 50% Completed | 2/4 [00:07<00:07, 3.71s/it]
Loading safetensors checkpoint shards: 75% Completed | 3/4 [00:10<00:03, 3.59s/it]
Loading safetensors checkpoint shards: 100% Completed | 4/4 [00:11<00:00, 2.49s/it]
Loading safetensors checkpoint shards: 100% Completed | 4/4 [00:11<00:00, 2.93s/it]
```

* The internal service url is used to run inference requests to the vLLM server. This service is only accessible from pods running within the same namespace i.e gaudi-validation. Run the below commands to create a sample pod and run requests.

```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/tests/gaudi/l2/test-pod.yaml
```

Check for the pod

```
$ oc get pods
NAME                           READY   STATUS      RESTARTS   AGE
test                           1/1     Running     0          2s
```

Use the command below to enter pod terminal to run curl requests

```
$ oc debug pod/test
```

```
sh-5.1# curl  "http://vllm-workload.gaudi-validation.svc.cluster.local:8000/v1/models"{"object":"list","data":[{"id":"meta-llama/Llama-3.1-8B","object":"model","created":1730317412,"owned_by":"vllm","root":"meta-llama/Llama-3.1-8B","parent":null,"max_model_len":131072,"permission":[{"id":"modelperm-452b2bd990834aa5a9416d083fcc4c9e","object":"model_permission","created":1730317412,"allow_create_engine":false,"allow_sampling":true,"allow_logprobs":true,"allow_search_indices":false,"allow_view":true,"allow_fine_tuning":false,"organization":"*","group":null,"is_blocking":false}]}]}
```

```
sh-5.1# curl http://vllm-workload.gaudi-validation.svc.cluster.local:8000/v1/completions   -H "Content-Type: application/json"   -d '{
        "model": "meta-llama/Llama-3.1-8B",
        "prompt": "A constellation is a",
        "max_tokens": 10
      }'
{"id":"cmpl-9a0442d0da67411081837a3a32a354f2","object":"text_completion","created":1730321284,"model":"meta-llama/Llama-3.1-8B","choices":[{"index":0,"text":" group of individual stars that forms a pattern or figure","logprobs":null,"finish_reason":"length","stop_reason":null}],"usage":{"prompt_tokens":5,"total_tokens":15,"completion_tokens":10}}
```

## Check firmware version with hl-smi
System Management Interface Tool (hl-smi) utility tool obtains information and monitors data of the Intel Gaudi AI accelerators.
Run below command to check firmware version with the tool:
```
$ oc apply -f https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/tests/gaudi/l2/hl-smi-firmware_job.yaml
```

Verify Output:
```
$ oc get pods -n gaudi-validation
NAME                      READY   STATUS      RESTARTS   AGE
hl-smi-firmware-pxhsn     0/1     Completed   0          11s
```
```
$ oc logs hl-smi-firmware-pxhsn -n gaudi-validation
        Firmware [SPI] Version          : Preboot version hl-gaudi2-1.16.0-fw-50.1.2-sec-9 (May 26 2024 - 11:33:04)
```
