package fib

import (
	"testing"
)

func TestFib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		arg  args
		want int
	}{
		{
			name: "Case 1",
			arg: args{
				n: 7,
			},
			want: 13,
		},
		{
			name: "Case 2",
			arg: args{
				n: 100,
			},
			want: 3736710778780434371,
		},
		{
			name: "Case 2",
			arg: args{
				n: 100000,
			},
			want: 2754320626097736315,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fib(tt.arg.n); got != tt.want {
				t.Errorf("Fib wont: %v got: %v", got, tt.want)
			}
		})
	}
}
