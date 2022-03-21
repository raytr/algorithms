package implement_str_str

/*
	problem: https://leetcode.com/problems/implement-strstr/

   clarificaiton:
       needle just include aphabet letters?
       what is maximum of length of haystack
       what is minimum of length of haystack

   iterate each character from haystack,
       if cur charcterr equal the first character from needle => check from this with needle
           needleI = 0
           allMatch = true
           j = i
           while needleI < len(needle) && i < n
               if haystack[j] != needle[needleI]
                   allMatch = false

               j++, needleI++

           if allMatch {
               return i
           }


   return -0


*/
//
func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		j := 0
		for ; j < len(needle); j++ {
			if needle[j] != haystack[i+j] {
				break
			}
		}
		if j == len(needle) {
			return i
		}
	}
	return -1
}

//func strStr(haystack string, needle string) int {
//	for i := 0; i < len(haystack)-len(needle)+1; i++ {
//		j := 0
//		for ; j < len(needle); j++ {
//			if needle[j] != haystack[i+j] {
//				break
//			}
//		}
//		if j == len(needle) {
//			return i
//		}
//	}
//	return -1
//}
