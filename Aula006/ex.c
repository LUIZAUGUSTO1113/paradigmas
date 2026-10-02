#include <stdio.h>

int main() {
    union { int i; float f; } u;
    u.f = 1.0f;
    printf("%d\n", u.i);
    return 0;
}