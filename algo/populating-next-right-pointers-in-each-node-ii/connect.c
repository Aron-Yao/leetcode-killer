/**
 * Definition for binary tree with next pointer.
 * struct TrqeeLinkNode {
 *  int val;
 *  struct TreeLinkNode *left, *right, *next;
 * };
 *
 */
struct TreeNode* findNext(struct TreeLinkNode* root) {
    if (root == NULL) {
        return NULL;
    }
    if (root->left) {
        return root->left;
    }
    if (root->right) {
        return root->right;
    }
    
    return findNext(root->next);
}

void connect(struct TreeLinkNode *root) {
    struct TreeLinkNode* cur;
    struct TreeLinkNode* pre;
    struct TreeLinkNode* next;
    struct TreeLinkNode* tmp;
    
    if (root == NULL) {
        return;
    } 
    
    pre = root;
    cur = NULL;
    
    while (1) {
        cur = pre;
        pre = findNext(pre);
        if (pre == NULL) {
            return;
        }
        while (cur) {
            if (cur->left && cur->right) {
                cur->left->next = cur->right;
            } 
            tmp = (cur->right) ? (cur->right) : (cur->left);
            if (tmp) {
                tmp->next = findNext(cur->next);
            }
            cur = cur->next;
        }
    }
}
