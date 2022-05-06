package homework

import (
	"reflect"
	"testing"
)

type testMapCase struct {
	name  string
	input map[int]string
	want  []string
}

var testMCases = []testMapCase{
	{"one", map[int]string{2: "a", 0: "b", 1: "c"}, []string{"b", "c", "a"}},
	{"two", map[int]string{10: "aa", 0: "bb", 500: "cc"}, []string{"bb", "aa", "cc"}},
	{"three", map[int]string{34: "aaa", 49: "bbb", 2: "ccc"}, []string{"ccc", "aaa", "bbb"}},
}

func TestSortMapValues(t *testing.T) {
	for _, testCase := range testMCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := sortMapValues(testCase.input)

			if !reflect.DeepEqual(actual, testCase.want) {
				t.Errorf("got: %s; want: %s", actual, testCase.want)
			}
		})
	}
}
