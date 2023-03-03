// Copyright (c) 2022 - 2023 Intel Corporation
// SPDX-License-Identifier: Apache-2.0 

package sgx 

import (
	"os"
	"path"
	"k8s.io/klog/v2"
)

const (
	devPath = "/dev/"
)

// DeviceCheck checks for sgx enclave and provision device files 
func DeviceCheck() (bool){
    enclavePath := path.Join(devPath,"sgx_enclave")
    provisionPath := path.Join(devPath,"sgx_provision")
    
    if _, err := os.Stat(enclavePath); err != nil{
        klog.Error("SGX enclave file not present: ", err)
        if _, err := os.Stat(provisionPath); err != nil{
            klog.Error("SGX provision file not present: ", err)    
        }
        return false
    }
    if _, err := os.Stat(provisionPath); err != nil{
        klog.Error("SGX provision file not present: ", err)
        return false
    }
    klog.Info("SGX device files present.")
    return true
}
