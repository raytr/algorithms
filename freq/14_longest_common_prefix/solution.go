package longest_common_prefix

/*
	problem: https://leetcode.com/problems/longest-common-prefix/

   shortest word in array

   iterate all of them and check the first character
   obj is "f" => iterate all array => get first character to compare with obj => if 1 word not match => return res
   obj is fl =>

   res = ""
   objToCompare = flo
   flower flow flight => return res


   runtime is O(n * lenth of shortest word)
   space: O(1)

*/

func longestCommonPrefixBruteForce(strs []string) string {

	//find the shortest word
	shortestWord := strs[0]
	for _, word := range strs {
		if len(word) < len(shortestWord) {
			shortestWord = word
		}
	}

	//find the prefix
	res := ""
	for i, c := range shortestWord {
		for _, curW := range strs {
			if c != int32(curW[i]) {
				return res
			}
		}
		res += string(c)
	}
	return res
}
