package main

import (
	"C"
	"log"
	"unsafe"

	"github.com/fluent/fluent-bit-go/output"
)
import (
	"fmt"
	"time"
)

//export FLBPluginRegister
func FLBPluginRegister(ctx unsafe.Pointer) int {
	// Gets called only once when the plugin.so is loaded
	log.Print("[prettyslack] register")

	return output.FLBPluginRegister(ctx, "prettyslack", "Slack output pretty")
}

//export FLBPluginInit
func FLBPluginInit(ctx unsafe.Pointer) int {
	webhook := output.FLBPluginConfigKey(ctx, "webhook")
	log.Printf("[prettyslack] webhook = %q", webhook)
	// Set the context to point to any Go variable
	output.FLBPluginSetContext(ctx, webhook)
	return output.FLB_OK
}

func FLBPluginFlush(data unsafe.Pointer, length C.int, tag *C.char) int {
	log.Print("[prettyslack] Flush called for unknown instance")
	return output.FLB_OK
}

//export FLBPluginFlushCtx
func FLBPluginFlushCtx(ctx, data unsafe.Pointer, length C.int, tag *C.char) int {
	// Gets called with a batch of records to be written to an instance.
	webhook := output.FLBPluginGetContext(ctx).(string)
	log.Printf("[prettyslack] Flush called for webhook: %s", webhook)

	dec := output.NewDecoder(data, int(length))

	count := 0
	for {
		ret, ts, record := output.GetRecord(dec)
		if ret != 0 {
			break
		}

		var timestamp time.Time
		switch t := ts.(type) {
		case output.FLBTime:
			timestamp = ts.(output.FLBTime).Time
		case uint64:
			timestamp = time.Unix(int64(t), 0)
		default:
			fmt.Println("time provided invalid, defaulting to now.")
			timestamp = time.Now()
		}

		// Print record keys and values
		fmt.Printf("[%d] %s: [%s, {", count, C.GoString(tag), timestamp.String())

		for k, v := range record {
			fmt.Printf("\"%s\": %v, ", k, v)
		}
		fmt.Printf("}\n")
		count++
	}

	return output.FLB_OK
}

//export FLBPluginExit
func FLBPluginExit() int {
	return output.FLB_OK
}

//export FLBPluginExitCtx
func FLBPluginExitCtx(ctx unsafe.Pointer) int {
	return output.FLB_OK
}

func main() {
}
