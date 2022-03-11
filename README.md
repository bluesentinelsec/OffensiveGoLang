# Offensive GoLang

Offensive GoLang is is a collection of Go packages containing commonly used cyber adversary emulation functions.
Offensive GoLang accomplishes nothing by itself; rather, it is intended to support rapid red team tool development by providing common functions in a modular format.

## New to Offensive Go?

Check out my presentation at [SANS Pen Test HackFest Summit 2021](https://youtu.be/AGLunpPtOgM).

Slides can be found [here](https://github.com/bluesentinelsec/OffensiveGoLang/blob/master/Offensive%20GoLang%202.0%20-%20SANS%20Pen%20Test%20HackFest%202021.pdf)

## Requirements

* Go 1.12.7

## Quick start

Clone this repository using Go get: 
```
go get github.com/bluesentinelsec/OffensiveGoLang
```

Create a new Go source file, main.go, with the following code:

```
package main

import (
	"github.com/bluesentinelsec/OffensiveGoLang/pkg/windows/execution"
)

func main() {

	execution.RunPowerShell("notepad.exe")
}
```

Execute source file:
```
go run main.go
```

Offensive GoLang includes common functions such as downloading/uploading files, executing shellcode, establishing persistence, and more.

Take a look at the various sub packages to find interesting functions.

## Developers

Offensive GoLang follows the basic feature branch GIT flow. Create a feature branch off of master and when ready, submit a merge 
request. Make branch names and commits descriptive. Keep features concise and modular. Overly clever code is discouraged. 

## Acknowledgements
Special thanks to the developers of these great projects, whose works served as helpful references throughout the development of Offensive GoLang:

[PowerSploit](https://github.com/PowerShellMafia/PowerSploit)

[EGESPLOIT](https://github.com/EgeBalci/EGESPLOIT)

[GoBot](https://github.com/SaturnsVoid/GoBot2)

[Empire](https://github.com/EmpireProject/Empire)

[Merlin](https://github.com/Ne0nd0g/merlin)
