#include <stdio.h>
#include <string.h>

void fit(char *, unsigned int);

int main() {
    char msg [] = "Things should be as simple as possible,"
        "but not simpler.";
    puts(msg);
    fit(msg, 38);
    puts(msg);
    puts("Let's look at some more of the string.");
    puts(msg + 39);

    return 0;
}

void fit(char *string, unsigned int size){
    if (strlen(string) > size) {
        string[size] = '\0';
    }
}
