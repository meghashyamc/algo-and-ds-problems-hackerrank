package graphs

import (
	"errors"
)

// Consider a matrix where each cell contains either a 0 or a 1 and any cell containing a 1 is called a filled cell. Two cells are said to be connected if they are adjacent to each other horizontally, vertically, or diagonally. In the diagram below, the two colored regions show cells connected to the filled cells. Black on white are not connected.
// https://s3.amazonaws.com/hr-assets/0/1528204809-ea89cbdef6-connected.png

// Cells adjacent to filled cells:

// If one or more filled cells are also connected, they form a region. Note that each cell in a region is connected to at least one other cell in the region but is not necessarily directly connected to all the other cells in the region.
// Regions:
// Given an n x m matrix,  find and print the number of cells in the largest region in the matrix.

// Complete the function maxRegion in the editor below. It must return an integer value, the size of the largest region.

// maxRegion has the following parameter(s):

// grid: a two dimensional array of integers
// more info: https://www.hackerrank.com/challenges/ctci-connected-cell-in-a-grid

// Complete the maxRegion function below.

const full = 1
const empty = 0

type position struct {
	i int
	j int
}
type stack []position

func (s *stack) push(p position) {

	if s == nil {
		s = &stack{}
	}
	*s = append((*s), p)
}

func (s *stack) pop() position {

	if s == nil || len(*s) == 0 {
		panic(errors.New("can't pop from empty stack"))
	}
	popped := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return popped
}

func maxRegion(grid [][]int32) int32 {

	if len(grid) == 0 {
		return 0
	}

	visited := make([][]bool, 0)
	for i := 0; i < len(grid); i++ {
		visited = append(visited, make([]bool, len(grid[i])))
	}

	var maxRegionVal int32
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == full && !visited[i][j] {
				if val := getValForRegion(visited, grid, position{i: i, j: j}); val > maxRegionVal {
					maxRegionVal = val
				}
			}
		}
	}
	return maxRegionVal
}

func getValForRegion(visited [][]bool, grid [][]int32, startPosition position) int32 {

	s := stack{}
	s.push(startPosition)
	visited[startPosition.i][startPosition.j] = true

	var regionVal int32 = 1
	for {
		if len(s) == 0 {
			return regionVal
		}
		popped := s.pop()
		for _, adjacentPos := range getAdjacents(grid, popped) {
			if !visited[adjacentPos.i][adjacentPos.j] {
				visited[adjacentPos.i][adjacentPos.j] = true
				s.push(adjacentPos)
				regionVal++
			}

		}
	}

}

func getAdjacents(grid [][]int32, pos position) []position {

	positions := make([]position, 0)

	if pos.j < len(grid[0])-1 && grid[pos.i][pos.j+1] == full {
		positions = append(positions, position{i: pos.i, j: pos.j + 1})
	}
	if pos.i < len(grid)-1 && grid[pos.i+1][pos.j] == full {
		positions = append(positions, position{i: pos.i + 1, j: pos.j})
	}

	if pos.j > 0 && grid[pos.i][pos.j-1] == full {
		positions = append(positions, position{i: pos.i, j: pos.j - 1})
	}

	if pos.i > 0 && grid[pos.i-1][pos.j] == full {
		positions = append(positions, position{i: pos.i - 1, j: pos.j})
	}

	if pos.i > 0 && pos.j < len(grid[0])-1 && grid[pos.i-1][pos.j+1] == full {
		positions = append(positions, position{i: pos.i - 1, j: pos.j + 1})
	}

	if pos.i > 0 && pos.j > 0 && grid[pos.i-1][pos.j-1] == full {
		positions = append(positions, position{i: pos.i - 1, j: pos.j - 1})
	}

	if pos.i < len(grid)-1 && pos.j < len(grid[0])-1 && grid[pos.i+1][pos.j+1] == full {
		positions = append(positions, position{i: pos.i + 1, j: pos.j + 1})
	}

	if pos.i < len(grid)-1 && pos.j > 0 && grid[pos.i+1][pos.j-1] == full {
		positions = append(positions, position{i: pos.i + 1, j: pos.j - 1})
	}
	return positions
}
