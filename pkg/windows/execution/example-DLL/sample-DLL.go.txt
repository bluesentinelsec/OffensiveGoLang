// Compile DLL on Windows:
// go build -o sample.dll -buildmode=c-shared sample.go

// Execute DLL:
// C:> rundll32.exe sample.dll,PopCalc

// Compile DLL on Linux:
// GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -buildmode=c-shared -o calc.dll calc.go

package main

import "C"
import (
	"fmt"
	"os/exec"
)

//export PopCalc
func PopCalc() {
	fmt.Println("Spawning calculator")
	cmd := exec.Command("cmd.exe", "/C", "C:\\Windows\\System32\\calc.exe")
	cmd.Run()
}

func main() {
	// leave blank
}
