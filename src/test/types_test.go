package types

import (
	"testing"
	"reflect"
)

func integers() reflect.Kind {
	var x int = 10
	return reflect.TypeOf(x).Kind()
}

func TestIntegers(t *testing.T) {
	tests := []struct {
		name string
		want reflect.Kind
	}{
		{
			name: "normal",
			want: reflect.Int,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := integers(); got != tt.want {
				t.Errorf("integers() = %v, want %v", got, tt.want)
			}
		})
	}
}
