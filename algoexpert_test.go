package algoexpert_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/algoexpert"
)

func TestIsValidSubsequence(t *testing.T) {
	cases := []struct {
		given struct {
			array, sequence []int
		}
		expected bool
	}{
		{
			struct {
				array, sequence []int
			}{
				[]int{5, 1, 22, 25, 6, -1, 8, 10},
				[]int{1, 6, -1, 10},
			},
			true,
		},
		{
			struct {
				array, sequence []int
			}{
				[]int{5, 1, 22, 25, 6, -1, 8, 10},
				[]int{0, 6, 11, 10},
			},
			false,
		},
		{
			struct {
				array, sequence []int
			}{
				[]int{5, 1, 22, 25, 6, -1, 8, 10},
				[]int{5, 1, 22, 22, 25, 6, -1, 8, 10},
			},
			false,
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("IsValidSubsequence(%v, %v)", tc.given.array, tc.given.sequence), func(t *testing.T) {
			actual := algoexpert.IsValidSubsequence(tc.given.array, tc.given.sequence)
			assert.Equal(t, tc.expected, actual)
		})
	}

}
