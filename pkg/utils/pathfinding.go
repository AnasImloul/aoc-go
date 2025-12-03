package utils

import (
	"container/heap"
	"math"
)

// Point2D represents a 2D point with int coordinates for pathfinding algorithms.
// For generic numeric coordinates, see solver.Point[T].
type Point2D struct {
	X, Y int
}

// Point3D represents a 3D point with int coordinates.
// For generic numeric coordinates, see solver.Point3D[T].
type Point3D struct {
	X, Y, Z int
}

// isInBoundsGrid checks if a point is within grid bounds
func isInBoundsGrid(y, x int, grid [][]byte) bool {
	return y >= 0 && y < len(grid) && x >= 0 && x < len(grid[0])
}

// Add adds two points
func (p Point2D) Add(other Point2D) Point2D {
	return Point2D{X: p.X + other.X, Y: p.Y + other.Y}
}

// Sub subtracts two points
func (p Point2D) Sub(other Point2D) Point2D {
	return Point2D{X: p.X - other.X, Y: p.Y - other.Y}
}

// ManhattanDist calculates Manhattan distance to another point
func (p Point2D) ManhattanDist(other Point2D) int {
	return Abs(p.X-other.X) + Abs(p.Y-other.Y)
}

// Neighbors4 returns 4-directional neighbors
func (p Point2D) Neighbors4() []Point2D {
	return []Point2D{
		{X: p.X + 1, Y: p.Y},
		{X: p.X - 1, Y: p.Y},
		{X: p.X, Y: p.Y + 1},
		{X: p.X, Y: p.Y - 1},
	}
}

// Neighbors8 returns 8-directional neighbors (including diagonals)
func (p Point2D) Neighbors8() []Point2D {
	return []Point2D{
		{X: p.X + 1, Y: p.Y},
		{X: p.X - 1, Y: p.Y},
		{X: p.X, Y: p.Y + 1},
		{X: p.X, Y: p.Y - 1},
		{X: p.X + 1, Y: p.Y + 1},
		{X: p.X + 1, Y: p.Y - 1},
		{X: p.X - 1, Y: p.Y + 1},
		{X: p.X - 1, Y: p.Y - 1},
	}
}

// Directions
var (
	North = Point2D{X: 0, Y: -1}
	South = Point2D{X: 0, Y: 1}
	East  = Point2D{X: 1, Y: 0}
	West  = Point2D{X: -1, Y: 0}
)

// BFS performs breadth-first search on a grid
func BFS(grid [][]byte, start, end Point2D, walkable func(byte) bool) ([]Point2D, int) {
	if !isInBoundsGrid(start.Y, start.X, grid) || !isInBoundsGrid(end.Y, end.X, grid) {
		return nil, -1
	}

	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	parent := make(map[Point2D]Point2D)
	queue := []Point2D{start}
	visited[start.Y][start.X] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			// Reconstruct path
			path := []Point2D{}
			for p := end; p != start; p = parent[p] {
				path = append([]Point2D{p}, path...)
			}
			path = append([]Point2D{start}, path...)
			return path, len(path) - 1
		}

		for _, neighbor := range current.Neighbors4() {
			if !isInBoundsGrid(neighbor.Y, neighbor.X, grid) {
				continue
			}
			if visited[neighbor.Y][neighbor.X] {
				continue
			}
			if !walkable(grid[neighbor.Y][neighbor.X]) {
				continue
			}

			visited[neighbor.Y][neighbor.X] = true
			parent[neighbor] = current
			queue = append(queue, neighbor)
		}
	}

	return nil, -1
}

// AStarNode represents a node in A* search
type AStarNode struct {
	Point  Point2D
	G      int // Cost from start
	H      int // Heuristic cost to end
	F      int // Total cost (G + H)
	Parent *AStarNode
	index  int
}

// AStarQueue implements heap.Interface for A* priority queue
type AStarQueue []*AStarNode

func (pq AStarQueue) Len() int { return len(pq) }

func (pq AStarQueue) Less(i, j int) bool {
	return pq[i].F < pq[j].F
}

func (pq AStarQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *AStarQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*AStarNode)
	node.index = n
	*pq = append(*pq, node)
}

func (pq *AStarQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	node.index = -1
	*pq = old[0 : n-1]
	return node
}

// AStar performs A* pathfinding on a grid
func AStar(grid [][]byte, start, end Point2D, walkable func(byte) bool, cost func(from, to Point2D) int) ([]Point2D, int) {
	if !isInBoundsGrid(start.Y, start.X, grid) || !isInBoundsGrid(end.Y, end.X, grid) {
		return nil, -1
	}

	openSet := &AStarQueue{}
	heap.Init(openSet)

	startNode := &AStarNode{
		Point: start,
		G:     0,
		H:     start.ManhattanDist(end),
		F:     start.ManhattanDist(end),
	}
	heap.Push(openSet, startNode)

	visited := make(map[Point2D]bool)
	gScore := make(map[Point2D]int)
	gScore[start] = 0

	for openSet.Len() > 0 {
		current := heap.Pop(openSet).(*AStarNode)

		if current.Point == end {
			// Reconstruct path
			path := []Point2D{}
			for node := current; node != nil; node = node.Parent {
				path = append([]Point2D{node.Point}, path...)
			}
			return path, current.G
		}

		if visited[current.Point] {
			continue
		}
		visited[current.Point] = true

		for _, neighbor := range current.Point.Neighbors4() {
			if !isInBoundsGrid(neighbor.Y, neighbor.X, grid) {
				continue
			}
			if !walkable(grid[neighbor.Y][neighbor.X]) {
				continue
			}
			if visited[neighbor] {
				continue
			}

			moveCost := 1
			if cost != nil {
				moveCost = cost(current.Point, neighbor)
			}

			tentativeG := current.G + moveCost

			if existingG, ok := gScore[neighbor]; !ok || tentativeG < existingG {
				gScore[neighbor] = tentativeG
				h := neighbor.ManhattanDist(end)
				neighborNode := &AStarNode{
					Point:  neighbor,
					G:      tentativeG,
					H:      h,
					F:      tentativeG + h,
					Parent: current,
				}
				heap.Push(openSet, neighborNode)
			}
		}
	}

	return nil, -1
}

// Dijkstra performs Dijkstra's algorithm on a graph
func Dijkstra(graph map[string]map[string]int, start string) map[string]int {
	dist := make(map[string]int)
	visited := make(map[string]bool)

	// Initialize distances to infinity
	for node := range graph {
		dist[node] = math.MaxInt32
	}
	dist[start] = 0

	// Priority queue: (distance, node)
	type queueItem struct {
		node string
		dist int
	}

	pq := []queueItem{{start, 0}}

	for len(pq) > 0 {
		// Get node with minimum distance
		minIdx := 0
		for i := range pq {
			if pq[i].dist < pq[minIdx].dist {
				minIdx = i
			}
		}

		current := pq[minIdx]
		pq = append(pq[:minIdx], pq[minIdx+1:]...)

		if visited[current.node] {
			continue
		}
		visited[current.node] = true

		// Update distances to neighbors
		for neighbor, weight := range graph[current.node] {
			if visited[neighbor] {
				continue
			}

			newDist := dist[current.node] + weight
			if newDist < dist[neighbor] {
				dist[neighbor] = newDist
				pq = append(pq, queueItem{neighbor, newDist})
			}
		}
	}

	return dist
}

// FloodFill performs flood fill on a grid
func FloodFill(grid [][]byte, start Point2D, target, replacement byte) {
	if !isInBoundsGrid(start.Y, start.X, grid) {
		return
	}
	if grid[start.Y][start.X] != target {
		return
	}
	if target == replacement {
		return
	}

	queue := []Point2D{start}
	grid[start.Y][start.X] = replacement

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, neighbor := range current.Neighbors4() {
			if !isInBoundsGrid(neighbor.Y, neighbor.X, grid) {
				continue
			}
			if grid[neighbor.Y][neighbor.X] == target {
				grid[neighbor.Y][neighbor.X] = replacement
				queue = append(queue, neighbor)
			}
		}
	}
}

// ConnectedComponents finds all connected components in a grid
func ConnectedComponents(grid [][]byte, walkable func(byte) bool) [][]Point2D {
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	components := [][]Point2D{}

	var dfs func(Point2D, *[]Point2D)
	dfs = func(p Point2D, component *[]Point2D) {
		if !isInBoundsGrid(p.Y, p.X, grid) || visited[p.Y][p.X] || !walkable(grid[p.Y][p.X]) {
			return
		}

		visited[p.Y][p.X] = true
		*component = append(*component, p)

		for _, neighbor := range p.Neighbors4() {
			dfs(neighbor, component)
		}
	}

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if !visited[y][x] && walkable(grid[y][x]) {
				component := []Point2D{}
				dfs(Point2D{X: x, Y: y}, &component)
				if len(component) > 0 {
					components = append(components, component)
				}
			}
		}
	}

	return components
}

// TopologicalSort performs topological sort on a directed acyclic graph
func TopologicalSort(graph map[string][]string) ([]string, bool) {
	inDegree := make(map[string]int)
	nodes := make(map[string]bool)

	// Initialize in-degrees
	for node := range graph {
		nodes[node] = true
		if _, ok := inDegree[node]; !ok {
			inDegree[node] = 0
		}
	}

	for _, neighbors := range graph {
		for _, neighbor := range neighbors {
			nodes[neighbor] = true
			inDegree[neighbor]++
		}
	}

	// Find all nodes with in-degree 0
	queue := []string{}
	for node := range nodes {
		if inDegree[node] == 0 {
			queue = append(queue, node)
		}
	}

	result := []string{}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)

		for _, neighbor := range graph[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// Check if all nodes were processed (no cycle)
	return result, len(result) == len(nodes)
}


