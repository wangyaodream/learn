#include <stdio.h>
#include <inttypes.h>  // 支持可移植类型
int main(void){
    int32_t me32;  /* 是一个32位无符号整型变量  */

    me32 = 45933945;
    printf("First, assume int32_t is int:");
    printf("me32 = %d\n", me32);
    printf("Next, let's not make any assumptions.\n");
    printf("instead, use a \"macro\" from inttypes h:");
    printf("me32 = %" "d" "\n", me32);

    return 0;

}
