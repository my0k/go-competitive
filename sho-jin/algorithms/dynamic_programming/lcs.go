package dynamic_programming

// LCS returns one of the Longest Common Subsequence of S and T.
// LCS only accepts len(S) and len(T) <= 5000.
func LCS(S, T []rune) []rune {
	dp := [5001][5001]int{}

	for i := 0; i < len(S); i++ {
		for j := 0; j < len(T); j++ {
			if S[i] == T[j] {
				ChMax(&dp[i+1][j+1], dp[i][j]+1)
			}
			ChMax(&dp[i+1][j+1], dp[i+1][j])
			ChMax(&dp[i+1][j+1], dp[i][j+1])
		}
	}

	revRes := make([]rune, 0, dp[len(S)][len(T)])
	si, ti := len(S), len(T)
	for si > 0 && ti > 0 {
		if dp[si][ti] == dp[si-1][ti] {
			si--
		} else if dp[si][ti] == dp[si][ti-1] {
			ti--
		} else {
			revRes = append(revRes, S[si-1])
			si--
			ti--
		}
	}

	res := make([]rune, len(revRes))
	for i := len(revRes) - 1; i >= 0; i-- {
		res[len(revRes)-1-i] = revRes[i]
	}

	return res
}
