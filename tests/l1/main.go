// Copyright (c) 2022 - 2023 Intel Corporation
// SPDX-License-Identifier: Apache-2.0 

package main

import (
        "l1/sgx"
        "l1/dgpu"
)

func main() {
        sgx.DeviceCheck()
        dgpu.NodeCheck()
}