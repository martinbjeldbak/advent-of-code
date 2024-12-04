package cmd

import "testing"

func Test_day04part2(t *testing.T) {
	type args struct {
		inputData []string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Example input 1",
			args: args{
				inputData: []string{
					"MMMSXXMASM",
					"MSAMXMSMSA",
					"AMXSXMAAMM",
					"MSAMASMSMX",
					"XMASAMXAMM",
					"XXAMMXXAMA",
					"SMSMSASXSS",
					"SAXAMASAAA",
					"MAMMMXMMMM",
					"MXMXAXMASX",
				},
			},
			want:    9,
			wantErr: false,
		},
		{
			name: "Simple input 1",
			args: args{
				inputData: []string{
					"M.S",
					".A.",
					"M.S",
				},
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := day04part2(tt.args.inputData)
			if (err != nil) != tt.wantErr {
				t.Errorf("day04part2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("day04part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
