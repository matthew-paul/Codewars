#include <algorithm>

class Solution {
public:
    int longestOnes(vector<int>& nums, int k) {
        int left = 0, right = 0, ans = 0, ones = 0;

        while (right < nums.size()) {
            if (nums[right]) ones++;

            while ((right - left+1)-ones > k) {
                if (nums[left]) ones--;
                left++;
            }

            ans = max(ans, right-left+1);
            right++;
        }

        return ans;
    }
};