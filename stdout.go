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
	rainbow("Guin is awesome! \n", 0)
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
	rainbow("Guin is awesome! \n", 0)
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
		for k, v := range record {
			logLine += fmt.Sprint(", {\"", k, "\": ", v)
		}
		logLine += "}]"
		rainbow(logLine, count)
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

func rainbow(s string, lineCount int) {
	for i := 0; i < len(s); i++ {

		subindex := (i + lineCount) % 30 // will give the index of each 30 char substring, and displace the start color by 1 character

		switch {
		case subindex >= 0 && subindex < 5:
			red(string(s[i]))
		case subindex >= 5 && subindex < 10:
			yellow(string(s[i]))
		case subindex >= 10 && subindex < 15:
			green(string(s[i]))
		case subindex >= 15 && subindex < 20:
			blue(string(s[i]))
		case subindex >= 20 && subindex < 25:
			magenta(string(s[i]))
		case subindex >= 25 && subindex < 30:
			cyan(string(s[i]))
		default:
			fmt.Printf(string(s[i]))
		}
	}
	fmt.Printf("\n")
}

func red(s string) {
	fmt.Printf("%s", strings.TrimSuffix(color.RedString(s), " "))
}

func yellow(s string) {
	fmt.Printf("%s", strings.TrimSuffix(color.YellowString(s), " "))
}

func blue(s string) {
	fmt.Printf("%s", strings.TrimSuffix(color.BlueString(s), " "))
}

func green(s string) {
	fmt.Printf("%s", strings.TrimSuffix(color.GreenString(s), " "))
}

func magenta(s string) {
	fmt.Printf("%s", strings.TrimSuffix(color.MagentaString(s), " "))
}

func cyan(s string) {
	fmt.Printf("%s", strings.TrimSuffix(color.CyanString(s), " "))
}
