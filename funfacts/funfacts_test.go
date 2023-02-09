package funfacts

import (
	"reflect"
	"testing"
)

func TestGetFunFacts(t *testing.T) {
	type test struct {
		input //må skrive inn riktig type for input
		want  //må skrive inn riktig type for return verdi
	}


	for _, tc := range tests {
		got := GetFunFacts(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected %v, got %v" tc.want, got)
		}
	}
}
