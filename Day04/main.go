package main

import "fmt"

func isValidpart1(mdp int) bool {

	if mdp > max || mdp < min {
		return false
	}

	arrMdp := []int{
		(mdp / 100000) % 10,
		(mdp / 10000) % 10,
		(mdp / 1000) % 10,
		(mdp / 100) % 10,
		(mdp / 10) % 10,
		(mdp) % 10,
	}

	flag := false
	for i := 0; i < 5; i++ {
		if arrMdp[i] > arrMdp[i+1] {
			return false
		}
		if arrMdp[i] == arrMdp[i+1] {
			flag = true
		}
	}

	return flag
}

func isValidpart2(mdp int) bool {

	if mdp > max || mdp < min {
		return false
	}

	arrMdp := []int{
		(mdp / 100000) % 10,
		(mdp / 10000) % 10,
		(mdp / 1000) % 10,
		(mdp / 100) % 10,
		(mdp / 10) % 10,
		(mdp) % 10,
	}

	flag := false
	digitCount := make([]int, 10)
	for i := 0; i < 5; i++ {
		if arrMdp[i] > arrMdp[i+1] {
			return false
		}
		if arrMdp[i] == arrMdp[i+1] {
			flag = true
		}
		digitCount[arrMdp[i]] += 1
	}
	digitCount[arrMdp[5]]++

	flagEx2 := false
	for _, value := range digitCount {
		if value == 2 {
			flagEx2 = true
		}
	}

	return flag && flagEx2
}

var max int = 624574

var min int = 158126

func main() {

	sumP1 := 0
	sumP2 := 0
	for i := min; i <= max; i++ {
		if isValidpart1(i) {
			sumP1++
		}

		if isValidpart2(i) {
			sumP2++
		}
	}
	fmt.Println("Part1", sumP1)
	fmt.Println("Part2", sumP2)
}

func test() {
	fmt.Println(isValidpart2(112233))
	// ok
	fmt.Println(isValidpart2(123444))
	// notok
	fmt.Println(isValidpart2(111122))
	// ok
	fmt.Println(isValidpart2(114444))
	// notok
}
