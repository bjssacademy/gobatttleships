package main

import (
	"fmt"
	"math/rand"
	"testing"
)

//you can run all you tests by typing
//go test -v
//in the terminal window

//this is a utility function for testing
//it will return a random square on the grid
//it does not keep track of any previously returned grids
func getRandomGridSquare() []int {
	
	row := []int{1,2,3,4,5,6,7,8,9}
	column := []int{1,2,3,4,5,6,7,8,9}

	return []int{rand.Intn(len(row))+1,rand.Intn(len(column))+1}

}

//these are the two tests we have for our functions in main
//the purpose of tests is to mimic interaction with our code
//there is no "user input" - the test is the calling code

//here is an example of a failing test - what do we need to do to fix it?
func TestCreateGrid(t *testing.T){
	grid := CreateGrid()
	if (len(grid) > 7){
		t.Error("Grid is too big! Expected max size of 7, got: ", len(grid))
	}
}

//one good place to start here is by using our utility function
//to target a random grid square rather than 1,1 co-ordinates every time

func TestPlayerOneTakingShot(t *testing.T) {
	grid := CreateGrid()
	shotResult := PlayerOneTurn(grid, []int{1,1})
	if (shotResult != false){
		t.Error("Shot should be false!")
	}
}

func TestPlayerTwoTakingShot(t *testing.T) {
	grid := CreateGrid()
	shotResult := PlayerTwoTurn(grid, []int{1,1})
	if (shotResult != true){
		t.Error("Shot should be true!")
	}
}

//other tests here that fail

//sometimes we write tests to test our own functions.
func TestGetRandomGridSquare(t *testing.T){
	gridSquare := getRandomGridSquare()
	
	//literally only exists here to show you the output
	//should not exist in a real test
	fmt.Println(gridSquare)

	//poor test making use of magic numbers
	//you should probably re-write it
	if (gridSquare[0] <= 0 || gridSquare[0] >= 10){
		t.Error("Grid square row should be >0 and <10, but got: ", gridSquare[0])
	}
	
	if (gridSquare[1] <= 0 || gridSquare[1] >= 10){
		t.Error("Grid square column should be >0 and <10, but got: ", gridSquare[1])
	}
}