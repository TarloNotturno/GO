package HilbertMaze

import (
	"fmt"
)

/* HilbertMaze https://app.codility.com/programmers/trainings/2/hilbert_maze/
A halfling is searching for treasure hidden in a maze in the dwarfs' mine. He has a map of the maze and would like to find the shortest path to the treasure.

The maze has a specific shape. It is placed on a square grid with M2 cells, where M = 2N+1+1 for some given size N. Each cell has coordinates (x, y), where 0 ≤ x, y < M, and can either be empty or contain a rock.

The mazes of sizes N = 1 and N = 2 are presented in the pictures below:



A maze of size N is constructed recursively from the layout of the maze of size N−1 (like the Hilbert curve). It contains four mazes of size N−1, each maze in one quarter. The maze in the bottom-left quarter is rotated by 90 degrees clockwise and the maze in the bottom-right quarter is rotated by 90 degrees counter-clockwise. The mazes in the top quarters are not rotated. There are three additional rocks (squares marked in green in the picture below) in the areas where the mazes intersect. The construction of the maze of size N = 3 is shown below:



The halfling would like to reach the treasure in the smallest number of steps possible. At each step, he can move to any one of the four adjacent cells (north, south, west, east) that does not contain a rock and is not outside of the grid.

For example, given N = 1, the halfling needs 8 steps to move from cell (2, 1) to cell (3, 4):



Write a function:

func Solution(N int, A int, B int, C int, D int) int

that, given the size of the maze N, coordinates of the halfling (A, B) and coordinates of the treasure (C, D), returns the minimum number of steps required for the halfling to reach the treasure.

Examples
Given N = 1, A = 2, B = 1, C = 3 and D = 4 the function should return 8, as shown above.

Given N = 2, A = 2, B = 5, C = 6 and D = 6 the function should return 7:



Given N = 3, A = 6, B = 6, C = 10 and D = 13 the function should return 39:



Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..25];
A, B, C, D are integers within the range [0..2N+1];
cells (A, B) and (C, D) do not contain rocks;
the result will be an integer smaller than 2,147,483,647
*/

type posMap struct {
	x int
	y int
}

type maze struct {
	// map, value 1 is equal to a wall, 2 a node visited and 3 the treasure
	mazeMap map[posMap]int
	// for each point in map i store the value of the shortest path that was found to reach it
	distances map[posMap]int
	//map dimension
	N int
}

// additional function, just to print the maze layout
const FULLSQUARE = 9638
const EMPTYSQUARE = 9723
const HALFLINGSTP = 9635
const TREASURE = 9714

func (inputMaze maze) printMaze() {
	N := inputMaze.N
	for y := N; y >= 0; y-- {
		if y < 100 && y >= 10 {
			fmt.Print(y, " |")
		} else if y < 10 {
			fmt.Print(y, "  |")
		} else {
			fmt.Print(y, "|")
		}
		for x := 0; x <= N; x++ {
			if inputMaze.mazeMap[posMap{x: x, y: y}] == 0 {
				fmt.Print(string(rune(EMPTYSQUARE)), " ")
			} else if inputMaze.mazeMap[posMap{x: x, y: y}] == 1 {
				fmt.Print(string(rune(FULLSQUARE)), " ")
			} else if inputMaze.mazeMap[posMap{x: x, y: y}] == 2 {
				fmt.Print(string(rune(HALFLINGSTP)), " ")
			} else {
				fmt.Print(string(rune(TREASURE)), " ")

			}
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}

// function to obtain the map of order 2^(N+1)+1
func (inputMaze *maze) updateMap() {
	N := inputMaze.N
	outputMap := maze{mazeMap: make(map[posMap]int), distances: make(map[posMap]int)}
	for y := 0; y <= N; y++ {
		for x := 0; x <= N; x++ {
			outputMap.distances[posMap{x: x, y: y}] = 0
			if x != inputMaze.N && y != N {
				//downleft
				outputMap.mazeMap[posMap{x: x, y: y}] = inputMaze.mazeMap[posMap{x: N - y, y: x}]
				//downright
				outputMap.mazeMap[posMap{x: N + x, y: y}] = inputMaze.mazeMap[posMap{x: N - y, y: N - x}]
				//uperleft
				outputMap.mazeMap[posMap{x: x, y: N + y}] = inputMaze.mazeMap[posMap{x: x, y: y}]
				//uperright
				outputMap.mazeMap[posMap{x: N + x, y: N + y}] = inputMaze.mazeMap[posMap{x: x, y: y}]
			} else {
				outputMap.mazeMap[posMap{x: N + x, y: N + y}] = 0
			}
		}
	}
	// add the wall at the intersection of new map pieces
	outputMap.mazeMap[posMap{x: N, y: N + 1}] = 1
	outputMap.mazeMap[posMap{x: 1, y: N}] = 1
	outputMap.mazeMap[posMap{x: 2*N - 1, y: N}] = 1
	inputMaze.N = 2 * N
	// copy the obtained maps to the maze
	inputMaze.mazeMap = outputMap.mazeMap
	inputMaze.distances = outputMap.distances
}

// dummy function that does 2^(N+1)
func mapDimension(N int) int {
	M := 1
	i := 1
	for i <= N+1 {
		M = 2 * M
		i++
	}
	return M
}

// a possible solution path
type path struct {
	currentPos posMap
	lenght     int
}

// function that check if i can go in a direction and update the path with the new step if possible
func (M *maze) upDatePath(xNew int, yNew int, currentLenght int, pathCovered *[]path, minPath *int) {
	//if we are still inside the maze with the index
	if xNew >= 0 && yNew >= 0 && xNew <= M.N && yNew <= M.N &&
		// and the next step was a not visited empty space
		(M.mazeMap[posMap{x: xNew, y: yNew}] == 0 ||
			// or the treasure
			M.mazeMap[posMap{x: xNew, y: yNew}] == 3 ||
			// or an empty space visited with a longer path
			(M.mazeMap[posMap{x: xNew, y: yNew}] == 2 && M.distances[posMap{x: xNew, y: yNew}] > currentLenght)) {
		// if next step is not the treasure add it to the queu
		if (M.mazeMap[posMap{x: xNew, y: yNew}] != 3) {
			*pathCovered = append(*pathCovered, path{currentPos: posMap{x: xNew, y: yNew}, lenght: currentLenght + 1})
			M.mazeMap[posMap{x: xNew, y: yNew}] = 2
			M.distances[posMap{x: xNew, y: yNew}] = currentLenght
		} else {
			//reached the tresure in next step, check if the path is the min one
			if *minPath > currentLenght {
				*minPath = currentLenght + 1
			}
		}
	}
}

// function to find all the path and the optimal one
func (M maze) move(start *posMap, goalX posMap) int {
	pathCovered := make([]path, 0) // a queu reporting the current paths checked
	pathCovered = append(pathCovered, path{currentPos: *start, lenght: 0})
	minPath := M.N * M.N // for sure the solution is less than the dimension of the maze
	// until i have element in queu i check their path to treasure
	for len(pathCovered) > 0 {
		endOfPath := len(pathCovered) - 1
		halflingPos := pathCovered[endOfPath].currentPos
		currentLenght := pathCovered[endOfPath].lenght
		if goalX.x == halflingPos.x && goalX.y == halflingPos.y {
			//reached the tresure, check if the path is the min one
			if minPath > currentLenght {
				minPath = currentLenght
			}
		} else {
			// for each direction i check if the corresponding step is free or if it was visited with a longer path

			//check left
			M.upDatePath(halflingPos.x-1, halflingPos.y, currentLenght, &pathCovered, &minPath)

			// check right
			M.upDatePath(halflingPos.x+1, halflingPos.y, currentLenght, &pathCovered, &minPath)

			//check down
			M.upDatePath(halflingPos.x, halflingPos.y-1, currentLenght, &pathCovered, &minPath)

			//check up
			M.upDatePath(halflingPos.x, halflingPos.y+1, currentLenght, &pathCovered, &minPath)

		}
		//remove the current step from queu. If no new path were found it will be deleted and moved back to an old element of the queu
		app := pathCovered[:endOfPath]
		app2 := pathCovered[endOfPath+1:]
		pathCovered = append(app, app2...)
		//M.printMaze()

	}
	//no more path available i return the shortest i found
	return minPath
}

func Solution(N int, A int, B int, C int, D int) int {

	// N is the dimension of the map 2^(N+1)+1

	//CREATE MAP ------------------------------------------------------------------
	M := mapDimension(N)

	//initialize the 4x4 map
	var N2Map = map[posMap]int{
		//first line
		posMap{x: 0, y: 0}: 0,
		posMap{x: 0, y: 1}: 0,
		posMap{x: 0, y: 2}: 0,
		posMap{x: 0, y: 3}: 0,
		posMap{x: 0, y: 4}: 0,
		//second line
		posMap{x: 1, y: 0}: 0,
		posMap{x: 1, y: 1}: 1,
		posMap{x: 1, y: 2}: 1,
		posMap{x: 1, y: 3}: 1,
		posMap{x: 1, y: 4}: 0,
		//third line
		posMap{x: 2, y: 0}: 0,
		posMap{x: 2, y: 1}: 0,
		posMap{x: 2, y: 2}: 0,
		posMap{x: 2, y: 3}: 1,
		posMap{x: 2, y: 4}: 0,
		//fourth line
		posMap{x: 3, y: 0}: 0,
		posMap{x: 3, y: 1}: 1,
		posMap{x: 3, y: 2}: 1,
		posMap{x: 3, y: 3}: 1,
		posMap{x: 3, y: 4}: 0,
		//fifth line
		posMap{x: 4, y: 0}: 0,
		posMap{x: 4, y: 1}: 0,
		posMap{x: 4, y: 2}: 0,
		posMap{x: 4, y: 3}: 0,
		posMap{x: 4, y: 4}: 0,
	}

	//initialize the distances for 4x4 map points
	var N2Distances = map[posMap]int{
		//first line
		posMap{x: 0, y: 0}: 0,
		posMap{x: 0, y: 1}: 0,
		posMap{x: 0, y: 2}: 0,
		posMap{x: 0, y: 3}: 0,
		posMap{x: 0, y: 4}: 0,
		//second line
		posMap{x: 1, y: 0}: 0,
		posMap{x: 1, y: 1}: 0,
		posMap{x: 1, y: 2}: 0,
		posMap{x: 1, y: 3}: 0,
		posMap{x: 1, y: 4}: 0,
		//third line
		posMap{x: 2, y: 0}: 0,
		posMap{x: 2, y: 1}: 0,
		posMap{x: 2, y: 2}: 0,
		posMap{x: 2, y: 3}: 1,
		posMap{x: 2, y: 4}: 0,
		//fourth line
		posMap{x: 3, y: 0}: 0,
		posMap{x: 3, y: 1}: 0,
		posMap{x: 3, y: 2}: 0,
		posMap{x: 3, y: 3}: 0,
		posMap{x: 3, y: 4}: 0,
		//fifth line
		posMap{x: 4, y: 0}: 0,
		posMap{x: 4, y: 1}: 0,
		posMap{x: 4, y: 2}: 0,
		posMap{x: 4, y: 3}: 0,
		posMap{x: 4, y: 4}: 0,
	}
	//set the maze as a 4x4 labyrinth
	labyrinth := maze{mazeMap: N2Map, distances: N2Distances, N: 4}
	//obtain the 2^(N+1)+1 dimension maze
	for labyrinth.N < M {
		labyrinth.updateMap()
	}
	halflingPos := posMap{x: A, y: B}
	treasurePos := posMap{x: C, y: D}
	//set the treasure position and halfling starting point
	labyrinth.mazeMap[halflingPos] = 2
	labyrinth.mazeMap[treasurePos] = 3
	labyrinth.printMaze()
	//find the path
	solution := labyrinth.move(&halflingPos, treasurePos)
	fmt.Println(solution)
	return solution
}
