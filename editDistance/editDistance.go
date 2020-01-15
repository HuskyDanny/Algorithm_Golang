package editDistance


func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func editDistanceRecursive(s1 string, s2 string) int{

	var res int = 10000

	var recursiveHelper func(s1 string, s2 string, i int, j int, counts int) 
	recursiveHelper = func(s1 string, s2 string, i int, j int, counts int) {
		
		if i == len(s1)-1 || j == len(s2)-1 {
			
			res = min(res, counts + (len(s1) - i - 1) + (len(s2) - j - 1))
			return
		}

		if s1[i] == s2[j] {
			recursiveHelper(s1,s2,i+1,j+1, counts)
		}else {

			recursiveHelper(s1,s2,i+1, j, counts  + 1)
			recursiveHelper(s1,s2,i, j+1, counts  + 1)
			recursiveHelper(s1,s2,i+1, j+1, counts  + 1)

		}


	}

	recursiveHelper(s1,s2,0,0,0)
	return res
}


func editDistanceDP(s1 string, s2 string) int {
	dp := make([][]int, len(s1))
	for i := range dp {
		dp[i] = make([]int, len(s2))
	}

	for i := 0; i < len(s1); i++ {
		if (s2[0] == s1[i]) {
			dp[i][0] = i
		}else if (i > 0) {
			dp[i][0] = dp[i-1][0] + 1
		}else {
			dp[i][0] = 1
		}
	}
	
	for i := 0; i < len(s2); i++ {
		if (s2[i] == s1[0]) {
			dp[0][i] = i
		}else if (i > 0) {
			dp[0][i] = dp[0][i-1] + 1
		}else {
			dp[0][i] = 1
		}
	}

	for i := 1; i < len(s1); i++ {
		for j := 1; j < len(s2); j++ {
			
			if s1[i] == s2[j] {
				dp[i][j] = min(min(dp[i][j-1] + 1, dp[i-1][j] + 1), dp[i-1][j-1])
			}else {
				dp[i][j] = min(min(dp[i][j-1] + 1, dp[i-1][j] + 1) , dp[i-1][j-1] + 1)
			}

		}
	}

	return dp[len(s1)-1][len(s2)-1]
}