static int _ = [](){
    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);
    return 0;
}();


class Solution {
public:
    int maxProfit(vector<int>& prices, int fee) {
        int buy_state = INT_MIN;
        int sell_state = 0;

        for ( int price : prices ) {
            buy_state = max(buy_state, sell_state - price);
            sell_state = max(sell_state, buy_state + price - fee);
        }

        return sell_state;

    }
};