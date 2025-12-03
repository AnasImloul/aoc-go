package utils

import (
	"testing"
)

func TestPoint2DAdd(t *testing.T) {
	p1 := Point2D{X: 1, Y: 2}
	p2 := Point2D{X: 3, Y: 4}
	result := p1.Add(p2)
	expected := Point2D{X: 4, Y: 6}

	if result != expected {
		t.Errorf("Point2D.Add = %v, want %v", result, expected)
	}
}

func TestPoint2DSub(t *testing.T) {
	p1 := Point2D{X: 5, Y: 7}
	p2 := Point2D{X: 2, Y: 3}
	result := p1.Sub(p2)
	expected := Point2D{X: 3, Y: 4}

	if result != expected {
		t.Errorf("Point2D.Sub = %v, want %v", result, expected)
	}
}

func TestPoint2DManhattanDist(t *testing.T) {
	tests := []struct {
		p1, p2   Point2D
		expected int
	}{
		{Point2D{0, 0}, Point2D{0, 0}, 0},
		{Point2D{0, 0}, Point2D{1, 1}, 2},
		{Point2D{0, 0}, Point2D{3, 4}, 7},
		{Point2D{1, 1}, Point2D{4, 5}, 7},
		{Point2D{5, 5}, Point2D{2, 1}, 7},
	}

	for _, tt := range tests {
		result := tt.p1.ManhattanDist(tt.p2)
		if result != tt.expected {
			t.Errorf("ManhattanDist(%v, %v) = %d, want %d", tt.p1, tt.p2, result, tt.expected)
		}
	}
}

func TestPoint2DNeighbors4(t *testing.T) {
	p := Point2D{X: 5, Y: 5}
	neighbors := p.Neighbors4()

	if len(neighbors) != 4 {
		t.Fatalf("Neighbors4 returned %d neighbors, want 4", len(neighbors))
	}

	// Check that all expected neighbors are present
	expected := map[Point2D]bool{
		{X: 6, Y: 5}: false,
		{X: 4, Y: 5}: false,
		{X: 5, Y: 6}: false,
		{X: 5, Y: 4}: false,
	}

	for _, n := range neighbors {
		if _, ok := expected[n]; !ok {
			t.Errorf("Unexpected neighbor: %v", n)
		}
		expected[n] = true
	}

	for n, found := range expected {
		if !found {
			t.Errorf("Missing expected neighbor: %v", n)
		}
	}
}

func TestPoint2DNeighbors8(t *testing.T) {
	p := Point2D{X: 5, Y: 5}
	neighbors := p.Neighbors8()

	if len(neighbors) != 8 {
		t.Fatalf("Neighbors8 returned %d neighbors, want 8", len(neighbors))
	}
}

func TestBFS(t *testing.T) {
	// Simple grid:
	// . . .
	// . # .
	// . . .
	grid := [][]byte{
		{'.', '.', '.'},
		{'.', '#', '.'},
		{'.', '.', '.'},
	}

	walkable := func(b byte) bool { return b == '.' }

	path, dist := BFS(grid, Point2D{0, 0}, Point2D{2, 2}, walkable)

	if dist == -1 {
		t.Fatal("BFS returned no path")
	}

	if dist != 4 {
		t.Errorf("BFS distance = %d, want 4", dist)
	}

	if len(path) != 5 {
		t.Errorf("BFS path length = %d, want 5", len(path))
	}

	// Verify path starts and ends correctly
	if path[0] != (Point2D{0, 0}) {
		t.Errorf("BFS path start = %v, want {0, 0}", path[0])
	}
	if path[len(path)-1] != (Point2D{2, 2}) {
		t.Errorf("BFS path end = %v, want {2, 2}", path[len(path)-1])
	}
}

func TestBFSNoPath(t *testing.T) {
	// Grid with no path:
	// . # .
	// # # #
	// . # .
	grid := [][]byte{
		{'.', '#', '.'},
		{'#', '#', '#'},
		{'.', '#', '.'},
	}

	walkable := func(b byte) bool { return b == '.' }

	path, dist := BFS(grid, Point2D{0, 0}, Point2D{2, 2}, walkable)

	if dist != -1 {
		t.Errorf("BFS should return -1 for unreachable target, got %d", dist)
	}
	if path != nil {
		t.Errorf("BFS should return nil path for unreachable target")
	}
}

func TestAStar(t *testing.T) {
	grid := [][]byte{
		{'.', '.', '.'},
		{'.', '#', '.'},
		{'.', '.', '.'},
	}

	walkable := func(b byte) bool { return b == '.' }

	path, dist := AStar(grid, Point2D{0, 0}, Point2D{2, 2}, walkable, nil)

	if dist == -1 {
		t.Fatal("AStar returned no path")
	}

	if dist != 4 {
		t.Errorf("AStar distance = %d, want 4", dist)
	}

	if path[0] != (Point2D{0, 0}) {
		t.Errorf("AStar path start = %v, want {0, 0}", path[0])
	}
}

func TestDijkstra(t *testing.T) {
	graph := map[string]map[string]int{
		"A": {"B": 1, "C": 4},
		"B": {"C": 2, "D": 5},
		"C": {"D": 1},
		"D": {},
	}

	distances := Dijkstra(graph, "A")

	expected := map[string]int{
		"A": 0,
		"B": 1,
		"C": 3, // A->B->C
		"D": 4, // A->B->C->D
	}

	for node, expectedDist := range expected {
		if distances[node] != expectedDist {
			t.Errorf("Dijkstra distance to %s = %d, want %d", node, distances[node], expectedDist)
		}
	}
}

func TestFloodFill(t *testing.T) {
	grid := [][]byte{
		{'.', '.', '#'},
		{'.', '.', '#'},
		{'#', '#', '#'},
	}

	FloodFill(grid, Point2D{0, 0}, '.', 'X')

	// Check that the filled region is correct
	if grid[0][0] != 'X' || grid[0][1] != 'X' || grid[1][0] != 'X' || grid[1][1] != 'X' {
		t.Error("FloodFill did not fill connected region correctly")
	}

	// Check that walls are not filled
	if grid[0][2] == 'X' || grid[2][0] == 'X' {
		t.Error("FloodFill incorrectly filled wall cells")
	}
}

func TestConnectedComponents(t *testing.T) {
	grid := [][]byte{
		{'.', '#', '.'},
		{'.', '#', '.'},
		{'#', '#', '#'},
	}

	walkable := func(b byte) bool { return b == '.' }
	components := ConnectedComponents(grid, walkable)

	if len(components) != 2 {
		t.Errorf("ConnectedComponents found %d components, want 2", len(components))
	}
}

func TestTopologicalSort(t *testing.T) {
	graph := map[string][]string{
		"A": {"B", "C"},
		"B": {"D"},
		"C": {"D"},
		"D": {},
	}

	order, ok := TopologicalSort(graph)

	if !ok {
		t.Fatal("TopologicalSort returned cycle detected for acyclic graph")
	}

	if len(order) != 4 {
		t.Errorf("TopologicalSort returned %d nodes, want 4", len(order))
	}

	// Verify ordering constraints
	indexOf := make(map[string]int)
	for i, node := range order {
		indexOf[node] = i
	}

	// A must come before B and C
	if indexOf["A"] > indexOf["B"] || indexOf["A"] > indexOf["C"] {
		t.Error("TopologicalSort: A should come before B and C")
	}

	// B and C must come before D
	if indexOf["B"] > indexOf["D"] || indexOf["C"] > indexOf["D"] {
		t.Error("TopologicalSort: B and C should come before D")
	}
}
