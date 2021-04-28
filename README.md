# Offensive GoLang

Offensive GoLang is is a collection of Go packages that implement common cyber adversary tactics, techniques, and procedures (TTPs).

Offensive GoLang accomplishes nothing by itself; rather, it is intended to support rapid red team tool development by providing common functions in a modular format. It also provides clean examples of adversary TTPs that can be used as a reference to implement your own versions, possibly in other languages.

## Pardon our dust

As of May 2021, this project is undergoing a major refactor to incorporate a CI/CD pipeline, automated tests, and improved cross platform functionality. The goal with these changes is to increase the quality, stability, and accessibility of this project.

`Breaking changes are expected.`

Please see the [releases section](https://github.com/bluesentinelsec/OffensiveGoLang/releases) for the original repository (version 1.0).

## Requirements

[Install GoLang](https://golang.org/dl/)

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

## Want to contribute?

1. Fork the repo.
2. Make a descriptive branch name, like `implement_credential_dumping`.
3. Implement your changes.
4. Implement automated tests that demonstrate your feature works.
5. Submit a pull request with your changes to the [dev branch](https://github.com/bluesentinelsec/OffensiveGoLang/tree/dev).

## Acknowledgements

Special thanks to the developers of these great projects, whose works served as helpful references throughout the development of Offensive GoLang:

[PowerSploit](https://github.com/PowerShellMafia/PowerSploit)

[EGESPLOIT](https://github.com/EgeBalci/EGESPLOIT)

[GoBot](https://github.com/SaturnsVoid/GoBot2)

[Empire](https://github.com/EmpireProject/Empire)

[Merlin](https://github.com/Ne0nd0g/merlin)
