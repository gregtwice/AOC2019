package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	file, err := os.Open("./Day01/puzzle.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	s1 := 0
	s2 := 0
	t := time.Now()
	for scanner.Scan() {
		mass := scanner.Text()
		m, _ := strconv.Atoi(mass)
		m = m/3 - 2
		s1 += m
		s2 += m
		for m >= 0 {
			m = m/3 - 2
			if m > 0 {
				s2 += m
			}
		}
	}
	t2 := time.Now().Sub(t)
	fmt.Println("Part 1: ", s1)
	fmt.Println("Part 2: ", s2)
	println(t2.Seconds())

}
