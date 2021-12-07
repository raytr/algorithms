//package design_search_autocomplete_system
package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
	problem: https://leetcode.com/problems/design-search-autocomplete-system/
*/

func main() {
	autocompleteSystem := Constructor([]string{"i love you", "island", "iroman", "i love leetcode"}, []int{5, 3, 2, 2})
	fmt.Println(autocompleteSystem.Input('i'))
	fmt.Println(autocompleteSystem.Input(' '))
	fmt.Println(autocompleteSystem.Input('l'))
	fmt.Println(autocompleteSystem.Input('#'))
}

type AutocompleteSystem struct {
	children    []*AutocompleteSystem
	end         bool
	wordFreqMap map[string]int
}
type WordFreq struct {
	word string
	freq int
}

var inputed string

func Constructor(sentences []string, times []int) AutocompleteSystem {
	/*
	   build trie, with each node, we add wordFreqMap follow sentences and time
	*/
	inputed = ""
	node := NewNode()
	for i := 0; i < len(sentences); i++ {
		node.Insert(sentences[i], times[i])
	}
	return *node
}

func (this *AutocompleteSystem) Input(c byte) []string {
	//store c
	cur := this
	if c == '#' {
		cur.Insert(inputed, 1)
	} else {
		inputed += string(c)
	}

	res := make([]string, 0, 3)

	for _, char := range inputed {
		idx := int32(0)
		if char == ' ' {
			idx = 27
		} else {
			idx = char - 'a'
		}

		if cur.children[idx] == nil {
			break
		}

		minHeap := initHeap()
		for k, v := range cur.children[idx].wordFreqMap {
			heap.Push(minHeap, WordFreq{k, v})
			if minHeap.Len() > 3 {
				heap.Pop(minHeap)
			}
		}

		for minHeap.Len() > 0 {
			res = append(res, heap.Pop(minHeap).(WordFreq).word)
		}
	}
	return res
}

func NewNode() *AutocompleteSystem {
	return &AutocompleteSystem{
		children:    make([]*AutocompleteSystem, 28, 28),
		wordFreqMap: make(map[string]int),
	}
}

func (this *AutocompleteSystem) Insert(word string, time int) { //time = 0 mean this word exist in trie and just +1
	cur := this
	idx := int32(0)
	if word == " " {
		idx = 27
	}

	for _, c := range word {
		if c != ' ' {
			idx = c - 'a'
		}
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
	if h[i].freq < h[j].freq {
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
