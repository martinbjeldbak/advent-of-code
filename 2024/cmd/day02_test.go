package cmd

import "testing"

func Test_day2(t *testing.T) {
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
			name: "Day 2",
			args: args{
				inputData: []string{
					"7 6 4 2 1",
					"1 2 7 8 9",
					"9 7 6 2 1",
					"1 3 2 4 5",
					"8 6 4 4 1",
					"1 3 6 7 9",
				},
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "ex",
			args: args{
				inputData: []string{
					"9 13 14 18 21 24 24",
				},
			},
			want:    0,
			wantErr: false,
		},
		// 524 too high
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := day2(tt.args.inputData)
			if (err != nil) != tt.wantErr {
				t.Errorf("day2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("day2() = %v, want %v", got, tt.want)
			}
		})
	}
}
