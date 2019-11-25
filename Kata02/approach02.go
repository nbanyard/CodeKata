package kata02

func KarateChop02(needle int, haystack []int) {
	if len(haystack) == 0 {
		return -1
	}

	mid := len(haystack) / 2
	if haystack[mid] < needle {
		return KarateChop02(needle, haystack[mid + 1:])
	} else if haystack[mid] > needle {
		return KarateChop02(needle, haystack[:mid])
	} else {
		return mid
	}
}
