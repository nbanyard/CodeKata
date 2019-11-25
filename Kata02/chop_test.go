package kata02

func TestApproach01(t *testing.T)
	testApproach(t, KarateChop01)
}

func testApproach(t *testing.T, approach([]int) int) {
	oneTest(t, approach, -1, 3, [])
	oneTest(t, approach, -1, 3, [1])
	oneTest(t, approach, 0, 1, [1])

	oneTest(t, approach, 0, 1, [1, 3, 5])
	oneTest(t, approach, 1, 3, [1, 3, 5])
	oneTest(t, approach, 2, 5, [1, 3, 5])
	oneTest(t, approach, -1, 0, [1, 3, 5])
	oneTest(t, approach, -1, 2, [1, 3, 5])
	oneTest(t, approach, -1, 4, [1, 3, 5])
	oneTest(t, approach, -1, 6, [1, 3, 5])

	oneTest(t, approach, 0, 1, [1, 3, 5, 7])
	oneTest(t, approach, 1, 3, [1, 3, 5, 7])
	oneTest(t, approach, 2, 5, [1, 3, 5, 7])
	oneTest(t, approach, 3, 7, [1, 3, 5, 7])
	oneTest(t, approach, -1, 0, [1, 3, 5, 7])
	oneTest(t, approach, -1, 2, [1, 3, 5, 7])
	oneTest(t, approach, -1, 4, [1, 3, 5, 7])
	oneTest(t, approach, -1, 6, [1, 3, 5, 7])
	oneTest(t, approach, -1, 8, [1, 3, 5, 7])

func oneTest(t *testing.T, approach(int, []int) int, expectation int, needle int, haystack []int) {
	actual := approach(needle, haystack)
	if expectation != actual {
		t.Errorf("Expected approach(%d, %v) to equal %d, got %d\n", needle, haystack, expectation, actual)
	}
}
