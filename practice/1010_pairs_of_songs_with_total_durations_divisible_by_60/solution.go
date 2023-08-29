package pairs_of_songs_with_total_durations_divisible_by_60

/*
	problem: https://leetcode.com/problems/pairs-of-songs-with-total-durations-divisible-by-60/

	what is range of value of each element
	what is minimum and maximum length of length of time

	pivot
	[30,   20,    150,   100,    40]


	2 for loop nested each other -> time complexity is O(n^2)


	better solution is:
							  [ 30,   20,    150,   100,     40 ]
remainder: 						30	   40      30     40      40
neededNum (60 - remainder):		30     20      30     20      20

remainderFreqMap[30] = 2
remainderFreqMap[40] = 3

neededNum:
key == num => cout++ , decrease remainderFreqMap

count == 1, remainderFreqMap[30]= 1
count == 2, remainderFreqMap[30]= 0 (delete)


return count = 2



*/

func numPairsDivisibleBy60(time []int) int {
	count := 0

	if len(time) == 1 {
		if time[0] == 60 {
			return 1
		}
		return 0
	}

	for i := 0; i < len(time)-1; i++ {
		for j := i + 1; j < len(time); j++ {
			if (time[i]+time[j])%60 == 0 {
				count++
			}
		}
	}

	return count
}
