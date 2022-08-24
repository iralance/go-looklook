package nonrepeatingsubstr

import "testing"

func Test_lengthOfNonRepeatingSubstr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"demo1",
			args{"abcabcbb"},
			3,
		},
		{
			"demo2",
			args{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"},
			8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfNonRepeatingSubstr(tt.args.s); got != tt.want {
				t.Errorf("lengthOfNonRepeatingSubstr() = %v, want %v", got, tt.want)
			}
		})
	}
}
