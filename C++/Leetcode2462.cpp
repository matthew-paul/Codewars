class Solution {
public:
    long long totalCost(vector<int>& costs, int k, int candidates) {
        long long ans = 0;
        int i = 0;
        int j = costs.size() - 1;
        priority_queue<int, vector<int>, greater<int>> pq1;
        priority_queue<int, vector<int>, greater<int>> pq2;
        while(k--) {
            while(pq1.size() < candidates && i <= j)  {
                pq1.push(costs[i++]);
            }
            while(pq2.size() < candidates && j >= i) {
                pq2.push(costs[j--]);
            }

            int c1 = pq1.size() > 0 ? pq1.top() : INT_MAX;
            int c2 = pq2.size() > 0 ? pq2.top() : INT_MAX;

            if (c1 <= c2) {
                ans += c1;
                pq1.pop();
            } else {
                ans += c2;
                pq2.pop();
            }
        }
        return ans;
    }
};