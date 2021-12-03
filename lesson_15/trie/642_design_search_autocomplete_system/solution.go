//package design_search_autocomplete_system
package main

/*
	problem: https://leetcode.com/problems/design-search-autocomplete-system/
*/

type MapSum struct {
	children []*MapSum
	sum      int
}

var curVal map[string]int

func NewMapSum() *MapSum {
	ms := &MapSum{
		children: make([]*MapSum, 26, 26),
	}
	return ms
}

func Constructor() MapSum {
	curVal = make(map[string]int)
	ms := NewMapSum()
	return *ms
}

func (this *MapSum) Insert(key string, val int) {
	/*
			   hasmap curVal is map[string]int
		       if exist
		            ex: [apple]3
		                new [apple]2
		                -> value to updata = new val - old value

		            curVal[key] = newVal

		                update value to updated to trie

	*/
	newVal := val
	oldVal := curVal[key]
	if _, exist := curVal[key]; exist {
		newVal = newVal - oldVal
	}
	curVal[key] = val

	cur := this
	for _, c := range key {
		if cur.children[c-'a'] == nil {
			cur.children[c-'a'] = NewMapSum()
			cur.children[c-'a'].sum = newVal
		} else {
			cur.children[c-'a'].sum += newVal
		}
		cur = cur.children[c-'a']
	}
}

func (this *MapSum) Sum(prefix string) int {
	/*
			   cur = this
			   go down in trie by character
		            if cur.children[c - 'a'] == nil => return 0
			   return cur.sum
	*/
	cur := this
	for _, c := range prefix {
		if cur.children[c-'a'] == nil {
			return 0
		}
		cur = cur.children[c-'a']
	}
	return cur.sum
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
