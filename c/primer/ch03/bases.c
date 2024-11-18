#include <stdio.h>

int main(void) {
    int x = 100;

    printf("dec = %d; octal = %o; hex = %x\n", x, x, x);
    // 显示各进制值的前缀使用#
    printf("dec = %d; octal = %#o; hex = %#x\n", x, x, x);

    return 0;
}
