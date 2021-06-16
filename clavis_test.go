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
			t.Errorf("\nWant: %v\n Got: %v", tt.want, tt.got)
		}
	}
}

func getTestCases() []test {
	var tests []test

	key := "key"
	val := "value"

	{
		got := clavis.Set(key, val)
		tests = append(tests, test{
			name: "Correct set",
			want: nil,
			got:  got,
		})
	}

	{
		_, gotErr := clavis.Get("")
		tests = append(tests, test{
			name: "Empty key",
			want: fmt.Errorf("Missing key"),
			got:  gotErr,
		})
	}

	{
		key := "random_key"
		_, gotErr := clavis.Get(key)
		tests = append(tests, test{
			name: "Key inexistent",
			want: fmt.Errorf("%s not found", key),
			got:  gotErr,
		})
	}

	{
		gotVal, gotErr := clavis.Get(key)

		tests = append(tests, test{
			name: "Found key-value nil error",
			want: nil,
			got:  gotErr,
		})

		tests = append(tests, test{
			name: "Found key-value nil error",
			want: val,
			got:  gotVal,
		})
	}

	return tests
}

func TestGet(t *testing.T) {
	for _, tt := range getTestCases() {
		if !reflect.DeepEqual(tt.got, tt.want) {
			t.Errorf("\nWant: %v\n Got: %v", tt.want, tt.got)
		}
	}
}
