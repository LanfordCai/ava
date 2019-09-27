package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	s1 := []string{"a", "b", "c"}
	s2 := []string{}

	assert.True(t, Contains(s1, "a"))
	assert.False(t, Contains(s1, "d"))
	assert.False(t, Contains(s2, "a"))
}
