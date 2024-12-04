package cmd

import "testing"

func Test_day04(t *testing.T) {
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
					"SAMX",
				},
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Example input 1",
			args: args{
				inputData: []string{
					".AMX",
					".AMX",
					"SAMX",
					"SAMX",
				},
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "Example input 1",
			args: args{
				inputData: []string{
					"XMAS",
				},
			},
			want:    1,
			wantErr: false,
		},
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
			want:    18,
			wantErr: false,
		},
		{
			name: "Example input 2",
			args: args{
				inputData: []string{
					"..X...",
					".SAMX.",
					".A..A.",
					"XMAS.S",
					".X....",
				},
			},
			want:    4,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := day04(tt.args.inputData)
			if (err != nil) != tt.wantErr {
				t.Errorf("day04() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("day04() = %v, want %v", got, tt.want)
			}
		})
	}
}
