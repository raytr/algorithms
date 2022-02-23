package longest_word_in_dictionary

/*
	problem: https://leetcode.com/problems/longest-word-in-dictionary/

	   build a Trie on words. With each node
	   dfs
*/
type Trie struct {
	word     string
	end      bool
	children []*Trie
}

type LongestWord struct {
	val string
	len int
}

var curLongestWord LongestWord

func longestWord(words []string) string {

	curLongestWord = LongestWord{
		val: "",
		len: 0,
	}

	//int Trie
	trie := newTrie()

	// push all words into Trie
	for _, w := range words {
		trie.insert(w)
	}

	//dfs
	count := 0
	dfs(trie, count)

	return curLongestWord.val
}

// -------- dfs ---------
func dfs(trie *Trie, count int) {
	count++
	if len(trie.children) == 0 && trie.end {
		return
	}
	for _, c := range trie.children {
		if c != nil && c.end {
			if count > curLongestWord.len {
				curLongestWord = Max(curLongestWord, LongestWord{
					len: count,
					val: c.word,
				})
			}
			dfs(c, count)
		}
	}
}

// -------- Trie ---------

func newTrie() *Trie {
	return &Trie{
		end:      false,
		children: make([]*Trie, 26, 26),
	}
}

func (this *Trie) insert(word string) {
	cur := this
	for _, c := range word {
		if cur.children[c-'a'] == nil {
			cur.children[c-'a'] = newTrie()
			cur.children[c-'a'].word = cur.word + string(c)
		}
		cur = cur.children[c-'a']
	}
	cur.end = true
}

func (this *Trie) search(word string) bool {
	cur := this
	for _, c := range word {
		if cur.children[c-'a'] == nil {
			return false
		}
		cur = cur.children[c-'a']
	}
	return cur.end
}

func Max(a, b LongestWord) LongestWord {
	if a.len == b.len {
		aC, bC := a.val[a.len-1], b.val[b.len-1]
		if aC < bC {
			return a
		}
		return b
	}
	if a.len > b.len {
		return a
	}
	return b
}
