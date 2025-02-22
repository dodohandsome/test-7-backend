package main

import (
	"fmt"
	"strings"
)

func isValid(nums []int, encoded string) bool {
	for i := 0; i < len(encoded); i++ {
		switch encoded[i] {
		case 'L':
			if nums[i] <= nums[i+1] {
				return false
			}
		case 'R':
			if nums[i] >= nums[i+1] {
				return false
			}
		case '=':
			if nums[i] != nums[i+1] {
				return false
			}
		}
	}
	return true
}

func backtrack(nums []int, index int, encoded string) bool {
	if index == len(nums) {
		return isValid(nums, encoded)
	}
	for i := 0; i <= 9; i++ {
		nums[index] = i
		if backtrack(nums, index+1, encoded) {
			return true
		}
	}
	return false
}

func decodeLowestSum(encoded string) string {
	nums := make([]int, len(encoded)+1)
	if !backtrack(nums, 0, encoded) {
		return ""
	}
	var result strings.Builder
	for _, num := range nums {
		result.WriteString(fmt.Sprintf("%d", num))
	}
	return result.String()
}

func main() {
	for {
		var input string
		fmt.Print("กรุณาใส่ค่า encoded string: ")
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("เกิดข้อผิดพลาดในการอ่านค่า input:", err)
			return
		}
		for _, char := range input {
			if char != 'L' && char != 'R' && char != '=' {
				fmt.Println("input ไม่ถูกต้อง: ต้องมีเฉพาะ 'L', 'R', และ '='")
				return
			}
		}

		result := decodeLowestSum(input)
		fmt.Printf("ผลลัพธ์: %s\n", result)

	}
}
