package cmd

import "testing"

func Test_day1(t *testing.T) {
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
			name:    "Day 1 example",
			args:    args{
				inputData: []string{
					"1abc2",
					"pqr3stu8vwx",
					"a1b2c3d4e5f",
					"treb7uchet",
				},
			},
			want:    142,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := day1(tt.args.inputData)
			if (err != nil) != tt.wantErr {
				t.Errorf("day1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("day1() = %v, want %v", got, tt.want)
			}
		})
	}
}
