package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCutLast(t *testing.T) {
	assert.Equal(
		t,
		[]string{"one", "two"},
		cutLast([]string{"one", "two", "three"}),
		"cutLast should return a copy of a slice without last element",
	)
}

func TestLongestStringIn(t *testing.T) {
	assert.Equal(
		t,
		"the longest string",
		longestStringIn([]string{"small", "string", "the longest string", "medium string"}),
		"longestStringIn should return the longest string from a slice",
	)
}