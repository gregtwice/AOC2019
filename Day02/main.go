package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./Day02/puzzle.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var nums []int
	var i string
	for scanner.Scan() {
		i = scanner.Text()
	}
	strnums := strings.Split(i, ",")
	for _, value := range strnums {
		intval, _ := strconv.Atoi(value)
		nums = append(nums, intval)
	}

	for noun := 1; noun < 100; noun++ {
		for verb := 1; verb < 100; verb++ {
			c, p1, p2, addr := 0, 1, 2, 3
			numst := make([]int, len(nums))
			copy(numst, nums)
			numst[p1] = noun
			numst[p2] = verb
			for numst[c] != 99 {
				switch numst[c] {
				case 1:
					numst[numst[addr]] = numst[numst[p1]] + numst[numst[p2]]
					c += 4
					p1 += 4
					p2 += 4
					addr += 4
				case 2:
					numst[numst[addr]] = numst[numst[p1]] * numst[numst[p2]]
					c += 4
					p1 += 4
					p2 += 4
					addr += 4
				default:
					println("Erreur", numst[c], c, noun, verb)
					os.Exit(0)
				}

			}
			if numst[0] == 19690720 {
				println(100*noun, verb)
			}
		}
	}

}
