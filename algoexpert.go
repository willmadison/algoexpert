package algoexpert

import (
	"errors"
)

func IsValidSubsequence(array, sequence []int) bool {
	if len(sequence) > len(array) {
		return false
	}

	var currentIndex int

	for _, v := range array {
		if v == sequence[currentIndex] {
			currentIndex++
		}

		if currentIndex >= len(sequence) {
			return true
		}
	}

	return false
}

type OrgChart struct {
	Name          string
	DirectReports []*OrgChart
}

type Queue interface {
	Enqueue(*OrgChart)
	Dequeue() (*OrgChart, error)
	Peek() (*OrgChart, error)
	IsEmpty() bool
}
type orgChartQueue struct {
	data []*OrgChart
	size int
}

func NewOrgChartQueue() Queue {
	return &orgChartQueue{data: []*OrgChart{}}
}

func (o *orgChartQueue) Enqueue(value *OrgChart) {
	o.data = append(o.data, value)
	o.size += 1
}

func (o *orgChartQueue) Dequeue() (*OrgChart, error) {
	if o.size > 0 {
		value := o.data[0]
		o.size--
		o.data = o.data[1:]

		return value, nil
	}

	return nil, errors.New("No Such Element")
}

func (o *orgChartQueue) Peek() (*OrgChart, error) {
	if o.size > 0 {
		value := o.data[0]
		return value, nil
	}

	return nil, errors.New("No Such Element")
}

func (o *orgChartQueue) IsEmpty() bool {
	return o.size == 0
}

type Ancestry []string

func GetLowestCommonManager(org, reportOne, reportTwo *OrgChart) *OrgChart {
	var employeesByName map[string]*OrgChart

	firstAncestry, lookupA := determineAncestry(org, reportOne)
	secondAncestry, lookupB := determineAncestry(org, reportTwo)

	employeesByName = lookupA

	if len(lookupB) > len(employeesByName) {
		employeesByName = lookupB
	}

	var lowestCommonManagerName string

	for lowestCommonManagerName == "" {
		if firstAncestry[0] == secondAncestry[0] {
			lowestCommonManagerName = firstAncestry[0]
			break
		}

		if len(firstAncestry) > len(secondAncestry) {
			firstAncestry = firstAncestry[1:]
		} else if len(secondAncestry) > len(firstAncestry) {
			secondAncestry = secondAncestry[1:]
		} else {
			firstAncestry = firstAncestry[1:]
			secondAncestry = secondAncestry[1:]
		}
	}

	return employeesByName[lowestCommonManagerName]
}

func determineAncestry(root, target *OrgChart) (Ancestry, map[string]*OrgChart) {
	queue := NewOrgChartQueue()

	var current *OrgChart

	managersByEmployee := map[string]*OrgChart{
		root.Name: nil,
	}
	employeesByName := map[string]*OrgChart{}
	queue.Enqueue(root)

	for !queue.IsEmpty() && current != target {
		current, _ = queue.Dequeue()

		employeesByName[current.Name] = current

		if current == target {
			break
		}

		for _, report := range current.DirectReports {
			managersByEmployee[report.Name] = current
			queue.Enqueue(report)
		}
	}

	manager := managersByEmployee[target.Name]

	var ancestry Ancestry
	ancestry = append(ancestry, target.Name)

	for manager != nil {
		ancestry = append(ancestry, manager.Name)
		manager = managersByEmployee[manager.Name]
	}

	return ancestry, employeesByName
}

var lettersByDigit = map[rune][]string{
	'1': {"1"},
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
	'0': {"0"},
}

func PhoneNumberMnemonics(phoneNumber string) []string {
	if len(phoneNumber) == 0 {
		return []string{""}
	}

	var mnemonics []string

	digit := phoneNumber[0]

	subMnemonics := PhoneNumberMnemonics(phoneNumber[1:])

	for _, letter := range lettersByDigit[rune(digit)] {
		for _, m := range subMnemonics {
			mnemonics = append(mnemonics, letter+m)
		}
	}

	return mnemonics
}
