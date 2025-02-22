package main

import (
	"fmt"
	"strings"
)

func decodeLowestSum(encoded string) string {
	n := len(encoded) + 1
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = 999999
		}
	}
	for i := 0; i < n; i++ {
		dp[0][i] = i
	}
	for i := 1; i < n; i++ {
		for prev := 0; prev < n; prev++ {
			for curr := 0; curr < n; curr++ {
				switch encoded[i-1] {
				case 'L':
					if prev <= curr {
						continue
					}
				case 'R':
					if prev >= curr {
						continue
					}
				case '=':
					if prev != curr {
						continue
					}
				}
				if dp[i-1][prev]+curr < dp[i][curr] {
					dp[i][curr] = dp[i-1][prev] + curr
				}
			}
		}
	}
	minSum := 999999
	lastDigit := 0
	for i := 0; i < n; i++ {
		if dp[n-1][i] < minSum {
			minSum = dp[n-1][i]
			lastDigit = i
		}
	}
	result := make([]int, n)
	result[n-1] = lastDigit
	for i := n - 2; i >= 0; i-- {
		for prev := 0; prev < n; prev++ {
			switch encoded[i] {
			case 'L':
				if prev > result[i+1] && dp[i][prev]+result[i+1] == dp[i+1][result[i+1]] {
					result[i] = prev
					break
				}
			case 'R':
				if prev < result[i+1] && dp[i][prev]+result[i+1] == dp[i+1][result[i+1]] {
					result[i] = prev
					break
				}
			case '=':
				if prev == result[i+1] && dp[i][prev]+result[i+1] == dp[i+1][result[i+1]] {
					result[i] = prev
					break
				}
			}
		}
	}
	var res strings.Builder
	for _, num := range result {
		res.WriteString(fmt.Sprintf("%d", num))
	}
	return res.String()
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
