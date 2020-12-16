package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_decodeBoardingPass(t *testing.T) {
	var row, column, id int

	row, column, id = decodeBoardingPass("AAA")
	assert.Equal(t, -1, row)
	assert.Equal(t, -1, column)
	assert.Equal(t, -1, id)

	row, column, id = decodeBoardingPass("FBFBBFFRLR")
	assert.Equal(t, 44, row)
	assert.Equal(t, 5, column)
	assert.Equal(t, 357, id)
}
