package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{30, 40, 70},
		{-1, -5, -6},
		{0, 0, 0},
		{15, -4, 11},
	}
	for _, val := range tests {
		result := Add(val.a, val.b)
		assert.Equal(t, val.want, result, "Not equal")
	}
}

func TestMul(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{30, 40, 1200},
		{-1, -5, 5},
		{0, 4654, 0},
		{15, -4, -60},
	}
	for _, val := range tests {
		result := Mul(val.a, val.b)
		assert.Equal(t, val.want, result, "Not equal")
	}
}
func TestSub(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{30, 40, 10},
		{40, 30, 10},
		{0, 4654, 4654},
		{15, -30, 45},
		{-50, -30, 20},
	}
	for _, val := range tests {
		result := Sub(val.a, val.b)
		assert.Equal(t, val.want, result, "Not equal")
	}
}
func TestReverse(t *testing.T) {
	var tests = []struct {
		a    int
		want int
	}{
		{30, 03},
		{-1, 0},
		{0, 0},
		{11, 11},
		{89848, 84898},
	}
	for _, val := range tests {
		result := ReverseNumber(val.a)
		assert.Equal(t, val.want, result, "Not equal")
	}
}
func TestCount(t *testing.T) {
	var testcase = []struct {
		str  string
		want int
	}{
		{"Hell", 4},
		{" ", 1},
		{"", 0},
	}
	for _, val := range testcase {
		result := Wordc(val.str)
		assert.Equal(t, val.want, result, "Not correct count")
	}
}
