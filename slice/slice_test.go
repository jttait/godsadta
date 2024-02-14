package slice

import (
	"testing"

	"github.com/jttait/godsadta/assert"
)

func TestShouldBeTrueWhenSlicesAreSame(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	result := AreSlicesEqual(a, b)
	assert.AssertTrue(result, t)
}

func TestShouldBeFalseWhenStringSlicesAreNotSame(t *testing.T) {
	a := []string{"one", "two", "three"}
	b := []string{"four", "five", "six"}
	result := AreSlicesEqual(a, b)
	assert.AssertFalse(result, t)
}

func TestShouldBeTrueForEmptyStringSlices(t *testing.T) {
	a := []string{}
	b := []string{}
	result := AreSlicesEqual(a, b)
	assert.AssertTrue(result, t)
}

func TestShouldBeTrueWhenStringSlicesAreSame(t *testing.T) {
	a := []string{"one", "two", "three"}
	b := []string{"one", "two", "three"}
	result := AreSlicesEqual(a, b)
	assert.AssertTrue(result, t)
}

func TestShouldBeFalseWhenSlicesAreNotSame(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	result := AreSlicesEqual(a, b)
	assert.AssertFalse(result, t)
}

func TestShouldBeTrueForEmptySlices(t *testing.T) {
	a := []int{}
	b := []int{}
	result := AreSlicesEqual(a, b)
	assert.AssertTrue(result, t)
}
