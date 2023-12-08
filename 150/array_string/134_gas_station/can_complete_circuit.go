package gas_station

/*
	problem: https://leetcode.com/problems/gas-station

	complexity: O(n)
*/

func canCompleteCircuit(gas []int, cost []int) int {
	currGain, totalGain, answer := 0, 0, 0

	for i := 0; i < len(gas); i++ {
		totalGain += gas[i] - cost[i]
		currGain += gas[i] - cost[i]

		// If we meet a "valley", start over from the next station
		// with 0 initial gas.
		if currGain < 0 {
			answer = i + 1
			currGain = 0
		}
	}

	if totalGain >= 0 {
		return answer
	}
	return -1
}
