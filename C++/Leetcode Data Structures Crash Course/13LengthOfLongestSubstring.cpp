class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int left=0, maxLen=0;
        unordered_set<char> uniqueLetters;

        for (int right=0; right < s.size(); right++) {
            if (!uniqueLetters.count(s[right])) {
                uniqueLetters.insert(s[right]);
                maxLen = max(maxLen, right-left+1);
            } else {
                while (s[left] != s[right])
                    uniqueLetters.erase(s[left++]);

                uniqueLetters.erase(s[left++]);
                uniqueLetters.insert(s[right]);
            }
        }

        return maxLen;
    }
};