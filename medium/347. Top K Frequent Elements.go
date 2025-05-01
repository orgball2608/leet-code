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

type Pair struct {
	Key   int
	Count int
}

func topKFrequent(nums []int, k int) []int {
	hash := make(map[int]int)
	for _, val := range nums {
		hash[val]++
	}

	counter := make([]Pair, 0, len(hash))
	for key, count := range hash {
		counter = append(counter, Pair{
			Key:   key,
			Count: count,
		})
	}

	rand.Seed(time.Now().UnixNano())
	quickSelect(counter, 0, len(counter)-1, k-1)

	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, counter[i].Key)
	}
	return res
}

func quickSelect(nums []Pair, start, end int, k int) {
	if end == start {
		return
	}

	randomIndex := start + rand.Intn(end-start+1)
	pivot := nums[randomIndex]
	swap(nums, randomIndex, end)

	currentLeft := start
	for i := start; i <= end; i++ {
		if nums[i].Count > pivot.Count {
			swap(nums, i, currentLeft)
			currentLeft++
		}
	}

	swap(nums, end, currentLeft)
	currentRight := end
	for i := end; i > currentLeft; i-- {
		if nums[i].Count < pivot.Count {
			swap(nums, i, currentRight)
			currentRight--
		}
	}

	if currentLeft <= k && k <= currentRight {
		return
	} else if currentLeft > k {
		quickSelect(nums, start, currentLeft-1, k)
	} else {
		quickSelect(nums, currentRight+1, end, k)
	}
}

func swap(nums []Pair, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

// O(N)
// O(N)
