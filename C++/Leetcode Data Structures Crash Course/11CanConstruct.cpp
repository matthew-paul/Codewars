class Solution {
public:
    bool canConstruct(string ransomNote, string magazine) {
        if (magazine.size() < ransomNote.size()) return false;
        int magMap[26] = {};
        for (char c : magazine) {
            magMap[c - 'a']++;
        }
        for (char c : ransomNote) {
            if (--magMap[c - 'a'] < 0) return false;
        }

        return true;
    }
};