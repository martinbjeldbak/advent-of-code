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
			name:    "Example data",
			args:    args{
				inputData: []string{
					"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
					"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
					"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
					"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
					"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
				},
			},
			want:    8,
			wantErr: false,
		},
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
