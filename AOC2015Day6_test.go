package main

import (
	"reflect"
	"testing"
)

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

func Test_Individual(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Test 1", args: args{"turn on 0,0 through 999,999"}, want: 1000000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var lightBoard LightBoard
			action, startx, starty, endx, endy := parseInput(tt.args.input)

			switch action {
			case "turn on":
				processTurnOn(startx, endx, starty, endy, lightBoard)
			case "turn off":
				processTurnOff(startx, endx, starty, endy, lightBoard)
			case "toggle":
				processToggle(startx, endx, starty, endy, lightBoard)
			}
			var total int = 0
			for _, row := range lightBoard {
				for _, val := range row {
					total += val
				}
			}
			if !reflect.DeepEqual(total, tt.want) {
				t.Errorf("total = %v, want %v", total, tt.want)
			}
		})
	}
}
