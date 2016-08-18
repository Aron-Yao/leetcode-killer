/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
#define NO_MEANNING -1
#define ADD_TO_TAIL(tail, num) do {                                     \
    tail->next = (struct ListNode*) malloc(sizeof(struct ListNode));    \
    tail = tail->next;                                                  \
    tail->val = num, tail->next = NULL;                                 \
} while(0)

struct ListNode* addTwoNumbers(struct ListNode* l1, struct ListNode* l2) {
    int tmp;
    struct ListNode* res;
    struct ListNode* tail;
    struct ListNode* a;
    struct ListNode* b;
    
    res = (struct ListNode*) malloc(sizeof(struct ListNode));
    res->val = NO_MEANNING, res->next = NULL, tail = res;
    tmp = 0, a = l1, b = l2;
    
    while (a || b) {
        tmp /= 10;
        if (a) {
            tmp += a->val;
            a = a->next;
        }
        if (b) {
            tmp += b->val;
            b = b->next;
        }
        ADD_TO_TAIL(tail, tmp % 10);
    }
    
    if (tmp / 10 == 1) {
        ADD_TO_TAIL(tail, 1);
    }
    
    a = res;
    res = res->next;
    free(a);
    return res;
}
