package array

import "sort"

// 题15： 三数之和
// 找出整型数组nums中， 满足a+b+c=0的元素组合
// nums = [-1,0,1,2,-1,-4]
// [[-1,-1,2],[-1,0,1]]

// 排序 + 双指针
func threeSums(nums []int) [][]int {
	sort.Ints(nums)

	var result [][]int

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 { // 当前数大于0， 则三数之和一定大于0
			break
		}

		if i > 0 && nums[i] == nums[i-1] { // 枚举去重
			continue
		}

		lft, rgt := i+1, len(nums)-1

		for lft < rgt {
			sum := nums[i] + nums[lft] + nums[rgt]

			if sum == 0 {
				result = append(result, []int{nums[i], nums[lft], nums[rgt]})

				for lft < rgt && nums[rgt] == nums[rgt-1] {
					rgt-- // 去重
				}

				for lft < rgt && nums[lft] == nums[lft+1] {
					lft++ // 去重
				}

				rgt--
				lft++
			} else if sum < 0 {
				lft++
			} else {
				rgt--
			}
		}
	}

	return result
}
