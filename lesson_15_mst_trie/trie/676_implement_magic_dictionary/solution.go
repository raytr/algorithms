//package implement_magic_dictionary
package main

import "fmt"

/*
	problem: https://leetcode.com/problems/implement-magic-dictionary
*/

func main() {
	trie := Constructor()
	//trie.BuildDict([]string{"hello", "leetcode"})
	//fmt.Println(trie.Search("hello"))
	//fmt.Println(trie.Search("hhllo"))
	//fmt.Println(trie.Search("hell"))
	//fmt.Println(trie.Search("leetcoded"))

	trie.BuildDict([]string{"hello", "hallo", "leetcode"})
	fmt.Println(trie.Search("hello"))
	fmt.Println(trie.Search("hhllo"))
	fmt.Println(trie.Search("hell"))
	fmt.Println(trie.Search("leetcoded"))

}

type MagicDictionary struct {
	children []*MagicDictionary
	char     string
	end      bool
}

func Constructor() MagicDictionary {
	return MagicDictionary{
		children: make([]*MagicDictionary, 26, 26),
	}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	/*
	   add # at the last of word
	   run from index 1
	       trie.cur = word[i-1]
	       trie.next = word[i]


	   go over word by word, insert to the trie
	   ex: hello

	       cur = h
	       next = e
	       if !children[e] -> new node {char: e, next = h}
	*/

	for _, word := range dictionary {
		this.Insert(word)
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	/*
	   bypass = true
	   for i:= 0; i<len(searchWord)-1; i++,
	       c = searchWord[i]
	       nextC = searchWord[i+1]
	       if !trie.children[c - 'a'] => check bypass =>  check next:
	           -  if next == next character; bypass = false => keep go down
	           - otherwise return false

	   if cur.end == true  => reuturn true
	*/
	bypass := true
	cur := this
	for i := 0; i < len(searchWord); i++ {
		c := searchWord[i]
		a := string(c)
		fmt.Println(a)

		//because search word longer than the word in dictionary
		if cur.end {
			return false
		}

		if cur.children[c-'a'] == nil {
			if cur.end {
				break
			}
			if !cur.end {
				nextC := searchWord[i+1]
				if bypass {
					for _, trie := range cur.children {
						b := nextC - 'a'
						if trie != nil && trie.children[b] != nil {
							cur = trie.children[nextC-'a']
							i++
							bypass = false
						}
					}
				}
			}
		} else {
			cur = cur.children[c-'a']
		}

	}

	return cur.end && !bypass
}

func newTrie() *MagicDictionary {
	return &MagicDictionary{
		children: make([]*MagicDictionary, 26, 26),
	}
}

func (this *MagicDictionary) Insert(word string) {
	cur := this
	for _, c := range word {
		if cur.children[c-'a'] == nil {
			cur.children[c-'a'] = newTrie()
		}
		cur.children[c-'a'].char = string(c)
		cur = cur.children[c-'a']
	}
	cur.end = true
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
