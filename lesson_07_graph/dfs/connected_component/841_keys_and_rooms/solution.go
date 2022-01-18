//package _41_keys_and_rooms
package main

import "fmt"

func main() {
	fmt.Println(canVisitAllRooms([][]int{{1}, {2}, {3}, {}}))
	fmt.Println(canVisitAllRooms([][]int{{1, 3}, {1, 0, 3}, {2}, {0}}))
}

/*
	problem: https://leetcode.com/problems/keys-and-rooms/

	first of all, convert to adjList
	loop all room and we have map
	roomMap:
	- roomMap[0] : [1]
	- roomMap[1] : [2]
	- roomMap[2] : [3]
	- roomMap[3] : []

	use dfs, if count == 1 => true , otherwise => false

	complexity: O(n)
*/
var count = 0
var visited = make(map[int]bool)
var roomMap = make(map[int][]int)

func canVisitAllRooms(rooms [][]int) bool {

	for i := 0; i < len(rooms); i++ {
		roomMap[i] = make([]int, 0, len(rooms))
		for j := 0; j < len(rooms[i]); j++ {
			roomMap[i] = append(roomMap[i], rooms[i][j])
		}
	}

	for i := 0; i < len(rooms); i++ {
		if !visited[i] {
			count++
		}
		dfs(i)
	}
	if count > 1 {
		return false
	}
	return true
}

func dfs(currentRoom int) {
	for _, r := range roomMap[currentRoom] {
		if !visited[r] {
			visited[r] = true
			dfs(r)
		}
	}
}
