#include <stdio.h>

int main() {
    int age;
    float assets;
    char pet[30];

    printf("Enter your age, assets, and favorite pet.\n");
    scanf("%d %f", &age, &assets);
    scanf("%s", pet);  // 不需要&符号，因为pet是一个字符串
    printf("%d $%.2f %s\n", age, assets, pet);

    return 0;
}
