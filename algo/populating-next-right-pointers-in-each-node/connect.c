/**
 * Definition for binary tree with next pointer.
 * struct TrqeeLinkNode {
 *  int val;
 *  struct TreeLinkNode *left, *right, *next;
 * };
 *
 */
void connect(struct TreeLinkNode *root) {
    struct TreeLinkNode* pre;
    struct TreeLinkNode* cur;
    
    if (root == NULL) {
        return;
    }
    
    pre = root;
    cur = NULL;
    
    while (pre->left) {
        cur = pre;
        pre = pre->left;
        while (cur) {
            cur->left->next = cur->right;
            if (cur->next) {
                cur->right->next = cur->next->left;
            }
            cur = cur->next;
        }
    }
}
