package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("--- Day 1: The Tyranny of the Rocket Equation ---")
	input, err := readInputFromFile("input")
	if err != nil {
		return
	}

	totalFuel := 0
	for _, value := range input {
		mass, _ := strconv.Atoi(value)
		if mass <= 0 {
			continue
		}
		fmt.Printf("Mass: %d\n", mass)
		// fuel := calculateFuel(mass)
		fuel := calculatePart2Fuel(mass)
		totalFuel += fuel
		fmt.Printf("Fuel: %d\n", fuel)
		fmt.Printf("Total: %d\n", totalFuel)
		fmt.Println()
	}

	fmt.Println(totalFuel)
}

func calculatePart2Fuel(mass int) int {
	result := (mass / 3) - 2
	if result <= 0 {
		return 0
	}
	return result + calculatePart2Fuel(result)
}

func calculateFuel(mass int) int {
	result := (mass / 3) - 2
	return result
}

func readInputFromFile(fileName string) ([]string, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return []string{}, err
	}
	input := strings.Split(string(data), "\n")
	return input, nil
}
