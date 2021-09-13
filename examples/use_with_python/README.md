# Use Offensive GoLang with Python

This example demonstrates how to use Offensive GoLang code with Python.

## Quick Start (MacOS)

1. Run invoke_discovery.py

```bash
python3 invoke_discovery.py
```

Output:

```
[i] Loading Library: OffensiveGoLang/examples/use_with_python/libDiscovery.so
[i] Setting return type to 'char *' for function, Discovery_GetCurrentUser
[i] Invoking Discovery_GetCurrentUser()
[+] Output for Discovery_GetCurrentUser():
b'bluesentinel'
[+] Congrats, you just called Go code from Python.
```

Note that the provided libaries were compiled on MacOS.