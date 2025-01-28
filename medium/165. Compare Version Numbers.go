package medium

func compareVersion(version1 string, version2 string) int {
	rev1 := strings.Split(version1, ".")
	rev2 := strings.Split(version2, ".")

	revInt1 := make([]int, 0, len(rev1))
	revInt2 := make([]int, 0, len(rev2))
	for _, val := range rev1 {
		cInt, _ := strconv.Atoi(val)
		revInt1 = append(revInt1, cInt)
	}

	for _, val := range rev2 {
		cInt, _ := strconv.Atoi(val)
		revInt2 = append(revInt2, cInt)
	}

	i, j := 0, 0
	for i < len(revInt1) && j < len(revInt2) {
		if revInt1[i] < revInt2[j] {
			return -1
		} else if revInt1[i] > revInt2[j] {
			return 1
		}
		i++
		j++
	}

	for i < len(revInt1) {
		if revInt1[i] > 0 {
			return 1
		}
		i++
	}

	for j < len(revInt2) {
		if revInt2[j] > 0 {
			return -1
		}
		j++
	}

	return 0
}

// Time: O(N)
// Space: O(N)
