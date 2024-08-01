package main

import (
	"fmt"
)

// Using Dynamic Programming, let's find a solution for the Largest Common Substring problem.
//
// This DP algorithm will use a 2d matrix to solve the problem.
// Where the values on the cell are what we are trying to optimize.
// Each cell is a subproblem.
//

// Problem: Find the largest commom substring between two words.
func LargestCommonSequence(s1, s2 string) string {
	n, m := len(s1), len(s2)
	if n == 0 || m == 0 {
		return ""
	}

	// Create a 2D slice to store lengths of longest common suffixes
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	longest := 0
	endIndex := 0

	// array starts at 1 to avoid dealing with negative indices
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1

				if dp[i][j] > longest {
					longest = dp[i][j]
					endIndex = i
				}
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return s1[endIndex-longest : endIndex]
}

func main() {
	word := LargestCommonSequence("fish", "fosh")

	fmt.Println(word)
}
