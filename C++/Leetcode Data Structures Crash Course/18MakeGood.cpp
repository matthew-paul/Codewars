class Solution {
public:
    string makeGood(string s) {
        string ans = "";
        int n = s.size();
        if (n == 1) return s;

        for (char c : s) {
            if (ans.empty()) {
                ans.push_back(c);
            } else {
                if (abs(c - ans.back()) == 32) {
                    ans.pop_back();
                } else {
                    ans.push_back(c);
                }
            }
        }

        return ans;
    }
};