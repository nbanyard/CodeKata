package kata02

func KarateChop01(needle int, haystack []int) {
	if needle == haystack[0] {
		return 0
	}

	for start, end := 0, len(haystack); start + 1 < end ; {
		mid := (start + end) / 2
		if haystack[mid] < needle {
			start = mid
		} else if haystack[mid] > needle {
			end = mid
		} else {
			return mid
		}
	}
}
