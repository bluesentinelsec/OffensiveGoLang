# Use Offensive GoLang with C

This example demonstrates how to use Offensive GoLang code with C.

## Quick Start (MacOS)

1. Install clang

```bash
command xcode-select --install
```

2. Compile program with the provided libraries

```bash
clang -o exampleProgram exampleProgram.c -L ./ -l Discovery
```

Note that the -L and -l options tell the compiler to link to libDiscovery.a in the current directory.

3. Test program

```
./exampleProgram
```

Output:

```
[i] Calling an OffensiveGoLang function: DiscoveryGetCurrentUser
[+] Username: michaellong
[!] Warning: you are responsible for freeing memory when Go functions return values to C code.
[i] See 'exampleProgram.c' for more information.
```

Note that the provided libaries were compiled on MacOS.