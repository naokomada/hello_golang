package slicetest

import "testing"
import "reflect"
import "fmt"

func m_slice(a, b, c string) []string {
	fruits := []string{a, b}
	fruits = append(fruits, c)
    fmt.Println("fruits: ", fruits)
	fmt.Println("type of fruits: ", reflect.TypeOf(fruits))
	fmt.Println("length: ", len(fruits), "capacity: ", cap(fruits))
	return fruits
}

func TestMSlice(t *testing.T) {
	type args struct {
		a string
		b string
		c string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "normal",
			args: args{a: "apple", b: "orange", c:"lemon"},
			want: []string{"apple", "orange", "lemon"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := m_slice(tt.args.a, tt.args.b, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("m_slice() = %v, want %v", got, tt.want)
			}
		})
	}
}
