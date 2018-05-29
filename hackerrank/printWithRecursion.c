#include <stdio.h>

struct node {
    int data;
    struct node *next;
};

void PrintWithRecursion(struct node *head){
    if(head == NULL) return;
    struct node *temp = head;
    printf("%d ", temp->data);
    PrintWithRecursion(temp->next);
}