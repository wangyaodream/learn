#include <stdio.h>
#define PI 3.1415926

int main(void){
    float area, circum, radius;

    printf("What is the radius of you pizza?\n");
    scanf("%f", &radius);
    area = PI * radius * radius;

    circum = 2.0 * PI * radius;
    printf("Your basic pizza parameter are as follows.\n");
    printf("circumference = %1.2f, area = %1.2f\n", circum, area);

    return 0;
}
