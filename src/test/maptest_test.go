package maptest

import "testing"
import "reflect"
import "fmt"

func m_map(a string, b int, c string, d int) map[string]int {
	my_map := map[string]int{
		a: b,
		c: d,
	}

	fmt.Println("my_map: ", my_map)
	fmt.Println("type of my_map: ", reflect.TypeOf(my_map))

	return my_map
}

func TestMMap(t *testing.T) {
	type args struct {
		a string
		b int
		c string
		d int
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "normal",
			args: args{a: "x", b: 100, c:"y", d: 200},
			want: map[string]int{
				      "x": 100,
					  "y": 200,
				  },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := m_map(tt.args.a, tt.args.b, tt.args.c, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("m_map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ref_map(a string, b int) int {
	my_map := map[string]int{ a: b }

	return my_map[a]
}

func TestRefMap(t *testing.T) {
	type args struct {
		a string
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "reference map by key",
			args: args{a: "x", b: 100},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ref_map(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ref_map() = %v, want %v", got, tt.want)
			}
		})
	}
}
