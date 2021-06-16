package clavis_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fiuskylab/clavis"
)

type test struct {
	name string
	want interface{}
	got  interface{}
}

func setTestCases() []test {
	var tests []test

	got := clavis.Set("", "Token Bearer")
	tests = append(tests, test{
		name: "Missing key",
		want: fmt.Errorf("Missing key"),
		got:  got,
	})

	got = clavis.Set("token-123", "")
	tests = append(tests, test{
		name: "Missing Value",
		want: fmt.Errorf("Missing value"),
		got:  got,
	})

	got = clavis.Set("token-123", "Bearer Token")
	tests = append(tests, test{
		name: "Correct set",
		want: nil,
		got:  got,
	})

	return tests
}

func TestSet(t *testing.T) {
	for _, tt := range setTestCases() {
		if !reflect.DeepEqual(tt.got, tt.want) {
			t.Errorf("Want: %v\n Got: %v", tt.want, tt.got)
		}
	}
}
