package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./Day03/in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	wire1 := []string{}
	wire2 := []string{}
	i := 0
	for scanner.Scan() {
		s := scanner.Text()
		if i == 0 {
			wire1 = strings.Split(s, ",")
			i++
		} else if i == 1 {
			wire2 = strings.Split(s, ",")
			i++
		}
	}
	fmt.Println(wire1)
	fmt.Println(wire2)

	minD := 100000000
	Ax := 0
	Ay := 0
	Cx := 0
	Cy := 0
	bestStep := 100000000
	stepsw1 := 0
	stepsw2 := 0
	for _, value := range wire1 {
		sd := value[1:]
		d, _ := strconv.Atoi(sd)
		var Bx, By int
		switch value[0] {
		case 'R': //right
			Bx = Ax + d
			By = Ay
		case 'D': // down
			Bx = Ax
			By = Ay - d
		case 'U': // up
			Bx = Ax
			By = Ay + d
		case 'L': //left
			Bx = Ax - d
			By = Ay
		}
		for _, v := range wire2 {
			sd := v[1:]
			d, _ := strconv.Atoi(sd)
			var x, y int
			var Dx, Dy int
			switch v[0] {
			case 'R': //right
				Dx = Cx + d
				Dy = Cy
				x = Ax
				y = Cy
				if value[0] == 'U' || value[0] == 'D' {
					if (Ay-y)*(By-y) < 0 && (Cx-x)*(Dx-x) < 0 {
						minD = min(minD, x, y)
						bestStep = int(math.Min(float64(stepsw1+stepsw2+int(math.Abs(float64(Cx-x))+math.Abs(float64(Ay-y)))), float64(bestStep)))
					}
				}
			case 'D': // down
				Dx = Cx
				Dy = Cy - d
				x = Cx
				y = Ay
				if value[0] == 'L' || value[0] == 'R' {
					if (Cy-y)*(Dy-y) < 0 && (Ax-x)*(Bx-x) < 0 {
						minD = min(minD, x, y)
						bestStep = int(math.Min(float64(stepsw1+stepsw2+int(math.Abs(float64(Ax-x))+math.Abs(float64(Cy-y)))), float64(bestStep)))

					}
				}
			case 'U': // up
				Dx = Cx
				Dy = Cy + d
				x = Cx
				y = Ay
				if value[0] == 'L' || value[0] == 'R' {
					if (Cy-y)*(Dy-y) < 0 && (Ax-x)*(Bx-x) < 0 {
						minD = min(minD, x, y)
						bestStep = int(math.Min(float64(stepsw1+stepsw2+int(math.Abs(float64(Ax-x))+math.Abs(float64(Cy-y)))), float64(bestStep)))
					}
				}
			case 'L': //left
				Dx = Cx - d
				Dy = Cy
				x = Ax
				y = Cy
				if value[0] == 'U' || value[0] == 'D' {
					if (Ay-y)*(By-y) < 0 && (Cx-x)*(Dx-x) < 0 {
						minD = min(minD, x, y)
						bestStep = int(math.Min(float64(stepsw1+stepsw2+int(math.Abs(float64(Cx-x))+math.Abs(float64(Ay-y)))), float64(bestStep)))
					}
				}
			}
			stepsw2 += d
			Cx = Dx
			Cy = Dy
		}
		stepsw2 = 0
		stepsw1 += d
		Cx = 0
		Cy = 0
		Ax = Bx
		Ay = By
	}
	fmt.Println(minD)
	fmt.Println(bestStep)
}

func min(minD int, x int, y int) int {
	return int(math.Min(float64(minD), math.Abs(float64(x))+math.Abs(float64(y))))
}
