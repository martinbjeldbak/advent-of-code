package cmd

import "testing"

func Test_day03(t *testing.T) {
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
			name:    "Puzzle input",
			args:    args{
				inputData: []string{
					"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
				},
			},
			want:    161,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := day03(tt.args.inputData)
			if (err != nil) != tt.wantErr {
				t.Errorf("day03() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("day03() = %v, want %v", got, tt.want)
			}
		})
	}
}
