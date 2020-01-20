// C(GCC 5.4.1)

#include <stdio.h>

int main() {
    int a, b, i;
    scanf("%d %d", &a, &b);

    if (a<b){
        for(i=0; i<b; i++)
            printf("%d", a);
    } else {
        for(i=0; i<a; i++)
            printf("%d", b);
    }
    return 0;
}
