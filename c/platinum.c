#include<stdio.h>

int main(void) 
{
    float weight;
    float value;

    printf("Are you worth your weight in platinum?\n");
    printf("Let's check it out.\n");
    printf("Please enter your weight in pounds.");

    /* 获取用户的输入 */
    scanf("%f", &weight);

    value = 1700.0 * weight * 14.5833;
    printf("Your weight in platinum is worth $%.2f.\n", value);

    printf("You are easily worth that! if platinum prices drop.\n");

    printf("eat more to maintain your value.\n");
    return 0;
}
