package remove_sub_folders_from_the_filesystem

import (
	"strings"
)

/*
	problem: https://leetcode.com/problems/remove-sub-folders-from-the-filesystem/

       build a trie, with each node :
           end bool
           children *node

   when create node: we have root is true

   when insert, go throughout top down,
       if meet a node has end == true, update the root is false

   for f is range folder
       func checkRoot(word string)
           go top down,
			if meet a node with end is true && word[last] != character => break
           	meet a node with end is true && word[last] == character => add to res

   return res
*/

type Trie struct {
	children map[string]*Trie
	name     string
	end      bool
}

var res []string

func removeSubfolders(folder []string) []string {
	//redefine res
	res = make([]string, 0, len(folder))
	trie := newNode()

	//insert all word to trie
	for _, path := range folder {
		trie.Insert(path)
	}

	for _, path := range folder {
		trie.checkRoot(path)
	}
	return res
}

func newNode() *Trie {
	return &Trie{
		children: make(map[string]*Trie),
		end:      false,
	}
}

func (this *Trie) checkRoot(word string) {
	//split to array of string
	arr := strings.Split(word, "/")
	arr = arr[1:]

	cur := this
	for _, ele := range arr {
		if cur.children[ele].end {
			if cur.children[ele].name == arr[len(arr)-1] {
				res = append(res, word)
			} else {
				break
			}
		}
		cur = cur.children[ele]
	}
}

func (this *Trie) Insert(word string) {
	//split to array of string
	arr := strings.Split(word, "/")
	arr = arr[1:]
	cur := this
	for _, ele := range arr {
		if _, exist := cur.children[ele]; !exist {
			cur.children[ele] = newNode()
			cur.children[ele].name = ele
		}
		cur = cur.children[ele]
	}
	cur.end = true
}
