package main

import (
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/alecthomas/assert"
)

func TestDay13_1(t *testing.T) {

	str := `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`

	// 	str = `Button A: X+19, Y+95
	// Button B: X+59, Y+23
	// Prize: X=6550, Y=10174

	// Button A: X+20, Y+58
	// Button B: X+60, Y+11
	// Prize: X=3140, Y=4691

	// Button A: X+92, Y+31
	// Button B: X+49, Y+87
	// Prize: X=6599, Y=2717

	// Button A: X+46, Y+72
	// Button B: X+29, Y+12
	// Prize: X=18754, Y=13304`

	str, err := loadInputFromFile("13")
	assert.NoError(t, err)

	lines := strings.Split(str, "\n\n")
	tokensA := 3
	tokensB := 1

	sum := 0

	for _, block := range lines {
		var xA, yA, xB, yB, prizeX, prizeY int

		// Use a more lenient format for names with spaces
		_, err := fmt.Sscanf(block,
			"Button A: X+%d, Y+%d\nButton B: X+%d, Y+%d\nPrize: X=%d, Y=%d",
			&xA, &yA, &xB, &yB, &prizeX, &prizeY)
		if err != nil {
			fmt.Println("Error parsing block:", err)
			continue
		}
		// numA, numB := job{
		// 	buttonA: item{xA, yA},
		// 	buttonB: item{xB, yB},
		// 	prize:   item{prizeX, prizeY},
		// }.process()

		// value := (numA * tokensA) + (numB * tokensB)

		fmt.Println("__________")

		value := job{
			buttonA: item{xA, yA},
			buttonB: item{xB, yB},
			prize:   item{prizeX, prizeY},
		}.process3(tokensA, tokensB)

		sum += value
	}

	fmt.Println("sum: ", sum)
	assert.Equal(t, 480, sum)

	// 38839

}

func (j job) process3(multiA, multiB int) int {

	numA := 0
	numB := 0
	value := 0

	for i := 0; i < 100; i++ {

		subAx := i * j.buttonA.x
		subAy := i * j.buttonA.y

		if subAx > j.prize.x || subAy > j.prize.y {
			break
		}

		modBx := math.Mod(float64(j.prize.x-subAx), float64(j.buttonB.x))
		modBy := math.Mod(float64(j.prize.y-subAy), float64(j.buttonB.y))

		if modBx == 0 && modBy == 0 {

			fmt.Println("valid-A", modBy, i)

			numA = i
			numB = (j.prize.x - subAx) / j.buttonB.x
			numBy := (j.prize.y - subAy) / j.buttonB.y

			if numB == numBy {

				v := (numA * multiA) + (numB * multiB)

				fmt.Println("v", v)
				if value > v || value == 0 {
					value = v
					fmt.Println("better value-A", numA, numB, value, i)
				}
			}
		}

	}

	if value == 0 {
		fmt.Println("not valid")
	}
	return value
}

func (j job) process2(multiA, multiB int) int {

	numA := 0
	numB := 0

	startPrize := j.prize
	value := 0

	for i := 0; i <= 100; i++ {
		if math.Mod(float64(j.prize.x), float64(j.buttonB.x)) == 0 {
			if math.Mod(float64(j.prize.y), float64(j.buttonB.y)) == 0 {

				numA = i
				numB = (startPrize.x - (j.buttonA.x * i)) / j.buttonB.x
				v := (numA * multiA) + (numB * multiB)

				if value == 0 {
					value = v
				}
				if v < value {
					value = v
					fmt.Println("better value", numA, numB, value, i)
				}
				fmt.Println("valid Y", numA, numB, value, i)
			}
		}
		j.prize.x -= j.buttonA.x
		j.prize.y -= j.buttonA.y

		if j.prize.x < 0 || j.prize.y < 0 {
			break
		}
	}

	return value
}

// func (j job) process() (int, int) {
// 	fmt.Println("Processing job:", j)

// 	numA := 0
// 	numB := 0
// 	for {
// 		// Calculate the distance from the prize to each button
// 		distA := distance(j.prize, j.buttonA)
// 		distB := distance(j.prize, j.buttonB)

// 		if distA < distB {
// 			j.prize.x -= j.buttonA.x
// 			j.prize.y -= j.buttonA.y
// 			numA++
// 		} else {
// 			j.prize.x -= j.buttonB.x
// 			j.prize.y -= j.buttonB.y
// 			numB++
// 		}

// 		if j.prize.x == 0 && j.prize.y == 0 {
// 			fmt.Println("valid")
// 			break
// 		}

// 		if j.prize.x < 0 || j.prize.y < 0 {
// 			fmt.Println("not valid")
// 			numA = 0
// 			numB = 0
// 			break
// 		}

// 	}
// 	return numA, numB
// }

// func distance(a, b item) int {
// 	return int(math.Sqrt(
// 		float64((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y)),
// 	))
// }

type job struct {
	buttonA item
	buttonB item
	prize   item
}

type item struct {
	x int
	y int
}
