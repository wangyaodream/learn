#include <stdio.h>

int main(void){
    unsigned int un = 30000000000;
    short end = 200;
    long big = 65537;
    long long verybig = 12345678908642;

    printf("un = %u and not %d\n", un, un);

    return 0;
}
