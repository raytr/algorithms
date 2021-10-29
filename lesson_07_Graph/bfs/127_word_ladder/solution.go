package word_ladder

/*
problem: https://leetcode.com/problems/word-ladder

bfs
build map using missed character, ex:
hot => *ot => transformWordMap[*ot] = [hot]
       h*t => transformWordMap[h*t] = [hot]
       ho* => transformWordMap[ho*] = [hot]

transWordMap

queue.push(beginWork)
while !queue.isEmpty
for word range queue
	if wordCanTransformWords.containt(endword) => return dis

	word => wordCanTransformWords
	if not => queue.push(wordCanTransformWords) (!visited)
}

Return -1

Complexity O(n)

*/

func ladderLength(beginWord string, endWord string, wordList []string) int {
	transformWordMap := make(map[string][]string)
	wordList = append([]string{beginWord}, wordList...)
	q := []string{beginWord}
	dis := 0
	visited := make(map[string]bool)

	//build transformWordMap
	for _, w := range wordList {
		for i := 0; i < len(w); i++ {
			missedCharWord := w[0:i] + "*" + w[i+1:]
			transformWordMap[missedCharWord] = append(transformWordMap[missedCharWord], w)
		}
	}

	for len(q) > 0 {
		for _, word := range q {
			visited[word] = true
			if word == endWord {
				return dis + 1
			}

			//check in transformWordMap
			for i := 0; i < len(word); i++ {
				newWord := word[0:i] + "*" + word[i+1:]
				for _, w := range transformWordMap[newWord] {
					if !visited[w] {
						q = append(q, w)
					}
				}
			}
			q = q[1:]
		}
		dis++
	}
	return 0
}
