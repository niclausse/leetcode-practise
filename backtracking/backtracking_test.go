package backtracking

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCombine(t *testing.T) {
	convey.Convey("test combine", t, func() {
		res := combine(4, 2)
		expected := [][]int{
			{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4},
		}
		convey.So(res, convey.ShouldResemble, expected)
	})

}

func TestCombinationSum2(t *testing.T) {
	convey.Convey("test combinationSum2", t, func() {
		// 1 2 2 4 5 6 7 10
		res := combinationSum2([]int{1, 2, 2, 4, 5, 6, 7, 10}, 8)
		expected := [][]int{
			{1, 2, 5},
			{1, 7},
			{2, 2, 4},
			{2, 6},
		}
		convey.So(res, convey.ShouldResemble, expected)
	})
}

func TestCombinationSum3(t *testing.T) {
	convey.Convey("test combinationSum3", t, func() {
		res := combinationSum3(7, 3)
		expected := [][]int{{1, 2, 4}}
		convey.So(res, convey.ShouldResemble, expected)
	})
}

func TestCombinationSum(t *testing.T) {
	convey.Convey("test combinationSum", t, func() {
		res := combinationSum([]int{2, 3, 6, 7}, 7)
		expected := [][]int{{2, 2, 3}, {7}}
		convey.So(res, convey.ShouldResemble, expected)
	})
}

func TestLetterCombination(t *testing.T) {
	convey.Convey("test letterCombination", t, func() {
		res := letterCombinations("23")
		expected := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
		convey.So(res, convey.ShouldResemble, expected)
	})
}

func TestPartition(t *testing.T) {
	convey.Convey("test partition", t, func() {
		res := partition("aab")
		expected := [][]string{
			{"a", "a", "b"},
			{"aa", "b"},
		}
		convey.So(res, convey.ShouldResemble, expected)
	})
}

func TestRestoreIP(t *testing.T) {
	convey.Convey("test RestoreIpAddress", t, func() {
		res := restoreIpAddresses("25525511135")
		expected := []string{"255.255.11.135", "255.255.111.35"}
		convey.So(res, convey.ShouldResemble, expected)
	})
}

func TestSubsets(t *testing.T) {
	convey.Convey("test subsets", t, func() {
		res := subsets([]int{1, 2, 3})
		expected := [][]int{
			{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3},
		}
		convey.So(res, convey.ShouldResemble, expected)
	})
}

func TestSubsets2(t *testing.T) {
	convey.Convey("test subsets2", t, func() {
		res := subsets2([]int{1, 2, 2})
		expected := [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}
		convey.So(res, convey.ShouldResemble, expected)
	})
}

func TestODSkipHouse(t *testing.T) {
	convey.Convey("test odSkipHouse", t, func() {
		convey.Convey("case1", func() {
			res := skipHouse(15, []int{1, 9, 4, 25, 10, 8, 7, 5})
			expected := []int{1, 4, 10}
			convey.So(res, convey.ShouldResemble, expected)
		})
		convey.Convey("case2", func() {
			res := skipHouse(9, []int{1, 4, 5, 2, 0, 2})
			expected := []int{4, 5, 0}
			convey.So(res, convey.ShouldResemble, expected)
		})
	})
}

func TestPermuteUniqueBySort(t *testing.T) {
	convey.Convey("test permuteUniqueBySort", t, func() {
		res := permuteUniqueBySort([]int{1, 1, 2})
		expected := [][]int{
			{1, 1, 2}, {1, 2, 3}, {2, 1, 1},
		}
		convey.So(res, convey.ShouldResemble, expected)
	})
}
