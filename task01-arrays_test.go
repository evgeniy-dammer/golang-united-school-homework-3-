package homework

import "testing"

type testArrayCase struct {
	name  string
	input [15]float32
	want  float32
}

var testACases = []testArrayCase{
	{"one", [15]float32{1, 2, 3, 4, 5, 6}, 3.5},
	{"two", [15]float32{9, 10, 12, 13, 13, 13, 15, 15, 16, 16}, 13.2},
	{"three", [15]float32{1, 2, 10, 15, 15, 18, 24, 33, 48, 59, 61, 64, 89, 99, 185}, 48.2},
}

func TestAverage(t *testing.T) {
	for _, testCase := range testACases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := average(testCase.input)

			if actual != testCase.want {
				t.Errorf("got: %f; want: %f", actual, testCase.want)
			}
		})
	}
}
