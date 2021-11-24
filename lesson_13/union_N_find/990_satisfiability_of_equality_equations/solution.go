package satisfiability_of_equality_equations

import "strings"

/*
	problem: https://leetcode.com/problems/satisfiability-of-equality-equations/



        we saprerate to 2 array : same and notSame
        same := [[a,b], [c,d]]
        nodeSame = [[d,f]]

        define a disjoinSet
            rank []int
            parent []string
            size int

        go over same array and union
        go over notSame array and find
            for ex: a & b
            if find(a) == find(b)=> return false
        erturn true




        func  union(a, b string)
            rootX = find(x)
            rootY = find(y)
            if range[rootX] > rank[rootY] => parent[rooY] = rootX
            else if rank[rootX] < rank[rootY] => parent[rootX] = rootY
            else parrent[rootY] = rootX; rank[rootX]++


        func find(x string){
            if parent[x] == x => return parent[x]
            else
                return find(x)

        }


	time complexity: O(len of equations)
	space complexity: O(len of equations)
*/

type disjoinSet struct {
	size   int
	parent map[string]string
	rank   map[string]int
}

func equationsPossible(equations []string) bool {

	theSame := make([][]string, 0, len(equations))
	theNotSame := make([][]string, 0, len(equations))
	uniqueChar := make(map[string]bool)
	for _, equation := range equations { //O(len of equations)
		if strings.Contains(equation, "==") {
			chars := strings.Split(equation, "==")
			uniqueChar[chars[0]] = true
			uniqueChar[chars[1]] = true
			theSame = append(theSame, chars)
		} else {
			chars := strings.Split(equation, "!=")
			uniqueChar[chars[0]] = true
			uniqueChar[chars[1]] = true
			theNotSame = append(theNotSame, chars)
		}
	}

	//init disjoin set
	djs := initDisjoinSet(uniqueChar)

	for _, pair := range theSame {
		djs.union(pair[0], pair[1])
	}

	for _, pair := range theNotSame {
		if djs.find(pair[0]) == djs.find(pair[1]) {
			return false
		}
	}

	return true
}

func initDisjoinSet(uniqueChar map[string]bool) disjoinSet {
	n := len(uniqueChar)
	djs := disjoinSet{
		size:   n,
		parent: make(map[string]string),
		rank:   make(map[string]int),
	}
	for k, _ := range uniqueChar {
		djs.parent[k] = k
		djs.rank[k] = 0
	}
	return djs
}

func (this *disjoinSet) union(x, y string) bool {
	rootX, rootY := this.parent[x], this.parent[y]

	if rootX == rootY {
		return false
	}

	if this.rank[rootX] > this.rank[rootY] {
		this.parent[rootY] = rootX
	} else if this.rank[rootX] < this.rank[rootY] {
		this.parent[rootX] = rootY
	} else {
		this.parent[rootY] = rootX
		this.rank[rootX]++
	}

	return true
}

func (this *disjoinSet) find(x string) string {
	if this.parent[x] == x {
		return this.parent[x]
	} else {
		return this.find(this.parent[x])
	}
}
