enum RES
{
    NONE,
    TRUE,
    FALSE
};

class Solution
{
    RES cache[21][31];

public:
    bool isMatch(string s, string p)
    {
        for (int i = 0; i <= s.length(); i++)
            for (int j = 0; j <= p.length(); j++)
                cache[i][j] = NONE;
        return dp(0, 0, s, p);
    }

private:
    bool dp(int i, int j, string s, string p)
    {
        if (cache[i][j] != NONE)
            return cache[i][j] == TRUE;

        bool ans;
        if (j == p.length())
            ans = i == s.length();
        else
        {
            bool a = i < s.length() && (s[i] == p[j] || p[j] == '.');
            if (j + 1 < p.length() && p[j + 1] == '*')
                ans = dp(i, j + 2, s, p) || (a && dp(i + 1, j, s, p));
            else
                ans = a && dp(i + 1, j + 1, s, p);
        }
        cache[i][j] = ans ? TRUE : FALSE;
        return ans;
    }
};
