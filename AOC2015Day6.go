package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Light struct {
	x int
	y int
}

var myRegexp = regexp.MustCompile(`(?P<action>turn on|turn off|toggle) (?P<startx>\d+),(?P<starty>\d+) through (?P<endx>\d+),(?P<endy>\d+)`)

func p(s string) {
	fmt.Println(s)
}

type LightBoard [1000][1000]int

// calculateLight calculates the number of lights that are on.
func calculateLight(beginLight Light, endLight Light) int {
	return (endLight.x + 1 - beginLight.x) * (endLight.y + 1 - beginLight.y)
}

func parseInput(input string) (string, int, int, int, int) {
	match := myRegexp.FindStringSubmatch(input)
	result := make(map[string]string)
	for i, name := range myRegexp.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	fmt.Printf("Action: %s\n", result["action"])
	fmt.Printf("StartX: %s\n", result["startx"])
	fmt.Printf("StartY: %s\n", result["starty"])
	fmt.Printf("EndX: %s\n", result["endx"])
	fmt.Printf("EndY: %s\n", result["endy"])

	startx, _ := strconv.Atoi(result["startx"])
	starty, _ := strconv.Atoi(result["starty"])
	endx, _ := strconv.Atoi(result["endx"])
	endy, _ := strconv.Atoi(result["endy"])

	return result["action"], startx, starty, endx, endy
}
func processTurnOn(startx int, endx int, starty int, endy int, lb LightBoard) {
	for i := startx; i <= endx; i++ {
		for j := starty; j <= endy; j++ {
			lb[i][j] = 1
		}
	}
}
func processTurnOff(startx int, endx int, starty int, endy int, lightBoard LightBoard) {
	for i := startx; i <= endx; i++ {
		for j := starty; j <= endy; j++ {
			lightBoard[i][j] = 0
		}
	}
}
func processToggle(startx int, endx int, starty int, endy int, lightBoard LightBoard) {
	for i := startx; i <= endx; i++ {
		for j := starty; j <= endy; j++ {
			if lightBoard[i][j] == 1 {
				lightBoard[i][j] = 0
			} else {
				lightBoard[i][j] = 1
			}
		}
	}
}

func main() {
	lightBoard := LightBoard{}

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	lightBoard[0][0] = 1

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Printf("Processing: %s\n", scanner.Text())
		action, startx, starty, endx, endy := parseInput(scanner.Text())

		if action == "turn on" {
			processTurnOn(startx, endx, starty, endy, lightBoard)
		}
		if action == "turn off" {
			processTurnOff(startx, endx, starty, endy, lightBoard)
		}
		if action == "toggle" {
			processToggle(startx, endx, starty, endy, lightBoard)
		}
	}
	//print out total
	var total int = 0
	for _, row := range lightBoard {
		for _, val := range row {
			total += val
		}
	}
	fmt.Printf("Total: %d\n", total)
}
