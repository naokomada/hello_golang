package multi

import "testing"
// import "fmt"

func multi(a, b int) int {
	// fmt.Print("multi")
	return a * b
}

func TestMulti(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal",
			args: args{a: 2, b: 2},
			want: 4,
		},
		{
			name: "normal2",
			args: args{a: 4, b: 2},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multi(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("multi() = %v, want %v", got, tt.want)
			}
		})
	}
}
