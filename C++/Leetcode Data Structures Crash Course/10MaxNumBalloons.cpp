#include <algorithm>

class Solution {
public:
    int maxNumberOfBalloons(string text) {
        int letters[256] = {};

        for (char c: text) {
            letters[c]++;
        }
        letters['l'] /= 2;
        letters['o'] /= 2;

        string balloon = "balloon";
        int maxWords = INT_MAX;
        for (auto c: balloon) {
            maxWords = min(maxWords, letters[c]);
        }

        return maxWords;
    }
};