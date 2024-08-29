#include <stdio.h>
#define NAME "CHONGQING CQU"
#define ADDRESS "101 Megabuck plaza"
#define PLACE "Megapilis, CA 949660"
#define WIDTH 40

void starbar(void);

int main(){
    starbar();
    printf("%s\n", NAME);
    printf("%s\n", ADDRESS);
    printf("%s\n",PLACE);
    starbar();

    return 0;
}

void starbar(){
    int count;

    for (count = 1; count <= WIDTH; count++)
        putchar('*');
    putchar('\n');
}

