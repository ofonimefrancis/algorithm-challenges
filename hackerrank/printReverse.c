#include <stdio.h>

struct node {
    int data;
    struct node *next;
};

void printReverse(struct node *head) {
    struct node *temp = head;
    int size = getSize(head);

    int lList[size];
    int count = 0;
    while(temp != NULL) {
        lList[count++] = temp->data;
        temp = temp->next;
    }
    
    for(int i= size-1; i >= 0; i--) {
        printf("%d ", lList[i]);
    }
}