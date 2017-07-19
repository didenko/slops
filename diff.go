package slops

// Diff takes sorted slices of strings and returns the non-
// overlapping entries
func Diff(left, right []string) (leftOnly, rightOnly []string) {

	leftOnly = make([]string, 0)
	rightOnly = make([]string, 0)

	for l, r := 0, 0; l < len(left) || r < len(right); {

		if l == len(left) || left[l] > right[r] {
			rightOnly = append(rightOnly, right[r])
			r++
			continue
		}

		if r == len(right) || left[l] < right[r] {
			leftOnly = append(leftOnly, left[l])
			l++
			continue
		}

		l++
		r++
	}

	return leftOnly, rightOnly
}
