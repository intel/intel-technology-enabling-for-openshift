// Copyright (c) 2022 - 2023 Intel Corporation
// SPDX-License-Identifier: Apache-2.0 

package dgpu  

import ( 
	"os"
	"regexp"
	"k8s.io/klog/v2"
)

const (
	devDriPath = "/dev/dri/"
	gpuCard = `^card[0-9]+$`
	gpuRender = `^renderD[0-9]+$`
)

// NodeCheck checks for gpu represented devices
func NodeCheck() (bool){
	devFiles, err := os.ReadDir(devDriPath)
	gpuCardExists := false
	gpuRenderExists := false
	if err != nil {
		klog.Error("/dev/dri path not present on node: ", err)
        return false
    }

	gpuCardReg:= regexp.MustCompile(gpuCard)
	gpuRenderReg:= regexp.MustCompile(gpuRender)
    for _, file := range devFiles{
		matchCard := gpuCardReg.MatchString(file.Name())
		matchRender := gpuRenderReg.MatchString(file.Name())
		if matchCard  {
			gpuCardExists = true
			continue
		}
		if matchRender {
			gpuRenderExists = true
			continue
		}
	}
    
	if gpuCardExists && gpuRenderExists{
		klog.Info("Card and device nodes present for dGPU")
		return true
	}

	klog.Error("Card or device nodes not present for dGPU")
		return false
}		