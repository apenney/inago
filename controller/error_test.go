package controller

import (
	"fmt"
	"testing"
)

func Test_Controller_maskAnyf(t *testing.T) {
	testCases := []struct {
		InputError  error
		InputFormat string
		InputArgs   []interface{}
		Expected    error
	}{
		{
			InputError:  nil,
			InputFormat: "",
			InputArgs:   []interface{}{},
			Expected:    nil,
		},
		{
			InputError:  fmt.Errorf("foo"),
			InputFormat: "bar",
			InputArgs:   []interface{}{},
			Expected:    nil,
		},
		{
			InputError:  fmt.Errorf("foo"),
			InputFormat: "bar %s",
			InputArgs:   []interface{}{"baz"},
			Expected:    fmt.Errorf("foo: bar baz"),
		},
	}

	for i, testCase := range testCases {
		var output error
		if len(testCase.InputArgs) == 0 {
			output = maskAnyf(testCase.InputError, testCase.InputFormat)
		} else {
			output = maskAnyf(testCase.InputError, testCase.InputFormat, testCase.InputArgs...)
		}

		if testCase.Expected != nil && output.Error() != testCase.Expected.Error() {
			t.Fatalf("test case %d: output '%s' != expected '%s'", i, output, testCase.Expected)
		}
	}
}
