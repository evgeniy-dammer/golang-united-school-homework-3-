package homework

import (
	"reflect"
	"testing"
)

type testSliceCase struct {
	name  string
	input []int64
	want  []int64
}

var testSCases = []testSliceCase{
	{"one", []int64{1, 2, 5, 15, 16}, []int64{16, 15, 5, 2, 1}},
	{"two", []int64{9, 10, 12, 13, 15, 16}, []int64{16, 15, 13, 12, 10, 9}},
	{"three", []int64{1, 2, 10, 15, 18, 24, 33}, []int64{33, 24, 18, 15, 10, 2, 1}},
}

func TestReverse(t *testing.T) {
	for _, testCase := range testSCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := reverse(testCase.input)

			if !reflect.DeepEqual(actual, testCase.want) {
				t.Errorf("got: %d; want: %d", actual, testCase.want)
			}
		})
	}
}
