package cmd

import "testing"

func Test_day10(t *testing.T) {
	type args struct {
		inputData []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "Example simple",
		// 	args: args{
		// 		inputData: []string{
		// 			".....",
		// 			".S-7.",
		// 			".|.|.",
		// 			".L-J.",
		// 			".....",
		// 		},
		// 	},
		// 	want: 4,
		// },
		// {
		// 	name: "Example",
		// 	args: args{
		// 		inputData: []string{
		// 			"..F7.",
		// 			".FJ|.",
		// 			"SJ.L7",
		// 			"|F--J",
		// 			"LJ...",
		// 		},
		// 	},
		// 	want: 8,
		// },
		{
			name: "Example 2",
			args: args{
				inputData: []string{
					"-L|F7",
					"7S-7|",
					"L|7||",
					"-L-J|",
					"L|-JF",
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day10(tt.args.inputData); got != tt.want {
				t.Errorf("day10() = %v, want %v", got, tt.want)
			}
		})
	}
}
