package kata02

func KarateChop04(needle int, haystack []int) {
	start, end := 0, len(haystack)
	for length := end - 1; length != 0; length >>= 1 {
		mid := (start + end) >> 1
		if haystack[mid] > needle {
			end = mid
		} else {
			start = mid
		}
	}

	if end > 0 && haystack[start] == needle {
		return start
	}
	return -1
}
