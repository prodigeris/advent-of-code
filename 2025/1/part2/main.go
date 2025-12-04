package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	pos := 50
	password := 0

	for scanner.Scan() {

		line := scanner.Text()
		rotation := line[0]
		times, _ := strconv.Atoi(line[1:])
		for i := 0; i < int(math.Abs(float64(times))); i++ {
			if rotation == 'L' {
				pos--
			} else {
				pos++
			}
			if pos > 99 {
				pos = pos - 100
			} else if pos < 0 {
				pos = pos + 100
			}
			if pos == 0 {
				password++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	fmt.Println("password is", password)
}
