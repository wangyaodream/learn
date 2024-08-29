/* print1.c - 演示printf()的一些特性 */
#include <stdio.h>

int main(void)
{
    int ten = 10;
    int two = 2;

    printf("Doing it right:");
    printf("%d minus %d is %d\n", ten, 2, ten - two);
    printf("Doning it wrong:");
    printf("%d minus %d is %d\n", ten);  // 缺少参数

    return 0;
}
