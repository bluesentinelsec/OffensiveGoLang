// This program demonstrates how to write a C binding
// for Offensive GoLang code.

package main

import "C"

import (
	"github.com/bluesentinelsec/OffensiveGoLang/discovery"
)

//export Discovery_GetCurrentUser
func Discovery_GetCurrentUser() *C.char {
	user, err := discovery.GetCurrentUser()
	if err != nil {
		e := err.Error()
		return C.CString(e)
	}
	return C.CString(user)
}

func main() {}
