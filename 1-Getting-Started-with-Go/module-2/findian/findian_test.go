package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”. The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.
*/
func TestFindian(t *testing.T) {
	assert := assert.New(t)
	positiveTestCases := [4]string{"ian", "Ian", "iuiygaygn", "I d skd a efju N"}
	negativeTestCases := [3]string{"ihhhhhn", "ina", "xian"}
	for _, testCase := range positiveTestCases {
		assert.Equal(findian(testCase), "Found!", "These cases are positive.")
	}
	for _, testCase := range negativeTestCases {
		assert.Equal(findian(testCase), "Not Found!", "These cases are negative.")
	}
}
