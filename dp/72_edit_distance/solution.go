package edit_distance

import (
	"fmt"
	"sort"
)

/*
	problem: https://leetcode.com/problems/edit-distance/


        stage:
            word1[1...i] (i is any indexing in word1)
            word2[1...j] (j is any indexing in word2)

            for example: HORSE & ROS
                we have word1[1...i] = HOR
                        word2[1...j] = RO
                D[i][j] = ?: the edit between HOR and RO
                D[i - 1][j] = 1: the edit between HO and RO
                D[i][j-1] = 2: the edit between HOR and R
                D[i-1][j-1] = 2: the edit between HO and R


        relation between sub-problems and problem:
            D[i][j]=1+min(D[i−1][j], D[i][j−1], D[i−1][j−1]); since R != 0


        base case:
            if word1[0] != word2[0] => D[0][0] = 1, ortherwise = 0
            if word1[1] != word2[0] => D[1][0] = 1, ortherwise = 0
            if word1[0] != word2[1] => D[0][1] = 1, ortherwise = 0


*/

var D [][]int

func minDistance(word1 string, word2 string) int {
	n, m := len(word1), len(word2)
	D = make([][]int, n)
	for i := 0; i < n; i++ {
		D[i] = make([]int, m, m)
	}

	// set base case
	if word1[0] == word2[0] {
		D[0][0] = 0
	} else {
		D[0][0] = 1
	}
	if word1[1] == word2[0] {
		D[1][0] = 0
	} else {
		D[1][0] = 1
	}
	if word1[0] == word2[1] {
		D[0][1] = 0
	} else {
		D[0][1] = 1
	}

	return call(word1, word2, n-1, m-1)
}

func call(word1, word2 string, i, j int) int {
	ax := D
	fmt.Println(ax)
	//base cases
	if i == 0 && j == 0 {
		return D[0][0]
	}
	if i == 0 && j == 1 {
		return D[0][1]
	}
	if i == 1 && j == 0 {
		return D[1][0]
	}

	if D[i-1][j-1] == 0 {
		D[i-1][j-1] = call(word1, word2, i-1, j-1)
	}
	if D[i-1][j] == 0 {
		D[i-1][j] = call(word1, word2, i-1, j)
	}
	if D[i][j-1] == 0 {
		D[i][j-1] = call(word1, word2, i, j-1)
	}

	return 1 + Min(D[i-1][j-1], D[i][j-1], D[i-1][j])
}

func Min(a, b, c int) int {
	arr := []int{a, b, c}
	sort.Ints(arr)
	return arr[0]
}
