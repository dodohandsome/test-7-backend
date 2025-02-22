package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBeef(t *testing.T) {
	mockText := `
		t-bone, bacon pastrami pork sausage jowl meatloaf ham
		fatback bresaola ham turkey pork. turkey bacon
		enim ipsum dolor sit amet consectetur adipising elit
	`
	expected := map[string]int{
		"t-bone":   1,
		"pastrami": 1,
		"pork":     2,
		"meatloaf": 1,
		"jowl":     1,
		"fatback":  1,
		"bresaola": 1,
		"enim":     1,
	}
	result := CountBeef(mockText)
	assert.Equal(t, expected, result)
}
