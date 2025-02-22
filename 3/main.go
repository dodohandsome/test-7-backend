package main

import (
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func fetchTextFromAPI() (string, error) {
	resp, err := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func CountBeef(text string) map[string]int {
	re := regexp.MustCompile(`\b[a-zA-Z-]+\b`)
	words := re.FindAllString(text, -1)
	beefCounts := make(map[string]int)
	for _, word := range words {
		word = strings.ToLower(word)
		if isBeef(word) {
			beefCounts[word]++
		}
	}
	return beefCounts
}

func isBeef(word string) bool {
	beefWords := map[string]bool{
		"t-bone":   true,
		"fatback":  true,
		"pastrami": true,
		"pork":     true,
		"meatloaf": true,
		"jowl":     true,
		"enim":     true,
		"bresaola": true,
	}
	return beefWords[word]
}

func GetBeefSummary(c *gin.Context) {
	text, err := fetchTextFromAPI()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}
	beefCounts := CountBeef(text)
	c.JSON(http.StatusOK, gin.H{
		"beef": beefCounts,
	})
}

func main() {
	router := gin.Default()
	router.GET("/beef/summary", GetBeefSummary)
	router.Run(":8080")
}
