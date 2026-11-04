package main

import (

	"fmt"
	"os"
	"os/exec"
//	"time"
	"log"

)

func spawn_barrel(grid [5][4]int) [5][4]int{

	var barrels_spawned int

	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			if grid[i][j] == 0 && barrels_spawned <= 0 {
				grid[i][j] = 1;
				barrels_spawned ++
			}
		}
	}

	return grid
}

func update_grid(grid [5][4]int) {
	for i := 0; i < 5; i++ {
		fmt.Println(grid[i])
	}
}

func clear_screen() {
	// -- makes clean value of the command to execute -- //
	clean := exec.Command("clear")
	// -- makes sure that user can see whatever is the command outputs -- //
	clean.Stdout = os.Stdout
	// -- runs the command and checks for error -- //
	err := clean.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	var grid [5][4]int

	clear_screen()

	// -- how to use spawning function -- //
	//grid = spawn_barrel(grid)

	for {
		clear_screen()
		grid = spawn_barrel(grid)
		update_grid(grid)

	}

}
