package main

import "testing"

func Test_calculateLight(t *testing.T) {
	type args struct {
		beginLight Light
		endLight   Light
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1, 0, 0, 999, 999", args{Light{0, 0}, Light{999, 999}}, 1000000},
		{"Test 2, 0, 0, 999, 0", args{Light{0, 0}, Light{999, 0}}, 1000},
		{"Test 3, 499, 499, 500, 500", args{Light{499, 499}, Light{500, 500}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateLight(tt.args.beginLight, tt.args.endLight); got != tt.want {
				t.Errorf("calculateLight() = %v, want %v", got, tt.want)
			}
		})
	}
}
