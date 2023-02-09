package funfacts

import (
	"reflect"
	"testing"
)

func TestGetFunFacts(t *testing.T) {
	type test struct {
		input string
		want  string
	}


	for _, tc := range tests {
		got := GetFunFacts(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected %v, got %v" tc.want, got)
		}
	}
}
