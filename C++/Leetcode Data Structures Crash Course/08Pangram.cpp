class Solution {
public:
    bool checkIfPangram(string sentence) {
        unordered_set<char> pangram(sentence.begin(), sentence.end());

        return pangram.size() == 26;
    }
}