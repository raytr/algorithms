package search_suggestions_system

import (
	"sort"
)

/*
	problem:  https://leetcode.com/problems/search-suggestions-system/

	   sort ~ O(nlogN)
	   build the trie with trie {
	       words []string
	       children []*trie
	       end bool
	   } ~ O(n)

	   with each character in search_word, start with(subword) ~ O(n)
	   push to heap
*/

type Trie struct {
	words    []string
	curWord  string
	end      bool
	children []*Trie
}

func suggestedProducts(products []string, searchWord string) [][]string {

	// sort
	sort.Strings(products)

	//init trie
	trie := newTrie()
	for _, p := range products {
		trie.insert(p)
	}

	res := make([][]string, 0, len(searchWord))
	//search het
	for i, w := range searchWord {
		res = append(res, []string{})
		if trie.startWith(string(w)) {
			res[i] = trie.children[w-'a'].words
			trie = trie.children[w-'a']
		}
	}
	return res
}

func newTrie() *Trie {
	return &Trie{
		end:      false,
		words:    make([]string, 0),
		children: make([]*Trie, 26, 26),
	}
}

func (this *Trie) insert(word string) {
	cur := this

	//insert to trie ~ O(n)
	for _, c := range word {
		if cur.children[c-'a'] == nil {
			cur.children[c-'a'] = newTrie()
			cur.children[c-'a'].curWord = cur.curWord + string(c)
		}
		if len(cur.children[c-'a'].words) < 3 {
			cur.children[c-'a'].words = append(cur.children[c-'a'].words, word)
		}
		cur = cur.children[c-'a']
	}
	cur.end = true
}

func (this *Trie) startWith(prefix string) bool {
	cur := this
	for _, c := range prefix {
		if cur.children[c-'a'] == nil {
			return false
		}
		cur = cur.children[c-'a']
	}
	return true
}
