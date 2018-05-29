#include <stdio.h>

struct SinglyLinkedListNode {
    int data;
    struct SinglyLinkedListNode *next;
};

typedef SinglyLinkedListNode  SinglyLinkedListNode;


SinglyLinkedListNode* insertNodeAtPosition(SinglyLinkedListNode* head, int data, int position) { 
    struct SinglyLinkedListNode *current;
    struct SinglyLinkedListNode *previous;
    struct SinglyLinkedListNode *newNode;

    newNode = malloc(sizeof(struct SinglyLinkedListNode));
    previous = malloc(sizeof(struct SinglyLinkedListNode));

    newNode->data = data;
    newNode->next = NULL;
    
    if(head == NULL) return newNode;
    
    current = head;
    int tempPosition = 0;

    while(current->next != NULL && tempPosition != position) {
        previous = current;    
        current = current->next;
        tempPosition++;
    }
    previous->next = newNode;
    newNode->next = current;

    return head;   
}