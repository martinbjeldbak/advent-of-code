package cmd

import "testing"

func Test_day7part2(t *testing.T) {
	type args struct {
		inputData []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example",
			args: args{
				inputData: []string{
					"32T3K 765",
					"T55J5 684",
					"KK677 28",
					"KTJJT 220",
					"QQQJA 483",
				},
			},
			want: 5905,
		},
		{
			name: "Martin example",
			args: args{
				inputData: []string{
					"AJAJA 165",
					"AJAAA 146",
				},
			},
			want: 457,
		},
		{
			name: "Martin example 2",
			args: args{
				inputData: []string{
					"222J2 165",
					"22222 146",
				},
			},
			want: 457,
		},
		{
			name: "Martin example 3",
			args: args{
				inputData: []string{
					"222J2 146",
					"JJJJJ 165",
				},
			},
			want: 457,
		},
		{
			name: "Martin example 4",
			args: args{
				inputData: []string{
					"333J3 146",
					"JJJJJ 165",
				},
			},
			want: 457,
		},
		// {
		// 	name: "Martin example 5",
		// 	args: args{
		// 		inputData: []string{
		// 			" 146",
		// 			" 165",
		// 		},
		// 	},
		// 	want: 457,
		// },
		{
			name: "Reddit example 1",
			args: args{
				[]string{
					"2345A 1",
					"Q2KJJ 13",
					"Q2Q2Q 19",
					"T3T3J 17",
					"T3Q33 11",
					"2345J 3",
					"J345A 2",
					"32T3K 5",
					"T55J5 29",
					"KK677 7",
					"KTJJT 34",
					"QQQJA 31",
					"JJJJJ 37",
					"JAAAA 43",
					"AAAAJ 59",
					"AAAAA 61",
					"2AAAA 23",
					"2JJJJ 53",
					"JJJJ2 41",
				},
			},
			want: 6839,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day7part2(tt.args.inputData); got != tt.want {
				t.Errorf("day7part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
