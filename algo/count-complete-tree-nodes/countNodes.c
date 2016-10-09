/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
int getDepth (struct TreeNode* node)
{
    int cnt = 0;

    while (node != NULL) 
    {
        ++cnt;
        node = node->left; 
    }

    return cnt;
}

int countNodes(struct TreeNode* root)
{
    static unsigned int full[] = {1, 3, 7, 15, 31, 63, 127, 255, 511, 1023, 2047, 4095, 8191, 16383, 32767, 65535, 131071, 262143, 524287, 1048575};

    if (root == NULL) {
        return 0;
    }

    struct TreeNode* p = root;
    int ldepth = 0, rdepth = 0, depth = 0;

    depth = getDepth (p);
    ldepth = depth - 1;
    rdepth = 0;
    int cnt = full[depth - 2];

    while (true) 
    {
        rdepth = getDepth (p->right); 

        if (rdepth == 0) {
            ++cnt;
            break;
        }

        if (ldepth == rdepth) {
            cnt += full[ldepth - 2] + 1; 
            p = p->right;
        }
        else {
            p = p->left;
        }

        --ldepth;
    }

    return cnt;
}

