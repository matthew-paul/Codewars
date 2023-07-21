class Solution {
public:
    vector<vector<int>> findWinners(vector<vector<int>>& matches) {
        map<int, int> lost;

        for (auto m : matches) {
            lost[m[0]];
            lost[m[1]]++;
        }

        vector<vector<int>> ans(2);
        for (auto[k, l] : lost) {
            if (l < 2) ans[l].push_back(k);
        }

        return ans;
    }
};