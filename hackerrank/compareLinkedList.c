#include <stdio.h>
#include <stdlib.h>

struct node {
    int data;
    struct Node *next;
};

int CompareLists(struct node *headA, struct node* headB){
  int sizeA = getSize(headA);
  int sizeB = getSize(headB);

  struct node *currentA = headA;
  struct node *currentB = headB;

  if(sizeA != sizeB) return 0;
  while(currentA != NULL) {
      if(currentA->data != currentB->data) { return 0; }
      currentA = currentA->next;
      currentB = currentB->next;
  }
  return 1;
}