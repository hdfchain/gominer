// Copyright (c) 2016 The Decred developers.

// +build cuda,!opencl

package main

/*
#cgo !windows LDFLAGS: -L/opt/cuda/lib64 -L/opt/cuda/lib -lcuda -lcudart -lstdc++ obj/hdfchain.a
#cgo windows LDFLAGS: -Lobj -lhdfchain -Lnvidia/CUDA/v7.0/lib/x64 -lcuda -lcudart -Lnvidia/NVSMI -lnvml
*/
import "C"
