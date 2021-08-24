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

func TestGetLowestCommonManager(t *testing.T) {
	e := &algoexpert.OrgChart{Name: "E"}
	f := &algoexpert.OrgChart{Name: "F"}
	g := &algoexpert.OrgChart{Name: "G"}
	h := &algoexpert.OrgChart{Name: "H"}
	i := &algoexpert.OrgChart{Name: "I"}
	d := &algoexpert.OrgChart{Name: "D", DirectReports: []*algoexpert.OrgChart{h, i}}
	c := &algoexpert.OrgChart{Name: "C", DirectReports: []*algoexpert.OrgChart{f, g}}
	b := &algoexpert.OrgChart{Name: "B", DirectReports: []*algoexpert.OrgChart{d, e}}
	a := &algoexpert.OrgChart{Name: "A", DirectReports: []*algoexpert.OrgChart{b, c}}

	org := a

	cases := []struct {
		given struct {
			org, reportOne, reportTwo *algoexpert.OrgChart
		}
		expected *algoexpert.OrgChart
	}{
		{
			struct {
				org, reportOne, reportTwo *algoexpert.OrgChart
			}{
				org:       org,
				reportOne: e,
				reportTwo: i,
			},
			b,
		},
		{
			struct {
				org, reportOne, reportTwo *algoexpert.OrgChart
			}{
				org:       org,
				reportOne: d,
				reportTwo: g,
			},
			a,
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("GetLowestCommonManager(%v, %v, %v)", tc.given.org, tc.given.reportOne, tc.given.reportTwo), func(t *testing.T) {
			actual := algoexpert.GetLowestCommonManager(tc.given.org, tc.given.reportOne, tc.given.reportTwo)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestPhoneNumberMnemonics(t *testing.T) {
	cases := []struct {
		given    string
		expected []string
	}{
		{
			"1905",
			[]string{
				"1w0j",
				"1w0k",
				"1w0l",
				"1x0j",
				"1x0k",
				"1x0l",
				"1y0j",
				"1y0k",
				"1y0l",
				"1z0j",
				"1z0k",
				"1z0l",
			},
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("PhoneNumberMnemonics(%v)", tc.given), func(t *testing.T) {
			actual := algoexpert.PhoneNumberMnemonics(tc.given)
			assert.ElementsMatch(t, tc.expected, actual)
		})
	}
}
