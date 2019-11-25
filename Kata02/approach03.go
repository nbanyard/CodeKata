package kata02

func KarateChop01(needle int, haystack []int) {
	for start, end := 0, len(haystack); start < end ; {
		mid := (start + end) / 2
		if haystack[mid] < needle {
			start = mid + 1
		} else if haystack[mid] > needle {
			end = mid
		} else {
			return mid
		}
	}

	return -1
}
