#include <stdio.h>
#include <stdlib.h>

int main(){
    int n;
    scanf("%d", &n);

    int arr[n][n];

    int diag = 0;
    for(int i=0; i<n; i++){
        for(int j=0; j<n; j++){
            scanf("%d", &(arr[i][j]));
        }
    }
    for(int i=0; i<n; i++){
        diag += arr[i][i];
    }


    int left_diag = 0;
    for(int i=0; i<n; i++){
        left_diag += arr[i][n-i-1];
    }
    printf("%d\n", abs(diag - left_diag));

    return 0;
}
