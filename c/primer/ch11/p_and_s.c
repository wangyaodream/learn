#include <stdio.h>

int main() {
    // const char * msg = "Don't be a fool!";
    // const char * copy;
    char * msg = "Don't be a fool!";
    char * copy;

    copy = msg;
    printf("%s\n", copy);
    printf("msg = %s; &msg = %p; value = %p\n", msg, &msg, msg);
    printf("copy = %s; &copy = %p; copy = %p\n", copy, &copy, copy);

    return 0;
}
