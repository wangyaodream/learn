#include <stdio.h>
int units = 0;
void critic(void);

int main() {
    extern int units;  /* 可选重复声明 */
    
    printf("How many pounds to a frikin of butter?\n");
    scanf("%d", &units);
    while (units != 56)
        critic();
    printf("You must have looked it up!\n");

    return 0;
}

void critic(void) {
    printf("No luck, my friend. Try again.\n");
    scanf("%d", &units);
}
