/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
bool hasCycle(struct ListNode *head) {
    struct ListNode* slow, *fast;
    
    if (head == NULL) {
        return false;
    }
    
    slow = fast = head;
    
    while (1) {
        slow = slow->next;
        fast = fast->next;
        if (fast == NULL) {
            return false;
        }
        fast = fast->next;
        if (fast == NULL) {
            return false;
        }
        if (fast == slow) {
            return true;
        }
    }
    return false;
}
