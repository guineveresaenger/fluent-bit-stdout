package main

import (
	"context"
	"testing"
	"unsafe"

	"github.com/fluent/fluent-bit-go/output"
)

func TestFLBPluginRegister(t *testing.T) {
	bk := context.Background()
	var ctx = unsafe.Pointer(&bk)

	result := FLBPluginRegister(ctx)
	if result != 0 {
		t.Error("Failed to register plugin, expected", 0, "but found", result)
	}
}

func TestFLBPluginInitFails(t *testing.T) {
	bk := context.Background()
	var ctx = unsafe.Pointer(&bk)

	result := FLBPluginInit(ctx)
	if result != output.FLB_OK {
		t.Error("Expected to successfully connect, but failed instead", result)
	}
}

func TestFLBPluginExit(t *testing.T) {

	result := FLBPluginExit()
	if result != output.FLB_OK {
		t.Error("Failed to exit plugin, expected", 0, "but found", result)
	}
}
