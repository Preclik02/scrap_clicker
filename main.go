package main

import (

	"fmt"

)

func spawn_barrel(grid [5][4]int) [5][4]int{

	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			if grid[i][j] == 0 {
				grid[i][j] = 1;
			}
		}
	}

	return grid
}

func main() {

	var grid [5][4]int

	for i := 0; i < 5; i++ {
		fmt.Println(grid[i])
	}

	grid = spawn_barrel(grid)

	fmt.Println("--- spawn_barrel ---")

	for i := 0; i < 5; i++ {
		fmt.Println(grid[i])
	}

}
