# Copyright (c) 2024 Intel Corporation
# SPDX-License-Identifier: Apache-2.0
#!/bin/sh

tag="v0.8"
namespace="opea-chatqna"
repo="https://github.com/opea-project/GenAIExamples.git"
yaml_url="https://raw.githubusercontent.com/intel/intel-technology-enabling-for-openshift/main/workloads/opea/chatqna/chatqna_megaservice_buildconfig.yaml"

oc project $namespace &&
    git clone --depth 1 --branch $tag $repo && 
        cd GenAIExamples/ChatQnA/deprecated/langchain/docker &&
            oc extract secret/knative-serving-cert -n istio-system --to=. --keys=tls.crt &&
                oc apply -f $yaml_url &&
                    oc start-build chatqna-megaservice --from-dir=./ --follow