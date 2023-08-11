/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    int minDepth(TreeNode* root) {
        if (root == nullptr) return 0;
        queue<TreeNode*> nodes;
        nodes.push(root);

        int depth = 0;
        while (!nodes.empty()) {
            int n = nodes.size();
            depth++;

            for (int i = 0; i < n; i++) {
                TreeNode* node = nodes.front();
                if (node->left != nullptr) nodes.push(node->left);
                if (node->right != nullptr) nodes.push(node->right);
                nodes.pop();
                if (node->left == nullptr && node->right == nullptr) return depth;
            }
        }

        return -1;
    }
};