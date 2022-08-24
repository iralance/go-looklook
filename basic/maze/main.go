package main

import (
	"fmt"
	"os"
)

type point struct {
	i, j int
}

var dirs = [4]point{
	{0, -1}, {-1, 0}, {1, 0}, {0, 1},
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}
		for _, dir := range dirs {
			next := cur.add(dir)
			//判断next是否在maze里
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}
			//判断next是否已经走过
			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}
			//判断next是否是起点
			if next == start {
				continue
			}
			curStep := steps[cur.i][cur.j]
			steps[next.i][next.j] = curStep + 1
			Q = append(Q, next)
		}
	}
	return steps
}

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var col, row int
	fmt.Fscanf(file, "%d %d", &row, &col)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

func findPath(steps [][]int) []point {
	start := point{0, 0}
	end := point{len(steps) - 1, len(steps[0]) - 1}
	//fmt.Println(start, end)
	value := steps[len(steps)-1][len(steps[0])-1]

	var list = []point{end}
	for end != start {
		for _, dir := range dirs {
			//4个方向 先判断 边界范围 其次 判断值
			prev := end.add(dir)
			if _, ok := prev.at(steps); ok {
				if steps[prev.i][prev.j] == (value - 1) {
					list = append(list, prev)
					end = prev
					value--
					break
				}
			}
		}
	}
	return list
}

func main() {
	maze := readMaze("maze.in")
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})

	fmt.Println(steps)
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}

	fmt.Println("最短需要走的步数：", steps[len(maze)-1][len(maze[0])-1])
	fmt.Print("最短路径坐标为：")
	paths := findPath(steps)
	for i := len(paths) - 1; i >= 0; i-- {
		fmt.Printf("(%d, %d) ", paths[i].i, paths[i].j)
	}
}
