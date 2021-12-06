//package implement_trie
package main

import "fmt"

/*
	problem: https://leetcode.com/problems/implement-trie-prefix-tree/
*/

func main() {
	t := Constructor()
	//t.Insert("apple")
	//fmt.Println(t.Search("apple"))
	//fmt.Println(t.Search("app"))
	//fmt.Println(t.StartsWith("app"))
	//t.Insert("app")
	//fmt.Println(t.Search("app"))

	t.Insert("hello")
	fmt.Println(t.Search("hell"))
	fmt.Println(t.Search("helloa"))
	fmt.Println(t.Search("hello"))
	fmt.Println(t.StartsWith("hell"))
	fmt.Println(t.StartsWith("helloa"))
	fmt.Println(t.StartsWith("hello"))
}

type Trie struct {
	val      string
	end      bool
	children []*Trie
}

func node() *Trie {
	node := &Trie{
		end:      false,
		children: make([]*Trie, 26, 26),
	}
	return node
}

func Constructor() Trie {
	node := node()
	return *node
}

func (this *Trie) Insert(word string) {
	cur := this
	for _, c := range word {
		if cur.children[c-'a'] == nil {
			cur.children[c-'a'] = node()
			cur.children[c-'a'].val = string(c)
		}
		cur = cur.children[c-'a']
	}
	cur.end = true
}

func (this *Trie) Search(word string) bool {
	cur := this
	for _, c := range word {
		if cur.children[c-'a'] == nil {
			return false
		}
		cur = cur.children[c-'a']
	}
	return cur.end
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for _, c := range prefix {
		if (cur.children[c-'a']) == nil {
			return false
		}
		cur = cur.children[c-'a']
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
