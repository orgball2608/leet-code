package medium

import "sort"

func threeSum(nums []int) [][]int {
	var results [][]int
	if len(nums) < 3 {
		return results
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		target := -nums[i]

		for left < right {
			sum := nums[left] + nums[right]

			if sum == target {
				results = append(results, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return results
}

//first version

// func threeSum(nums []int) [][]int {
//     sort.Ints(nums)
//     var results [][]int
//     for i:= 0; i < len(nums) - 2; i ++ {
//         if i> 0 && nums[i] == nums[i-1] {
//             continue;
//         }
//         twoSumResults := searchPair(nums, -nums[i], i + 1)
//         // Update results array
//         results = combine2dArrays(results, twoSumResults)
//     }
//     return results
// }

// func combine2dArrays(leftArr [][]int, rightArr [][]int) [][]int {
//     return append(leftArr, rightArr...)
// }

// func searchPair(nums []int, target int, left int) [][]int {
//     var results [][]int
//     right := len(nums) - 1
//     for left < right {
//         curSum := nums[left] + nums[right]
//         if curSum == target {
//             results = append(results, []int{nums[left], nums[right], target * -1})
//             left++
//             right--
//             for left < right && nums[left] == nums[left-1] {
//                 left++
//             }
//             for left < right && nums[right] == nums[right+1] {
//                 right--
//             }
//         } else if curSum < target {
//             left++
//         } else {
//             right--
//         }
//     }
//     return results

// }
