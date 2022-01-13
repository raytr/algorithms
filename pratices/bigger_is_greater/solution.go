package bigger_is_greater

//time complexity: O(n)
//space complexity: O(1)
func biggerIsGreater(w string) string {

	n := len(w)
	a := uint8(0)
	aIndex, bIndex := -1, n-1
	for i := n - 1; i > 0; i-- {
		if w[i-1] < w[i] {
			a = w[i-1]
			aIndex = i - 1
			break
		}
	}

	//not found a
	if aIndex == -1 {
		return "no answer"
	}

	for i := n - 1; i > aIndex; i-- {
		if w[i] > a {
			bIndex = i
			break
		}
	}

	r := []rune(w)

	r[aIndex], r[bIndex] = r[bIndex], r[aIndex]

	//reverse all from aIndex + 1 to end
	for i := aIndex + 1; i < n-1; i++ {
		r[i], r[n-i+aIndex] = r[n-i+aIndex], r[i]
	}

	return string(r)
}
