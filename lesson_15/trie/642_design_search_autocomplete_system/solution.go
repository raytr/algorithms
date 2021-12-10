//package design_search_autocomplete_system
package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
	problem: https://leetcode.com/problems/design-search-autocomplete-system/

	func Input()
		if c == '#' {
			cur.Insert(curInputed, 1)
			curInputed = ""
			return res
		} else {
			curInputed += string(c)
		}

		//search in trie
		for char := range curInputed


*/

func main() {
	//autocompleteSystem := Constructor([]string{"i love you", "island", "iroman", "i love leetcode"}, []int{5, 3, 2, 2})
	//fmt.Println(autocompleteSystem.Input('i'))
	//fmt.Println(autocompleteSystem.Input(' '))
	//fmt.Println(autocompleteSystem.Input('a'))
	//fmt.Println(autocompleteSystem.Input('#'))

	autocompleteSystem := Constructor([]string{"abc", "abbc", "a"}, []int{3, 3, 3})
	fmt.Println(autocompleteSystem.Input('b'))
	fmt.Println(autocompleteSystem.Input('c'))
	fmt.Println(autocompleteSystem.Input('#'))
	fmt.Println(autocompleteSystem.Input('b'))
	fmt.Println(autocompleteSystem.Input('c'))
	fmt.Println(autocompleteSystem.Input('#'))
	fmt.Println(autocompleteSystem.Input('a'))
	fmt.Println(autocompleteSystem.Input('b'))
	fmt.Println(autocompleteSystem.Input('c'))
	fmt.Println(autocompleteSystem.Input('#'))
	fmt.Println(autocompleteSystem.Input('a'))
	fmt.Println(autocompleteSystem.Input('b'))
	fmt.Println(autocompleteSystem.Input('c'))
	fmt.Println(autocompleteSystem.Input('#'))

}

//["AutocompleteSystem","input","input","input","input","input","input","input","input","input","input","input","input","input","input"]
//[[["abc","abbc","a"],[3,3,3]],["b"],["c"],["#"],["b"],["c"],["#"],["a"],["b"],["c"],["#"],["a"],["b"],["c"],["#"]]

type AutocompleteSystem struct {
	children    []*AutocompleteSystem
	end         bool
	wordFreqMap map[string]int
}
type WordFreq struct {
	word string
	freq int
}

var curInputed string
var isExist bool // detect if the previous letter doesn't exist => return [] too
func Constructor(sentences []string, times []int) AutocompleteSystem {
	/*
	   build trie, with each node, we add wordFreqMap follow sentences and time
	*/
	curInputed = ""
	isExist = true
	node := NewNode()
	for i := 0; i < len(sentences); i++ {
		node.Insert(sentences[i], times[i])
	}
	return *node
}

func (this *AutocompleteSystem) Input(c byte) []string {
	res := make([]string, 0, 3)

	//store c
	cur := this
	if c == '#' {
		cur.Insert(curInputed, 1)
		curInputed = ""
		isExist = true
		return res
	} else {
		curInputed += string(c)
		if !isExist {
			return []string{}
		}
	}

	//find the node
	for _, char := range curInputed {
		idx := getIndex(char)

		// find in trie
		if cur.children[idx] == nil {
			isExist = false
			return []string{}
		}

		if char == rune(c) {
			//add to res
			minHeap := initHeap()
			for k, v := range cur.children[idx].wordFreqMap {
				heap.Push(minHeap, WordFreq{k, v})
			}

			for i := 0; i < 3; i++ {
				if minHeap.Len() == 0 {
					break
				}
				a := heap.Pop(minHeap).(WordFreq).word
				res = append(res, a)
			}
			break
		}
		cur = cur.children[idx]
	}

	return res
}

func getIndex(c int32) int {
	if c == 32 {
		return 26
	} else {
		return int(c - 'a')
	}
}

func NewNode() *AutocompleteSystem {
	return &AutocompleteSystem{
		children:    make([]*AutocompleteSystem, 27, 27),
		wordFreqMap: make(map[string]int),
	}
}

func (this *AutocompleteSystem) Insert(word string, time int) { //time = 0 mean this word exist in trie and just +1
	cur := this

	for _, c := range word {
		idx := getIndex(c)

		if cur.children[idx] == nil {
			cur.children[idx] = NewNode()
		}
		cur.children[idx].wordFreqMap[word] += time
		//next node
		cur = cur.children[idx]
	}
	cur.end = true
}

/**
 * Your AutocompleteSystem object will be instantiated and called as such:
 * obj := Constructor(sentences, times);
 * param_1 := obj.Input(c);
 */

//--------------- min heap------------------
type MinHeap []WordFreq

func initHeap() *MinHeap {
	h := MinHeap{}
	heap.Init(&h)
	return &h
}

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	if h[i].freq == h[j].freq {
		temp := []string{h[i].word, h[j].word}
		sort.Strings(temp)
		if temp[0] == h[i].word {
			return true
		}
		return false
	}
	if h[i].freq > h[j].freq {
		return true
	}
	return false
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(WordFreq))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
