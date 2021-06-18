package clavis_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/fiuskylab/clavis"
)

type test struct {
	name string
	want interface{}
	got  interface{}
}

func setTestCases() []test {
	var tests []test

	client := clavis.NewClavis(clavis.DefaultConfig())

	got := client.Set("", "Token Bearer", -1)
	tests = append(tests, test{
		name: "Missing key",
		want: fmt.Errorf("Missing key"),
		got:  got,
	})

	got = client.Set("token-123", "", -1)
	tests = append(tests, test{
		name: "Missing Value",
		want: fmt.Errorf("Missing value"),
		got:  got,
	})

	got = client.Set("token-123", "Bearer Token", -1)
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
			t.Run(tt.name, func(t *testing.T) {
				t.Errorf("\nWant: %v\n Got: %v", tt.want, tt.got)
			})
		}
	}
}

func getTestCases() []test {
	var tests []test

	key_inf := "key-infinite"
	val := "value"

	client := clavis.NewClavis(clavis.DefaultConfig())

	{
		got := client.Set(key_inf, val, -1)
		tests = append(tests, test{
			name: "Correct set",
			want: nil,
			got:  got,
		})
	}

	key_fin := "key-finite"
	{
		got := client.Set(key_fin, val, 1)
		time.Sleep(time.Second * 1)
		tests = append(tests, test{
			name: "Correct set",
			want: nil,
			got:  got,
		})
	}

	{
		_, gotErr := client.Get("")
		tests = append(tests, test{
			name: "Empty key",
			want: fmt.Errorf("Missing key"),
			got:  gotErr,
		})
	}

	{
		key := "random_key"
		_, gotErr := client.Get(key)
		tests = append(tests, test{
			name: "Key inexistent",
			want: fmt.Errorf("%s not found", key),
			got:  gotErr,
		})
	}

	{
		gotVal, gotErr := client.Get(key_inf)

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

	{
		_, gotErr := client.Get(key_fin)

		tests = append(tests, test{
			name: "Found expired key-value",
			want: fmt.Errorf("%s is expirated", key_fin),
			got:  gotErr,
		})
	}

	return tests
}

func TestGet(t *testing.T) {
	for _, tt := range getTestCases() {
		if !reflect.DeepEqual(tt.got, tt.want) {
			t.Run(tt.name, func(t *testing.T) {
				t.Errorf("\nWant: %v\n Got: %v", tt.want, tt.got)
			})
		}
	}
}
