package sentence_similarity_ii

/*
	problem: https://leetcode.com/problems/sentence-similarity-ii/


        at first, we will check 2 sentence follow index to remove all of similar words in 2 sentence
        build disjont set rely similarPair
			uniqueWordMap = map[string]bool
			run over pairs of sentence => create uniqueWordMap
			build disjontSet via uniqueWordMap, sentence1, sentence2
			if disjontSet.len

        run over index in sentence
            if root[sentence1[i]] != root[sentence2[i]] => return false

        return true


	time complexity: O(m+n)
		m = length of sentence
		n = length of pairs of sentence

*/

type disjointSet struct {
	size   int
	parent map[string]string
	rank   map[string]int
}

func areSentencesSimilarTwo(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {

	if len(sentence1) != len(sentence2) {
		return false
	}

	//remove all of same words in 2 sentence follow index
	for i := 0; i < len(sentence1); i++ {
		if sentence1[i] == sentence2[i] {
			sentence1 = append(sentence1[:i], sentence1[i+1:]...)
			sentence2 = append(sentence2[:i], sentence2[i+1:]...)
		}
	}
	if len(sentence1) == 0 {
		return true
	}

	// build uniqueWordMap
	uniqueWordMap := make(map[string]bool)
	for _, pair := range similarPairs {
		uniqueWordMap[pair[0]] = true
		uniqueWordMap[pair[1]] = true
	}
	for _, w := range sentence1 {
		uniqueWordMap[w] = true
	}
	for _, w := range sentence2 {
		uniqueWordMap[w] = true
	}

	//build disjoint set via similarPairs
	djs := initDisjointSet(uniqueWordMap)
	for _, pair := range similarPairs {
		djs.union(pair[0], pair[1])
	}

	//check sentence1 similar with sentence2 follow word by word
	for i := 0; i < len(sentence1); i++ {
		if djs.find(sentence1[i]) != djs.find(sentence2[i]) {
			return false
		}
	}
	return true
}

func initDisjointSet(uniqueWordMap map[string]bool) disjointSet {
	djs := disjointSet{
		size:   len(uniqueWordMap),
		parent: make(map[string]string),
		rank:   make(map[string]int),
	}
	for word, _ := range uniqueWordMap {
		djs.parent[word] = word
		djs.rank[word] = 0
	}
	return djs
}

func (this *disjointSet) union(x, y string) bool {
	rootX, rootY := this.find(x), this.find(y)
	if rootX == rootY {
		return false
	}

	if this.rank[rootX] > this.rank[rootY] {
		this.parent[rootY] = rootX
	} else if this.rank[rootY] > this.rank[rootX] {
		this.parent[rootX] = rootY
	} else {
		this.parent[rootY] = rootX
		this.rank[rootX]++
	}
	return true
}

func (this *disjointSet) find(x string) string {
	if this.parent[x] == x {
		return this.parent[x]
	} else {
		return this.find(this.parent[x])
	}
}
