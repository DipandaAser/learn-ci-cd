package learn_ci_cd

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Sum of no value",
			args: args{[]int{}},
			want: 0,
		},
		{
			name: "Sum of one value",
			args: args{[]int{5}},
			want: 5,
		},
		{
			name: "Sum of multiples values",
			args: args{[]int{5, 5, 5, 15}},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.numbers...); got != tt.want {
				t.Errorf("Addition() = %v, want %v", got, tt.want)
			}
		})
	}
}
