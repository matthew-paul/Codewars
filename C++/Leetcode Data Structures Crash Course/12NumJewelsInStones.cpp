class Solution {
public:
    int numJewelsInStones(string jewels, string stones) {
        int jewelsMap[256] = {};

        for (char c : jewels) {
            jewelsMap[c]++;
        }

        int ans = 0;
        for (char c : stones) {
            if (jewelsMap[c]) ans++;
        }

        return ans;
    }
};