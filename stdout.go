package main

import (
	"strings"

	"github.com/fatih/color"
	"github.com/fluent/fluent-bit-go/output"
)
import (
	"C"
	"fmt"
	"unsafe"
)

//export FLBPluginRegister
func FLBPluginRegister(ctx unsafe.Pointer) int {
	return output.FLBPluginRegister(ctx, "guin_stdout", "Stdout GO!")
}

//export FLBPluginInit
// (fluentbit will call this)
// ctx (context) pointer to fluentbit context (state/ c code)
func FLBPluginInit(ctx unsafe.Pointer) int {
	// Example to retrieve an optional configuration parameter
	param := output.FLBPluginConfigKey(ctx, "param")
	fmt.Printf("[flb-go] plugin parameter = '%s'\n", param)
	rainbow("Guin is awesome! \n")
	return output.FLB_OK
}

//export FLBPluginFlush
func FLBPluginFlush(data unsafe.Pointer, length C.int, tag *C.char) int {
	var count int
	var ret int
	var ts interface{}
	var record map[interface{}]interface{}

	// Create Fluent Bit decoder
	dec := output.NewDecoder(data, int(length))
	rainbow("Guin is awesome! \n")
	// Iterate Records
	count = 0
	for {
		// Extract Record
		ret, ts, record = output.GetRecord(dec)
		if ret != 0 {
			break
		}

		// Print record keys and values
		timestamp := ts
		logLine := fmt.Sprint("[", count, "] ", C.GoString(tag), ": [", timestamp)
		// fmt.Printf("[%d] %s: [%d, {", count, C.GoString(tag), timestamp)
		for k, v := range record {
			logLine += fmt.Sprint(", {\"", k, "\": ", v)
		}
		logLine += "}]"
		rainbow(logLine)
		fmt.Printf("\n")
		count++
	}

	// Return options:
	//
	// output.FLB_OK    = data have been processed.
	// output.FLB_ERROR = unrecoverable error, do not try this again.
	// output.FLB_RETRY = retry to flush later.
	return output.FLB_OK
}

//export FLBPluginExit
func FLBPluginExit() int {
	return output.FLB_OK
}

func main() {
}

func rainbow(s string) {
	for i := 0; i < len(s); i += 30 {

		// fmt.Printf("%d", i)
		if i+5 > len(s) {
			fmt.Printf("%s", strings.TrimSuffix(color.RedString(s[i:len(s)]), " "))
			break
		} else {
			fmt.Printf("%s", strings.TrimSuffix(color.RedString(s[i:(i+5)]), " "))
		}
		if i+10 > len(s) {
			fmt.Printf("%s", strings.TrimSuffix(color.YellowString(s[i+5:len(s)]), " "))
			break
		} else {
			fmt.Printf("%s", strings.TrimSuffix(color.YellowString(s[(i+5):(i+10)]), " "))
		}
		if i+15 > len(s) {
			fmt.Printf("%s", strings.TrimSuffix(color.GreenString(s[i+10:len(s)]), " "))
			break
		} else {
			fmt.Printf("%s", strings.TrimSuffix(color.GreenString(s[(i+10):(i+15)]), " "))
		}
		if i+20 > len(s) {
			fmt.Printf("%s", strings.TrimSuffix(color.BlueString(s[i+15:len(s)]), " "))
			break
		} else {
			fmt.Printf("%s", strings.TrimSuffix(color.BlueString(s[(i+15):(i+20)]), " "))
		}
		if i+25 > len(s) {
			fmt.Printf("%s", strings.TrimSuffix(color.MagentaString(s[i+20:len(s)]), " "))
			break
		} else {
			fmt.Printf("%s", strings.TrimSuffix(color.MagentaString(s[(i+20):(i+25)]), " "))
		}
		if i+30 > len(s) {
			fmt.Printf("%s", strings.TrimSuffix(color.CyanString(s[i+25:len(s)]), " "))
			break
		} else {
			fmt.Printf("%s", strings.TrimSuffix(color.CyanString(s[(i+25):(i+30)]), " "))
		}
	}
	fmt.Printf("\n")
}
