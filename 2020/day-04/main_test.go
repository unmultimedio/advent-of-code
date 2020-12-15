package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validInt(t *testing.T) {
	assert.True(t, validInt("10", 10, 15))
	assert.True(t, validInt("15", 10, 15))
	assert.False(t, validInt("9", 10, 15))
	assert.False(t, validInt("16", 10, 15))
}

func Test_validHeight(t *testing.T) {
	assert.True(t, validHeight("190cm"))
	assert.True(t, validHeight("150cm"))
	assert.True(t, validHeight("151cm"))
	assert.True(t, validHeight("193cm"))
	assert.True(t, validHeight("60in"))
	assert.True(t, validHeight("59in"))
	assert.True(t, validHeight("76in"))

	assert.False(t, validHeight("149cm"))
	assert.False(t, validHeight("194cm"))
	assert.False(t, validHeight("190in"))
	assert.False(t, validHeight("58in"))
	assert.False(t, validHeight("77in"))
	assert.False(t, validHeight("190"))
}

func Test_validHairColor(t *testing.T) {
	assert.True(t, validHairColor("#123abc"))
	assert.False(t, validHairColor("#123abz"))
	assert.False(t, validHairColor("123abc"))
}

func Test_validEyeColor(t *testing.T) {
	assert.True(t, validEyeColor("brn"))
	assert.False(t, validEyeColor("wat"))
}

func Test_validPassportID(t *testing.T) {
	assert.True(t, validPassportID("000000001"))
	assert.False(t, validPassportID("0123456789"))
}
