package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, _ := os.Open("calories.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	maxCaloriesElf1 := 0
	maxCaloriesElf2 := 0
	maxCaloriesElf3 := 0

	currentElfCalories := 0

	for sc.Scan() {
		snack, err := strconv.Atoi(sc.Text())

		currentElfCalories += snack

		if err != nil {
			if currentElfCalories > maxCaloriesElf3 {
				maxCaloriesElf3 = currentElfCalories
			}
			if maxCaloriesElf3 > maxCaloriesElf2 {
				maxCaloriesElf3, maxCaloriesElf2 = maxCaloriesElf2, maxCaloriesElf3
			}
			if maxCaloriesElf2 > maxCaloriesElf1 {
				maxCaloriesElf2, maxCaloriesElf1 = maxCaloriesElf1, maxCaloriesElf2
			}
			//I start with a new elf
			currentElfCalories = 0
		}
	}
	fmt.Println(maxCaloriesElf1 + maxCaloriesElf2 + maxCaloriesElf3)
}
