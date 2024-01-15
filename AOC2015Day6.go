package main

type Light struct {
	x int
	y int
}

type LightBoard [1000][1000]int

type action int

const (
	turnOn action = iota
	turnOff
	toggle
)

// calculateLight calculates the number of lights that are on.
func calculateLight(beginLight Light, endLight Light) int {
	return (endLight.x + 1 - beginLight.x) * (endLight.y + 1 - beginLight.y)
}

func parseInput(input string) (Light, Light, string) {
	var beginLight Light
	var endLight Light
	var action string

	//parse input
	//return beginLight, endLight, action
	return beginLight, endLight, action
}

func main() {

	//open file
	//read file
	//iterate over each line
	//parse line

	lightBoard := LightBoard{}

}
