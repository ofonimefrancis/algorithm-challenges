#include <stdio.h>
#include <string.h>
#include <math.h>
#include <stdlib.h>

static int posCount = 0;
static int negCount = 0;
static int zerCount = 0;


void countLogic(int element){
    if(element < 0)
        negCount ++;
    else if(element > 0)
        posCount ++;
    else if(element == 0)
        zerCount ++;
}

int main(){
    int n;
    
    scanf("%d", &n);
    
    int arr[n];
    
    for(int i=0; i<n; i++){
        scanf("%d", &(arr[i]));
        countLogic(arr[i]);
        
    }
    
    printf("%.3f\n", (double)posCount/n);
    printf("%.3f\n", (double)negCount/n);
    printf("%.3f\n", (double)zerCount/n);
    
    return 0;
}