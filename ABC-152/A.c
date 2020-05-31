// C(GCC 5.4.1)

#include <stdio.h>
 
int main(void) {
    int n, m;
 
    scanf("%d %d", &n, &m);
    if (n == m && (1 <= n <= 100) && (0 <= m <= n)) {
        printf("Yes");
    } else {
        printf("No");
    }
    return 0;
}

// Better answer

#include <stdio.h>
 
int main(void) {
    int n, m;
    scanf("%d %d", &n, &m);
    if (m<n){
        printf("No");
    } else if(n==m){
        printf("Yes");
    }
    return 0;
}
