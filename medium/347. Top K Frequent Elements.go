package medium

func topKFrequent(nums []int, k int) []int {
	hashmap := make(map[int]int)
	for _, val := range nums {
		hashmap[val]++
	}

	pairs := make([][]int, len(nums))

	for k, el := range hashmap {
		index := el - 1
		if pairs[index] == nil {
			pairs[index] = []int{k}
		} else {
			pairs[index] = append(pairs[index], k)
		}
	}

	result := make([]int, k)
	idx := 0
	for i := len(pairs) - 1; i >= 0; i-- {
		if idx == k {
			break
		}

		if pairs[i] != nil {
			for _, val := range pairs[i] {
				result[idx] = val
				idx++
			}
		}
	}
	return result
}

// O(N)
// O(N)

//func topKFrequent(nums []int, k int) []int {
//	hashmap := make(map[int]int)
//	for _, val := range nums {
//		hashmap[val]++
//	}
//
//	type pair struct {
//		value     int
//		frequency int
//	}
//
//	pairs := make([]pair, 0, len(hashmap))
//
//	for k, el := range hashmap {
//		pairs = append(pairs, pair{value: k, frequency: el})
//	}
//
//	sort.Slice(pairs, func(i, j int) bool {
//		return pairs[i].frequency > pairs[j].frequency
//	})
//
//	var result []int
//	for i := 0; i < k; i++ {
//		result = append(result, pairs[i].value)
//	}
//	return result
//}
//
//// O(NlogN)
//// O(N)
