/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
struct ListNode *detectCycle(struct ListNode *head) {
    struct ListNode* slow;
    struct ListNode* fast;
    struct ListNode* meet;
    
    if (head == NULL) {
        return NULL;
    }
    
    slow = fast = meet = head;
    
    while (1) {
        slow = slow->next;
        fast = fast->next;
        if (fast == NULL) {
            return NULL;
        }
        fast = fast->next;
        if (fast == NULL) {
            return NULL;
        }
        if (slow == fast) {
            while (slow != meet) {
                slow = slow->next;
                meet = meet->next;
            }
            return meet;
        }
    }
}
