class Solution {
public:
    int minStartValue(vector<int>& nums) {
        int min = INT_MAX;
        int sum = 0;
        for (auto n : nums) {
            sum += n;
            if (sum < min) min = sum;
        }

        if (min > 0) return 1;
        else return -min+1;
    }
};