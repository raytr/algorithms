package football_scores

import "sort"

/*
use 2 pointer to solve
sort A and B

aPointer at head of A
bPointer at head of B

for i:=0; i<B.len; i++
	while aPointer.val <= bPointer.val
		aPointer++
		count++
	res.push(count)

return res
*/

type teamIndex struct {
	score int32
	index int
}

func counts(teamA []int32, teamB []int32) []int32 {
	B := make([]teamIndex, 0, len(teamB))
	for i, score := range teamB {
		B = append(B, teamIndex{
			score: score,
			index: i,
		})
	}

	//sort teamA and teamB
	sort.Slice(teamA, func(i, j int) bool {
		return teamA[i] < teamA[j]
	})
	sort.Slice(B, func(i, j int) bool {
		return B[i].score < B[j].score
	})

	aPointer, count := 0, int32(0)
	ori := make(map[int]int32)
	for bPointer := 0; bPointer < len(B); bPointer++ {
		for aPointer < len(teamA) && teamA[aPointer] <= B[bPointer].score {
			aPointer++
			count++
		}
		ori[B[bPointer].index] = count
	}

	//convert to res
	res := make([]int32, 0, len(teamB))
	for i := 0; i < len(B); i++ {
		res = append(res, ori[i])
	}
	return res
}
