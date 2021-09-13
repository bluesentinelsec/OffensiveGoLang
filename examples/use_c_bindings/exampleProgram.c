#include <stdio.h>
#include <stdlib.h>

#include "libDiscovery.h"

int main(int argc, char *argv[])
{
    printf("[i] Calling an OffensiveGoLang function: DiscoveryGetCurrentUser\n");
    
    char *username = Discovery_GetCurrentUser();
    
    printf("[+] Username: %s\n", username);
    
    printf("[!] Warning: you are responsible for freeing memory when Go functions return values to C code.\n");

    printf("[i] See 'exampleProgram.c' for more information.\n");
    
    free(username);

    return 0;
}
