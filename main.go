package main

/*
This game of battleships is very simple to start:
There are 2 players
Each player has a grid which is 7*7
Each player has 9 Battleships, each of which can occupy only one square on their grid
Each player can place their battleships anywhere on this grid
Players take it in turns to pick any grid square reference
If the player hits a battleship, then it is sunk, and the turn passes to the opponent
If the player misses a battleship then it is called a miss, and the turn passes to the opponent
The player to first sink all their opponent's battleships is the winner
*/

//All code in here is example code, you do not have to keep any of it.

func PlayerOneTurn(playerTwoGrid [8][8]string, shotCoordinates []int) (shotStatus bool) {
	return false //shot missed
}

func PlayerTwoTurn(playerOneGrid [8][8]string, shotCoordinates []int) (shotStatus bool) {
	return true //shot hit
}

func CreateGrid() (grid [8][8]string) {
	//this is a fixed array, not a slice
	return [8][8]string{}
}
