package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./Day05/in")
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
	reader := bufio.NewReader(os.Stdin)
	strnums := strings.Split(i, ",")
	for _, value := range strnums {
		intval, _ := strconv.Atoi(value)
		nums = append(nums, intval)
	}

	cursor := 0
	numst := make([]int, len(nums))
	copy(numst, nums)
	for numst[cursor] != 99 {
		switch numst[cursor] % 100 {
		case 1: // Add
			str := fmt.Sprintf("%05d", numst[cursor])
			var param1, param2 int
			if str[1] == '1' {
				param2 = numst[cursor+2]
			} else {
				param2 = numst[numst[cursor+2]]
			}
			if str[2] == '1' {
				param1 = numst[cursor+1]
			} else {
				param1 = numst[numst[cursor+1]]
			}
			numst[numst[cursor+3]] = param1 + param2
			cursor += 4
		case 2: // Mult
			str := fmt.Sprintf("%05d", numst[cursor])
			var param1, param2 int
			if str[1] == '1' {
				param2 = numst[cursor+2]
			} else {
				param2 = numst[numst[cursor+2]]
			}
			if str[2] == '1' {
				param1 = numst[cursor+1]
			} else {
				param1 = numst[numst[cursor+1]]
			}
			numst[numst[cursor+3]] = param1 * param2

			cursor += 4
		case 3:
			in, _ := reader.ReadString('\n')
			in = in[:len(in)-1]
			i, _ := strconv.Atoi(in)
			numst[numst[cursor+1]] = i
			cursor += 2
		case 4:
			str := fmt.Sprintf("%03d", numst[cursor])
			var param int
			if str[0] == '1' {
				param = numst[cursor+1]
			} else {
				param = numst[numst[cursor+1]]
			}
			fmt.Println(param)
			cursor += 2
		case 5:

			str := fmt.Sprintf("%04d", numst[cursor])
			var param1, param2 int
			if str[0] == '1' {
				param2 = numst[cursor+2]
			} else {
				param2 = numst[numst[cursor+2]]
			}
			if str[1] == '1' {
				param1 = numst[cursor+1]
			} else {
				param1 = numst[numst[cursor+1]]
			}

			if param1 != 0 {
				cursor = param2
			} else {
				cursor += 3
			}
		case 6:
			str := fmt.Sprintf("%04d", numst[cursor])
			var param1, param2 int
			if str[0] == '1' {
				param2 = numst[cursor+2]
			} else {
				param2 = numst[numst[cursor+2]]
			}
			if str[1] == '1' {
				param1 = numst[cursor+1]
			} else {
				param1 = numst[numst[cursor+1]]
			}
			if param1 == 0 {
				cursor = param2
			} else {
				cursor += 3
			}
		case 7:
			str := fmt.Sprintf("%05d", numst[cursor])
			var param1, param2 int
			if str[1] == '1' {
				param2 = numst[cursor+2]
			} else {
				param2 = numst[numst[cursor+2]]
			}
			if str[2] == '1' {
				param1 = numst[cursor+1]
			} else {
				param1 = numst[numst[cursor+1]]
			}
			if param1 < param2 {
				if str[0] == '1' {
					numst[cursor+3] = 1
				} else {
					numst[numst[cursor+3]] = 1
				}

			} else {
				if str[0] == '1' {
					numst[cursor+3] = 0
				} else {
					numst[numst[cursor+3]] = 0
				}
			}
			cursor += 4
		case 8:
			str := fmt.Sprintf("%05d", numst[cursor])
			var param1, param2 int
			if str[1] == '1' {
				param2 = numst[cursor+2]
			} else {
				param2 = numst[numst[cursor+2]]
			}
			if str[2] == '1' {
				param1 = numst[cursor+1]
			} else {
				param1 = numst[numst[cursor+1]]
			}
			if param1 == param2 {
				if str[0] == '1' {
					numst[cursor+3] = 1
				} else {
					numst[numst[cursor+3]] = 1
				}

			} else {
				if str[0] == '1' {
					numst[cursor+3] = 0
				} else {
					numst[numst[cursor+3]] = 0
				}
			}
			cursor += 4
		default:
			fmt.Println("Error instruction", numst[cursor], "unknown")
			os.Exit(0)
		}
	}

}
