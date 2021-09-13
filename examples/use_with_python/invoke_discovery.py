#!/usr/bin/python3
import ctypes
import os

def main():
    # load libDiscovery
    libDiscovery_path = os.path.abspath("libDiscovery.so")
    print(f"[i] Loading Library: {libDiscovery_path}")
    libDiscovery = ctypes.cdll.LoadLibrary(libDiscovery_path)

    # set Discovery_GetCurrentUser return type to char *
    print("[i] Setting return type to 'char *' for function, Discovery_GetCurrentUser")
    libDiscovery.Discovery_GetCurrentUser.restype = ctypes.c_char_p

    # invoke Discovery_GetCurrentUser()
    print("[i] Invoking Discovery_GetCurrentUser()")
    user = libDiscovery.Discovery_GetCurrentUser()

    # print the result
    print("[+] Output for Discovery_GetCurrentUser():")
    print(user)
    print("[+] Congrats, you just called Go code from Python. :)")

if __name__ == "__main__":
    main()