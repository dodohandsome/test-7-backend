package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func readTriangleFromFile(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("ไม่สามารถเปิดไฟล์: %v", err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("เกิดข้อผิดพลาดในการอ่านไฟล์: %v", err)
	}
	var triangle [][]int
	err = json.Unmarshal(data, &triangle)
	if err != nil {
		log.Fatalf("เกิดข้อผิดพลาดในการแปลง JSON: %v", err)
	}
	return triangle
}

func maxPathSum(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += max(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	filename := "input.json"
	triangle := readTriangleFromFile(filename)
	result := maxPathSum(triangle)
	fmt.Println("ผลรวมเส้นทางที่มากที่สุด:", result)
}
