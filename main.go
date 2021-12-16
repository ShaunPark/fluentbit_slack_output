package main

import (
	"C"
	"log"
	"unsafe"

	"github.com/fluent/fluent-bit-go/output"
)

//export FLBPluginRegister
func FLBPluginRegister(ctx unsafe.Pointer) int {
	// Gets called only once when the plugin.so is loaded

	return output.FLBPluginRegister(ctx, "prettyslack", "Slack output pretty")
}

//export FLBPluginInit
func FLBPluginInit(ctx unsafe.Pointer) int {
	webhook := output.FLBPluginConfigKey(ctx, "webhook")
	log.Printf("[multiinstance] webhook = %q", webhook)
	// Set the context to point to any Go variable
	output.FLBPluginSetContext(ctx, webhook)
	return output.FLB_OK
}

//export FLBPluginFlushCtx
func FLBPluginFlushCtx(ctx, data unsafe.Pointer, length C.int, tag *C.char) int {
	// Gets called with a batch of records to be written to an instance.
	webhook := output.FLBPluginGetContext(ctx).(string)
	log.Printf("[multiinstance] Flush called for webhook: %s", webhook)
	return output.FLB_OK
}

//export FLBPluginExit
func FLBPluginExit() int {
	return output.FLB_OK
}

func main() {
}
