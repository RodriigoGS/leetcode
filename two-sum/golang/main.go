package main

func TwoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if j, ok := hashMap[complement]; ok {
			return []int{i, j}
		}

		hashMap[num] = i
	}

	return nil
}
