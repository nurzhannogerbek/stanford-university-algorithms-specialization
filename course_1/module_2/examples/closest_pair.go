package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

type Point struct {
	X, Y float64
}

// ClosestPairSolver initializes the solver with the list of points.
type ClosestPairSolver struct {
	points []Point
}

func NewClosestPairSolver(points []Point) *ClosestPairSolver {
	return &ClosestPairSolver{points: points}
}

// EuclideanDistance calculates the Euclidean distance between two points.
func EuclideanDistance(p1, p2 Point) float64 {
	return math.Sqrt((p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y))
}

// BruteForce finds the closest pair using brute-force search.
func (solver *ClosestPairSolver) BruteForce(points []Point) (float64, [2]Point) {
	minDist := math.Inf(1)
	var closestPair [2]Point
	n := len(points)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist := EuclideanDistance(points[i], points[j])
			if dist < minDist {
				minDist = dist
				closestPair = [2]Point{points[i], points[j]}
			}
		}
	}
	return minDist, closestPair
}

// ClosestSplitPair checks for the closest pair across the divide.
func (solver *ClosestPairSolver) ClosestSplitPair(px, py []Point, delta float64, bestPair [2]Point) (float64, [2]Point) {
	midX := px[len(px)/2].X
	var sy []Point

	// Collect points within the delta-width vertical strip.
	for _, p := range py {
		if math.Abs(p.X-midX) < delta {
			sy = append(sy, p)
		}
	}

	bestDist := delta
	lenSy := len(sy)

	// Compare up to 7 closest neighbors.
	for i := 0; i < lenSy; i++ {
		for j := 1; j < 8 && i+j < lenSy; j++ {
			p, q := sy[i], sy[i+j]
			dist := EuclideanDistance(p, q)
			if dist < bestDist {
				bestDist = dist
				bestPair = [2]Point{p, q}
			}
		}
	}
	return bestDist, bestPair
}

// ClosestPairRecursive is a recursive divide-and-conquer method.
func (solver *ClosestPairSolver) ClosestPairRecursive(px, py []Point) (float64, [2]Point) {
	n := len(px)
	if n <= 3 {
		return solver.BruteForce(px)
	}

	mid := n / 2
	Qx := px[:mid]
	Rx := px[mid:]

	// Split py into Qy and Ry.
	var Qy, Ry []Point
	midpointX := px[mid].X
	for _, p := range py {
		if p.X <= midpointX {
			Qy = append(Qy, p)
		} else {
			Ry = append(Ry, p)
		}
	}

	// Recursive calls for left and right halves.
	distLeft, pairLeft := solver.ClosestPairRecursive(Qx, Qy)
	distRight, pairRight := solver.ClosestPairRecursive(Rx, Ry)

	// Choose the minimum distance from two halves.
	var delta float64
	var bestPair [2]Point
	if distLeft < distRight {
		delta = distLeft
		bestPair = pairLeft
	} else {
		delta = distRight
		bestPair = pairRight
	}

	// Check for closest split pair.
	distSplit, pairSplit := solver.ClosestSplitPair(px, py, delta, bestPair)

	// Return the smallest of the distances found.
	if delta < distSplit {
		return delta, bestPair
	} else {
		return distSplit, pairSplit
	}
}

// FindClosestPair initializes and runs the closest-pair search.
func (solver *ClosestPairSolver) FindClosestPair() (float64, [2]Point, error) {
	if len(solver.points) < 2 {
		return 0, [2]Point{}, errors.New("at least two points are required to find the closest pair")
	}

	// Sort points by X and Y coordinates.
	px := make([]Point, len(solver.points))
	py := make([]Point, len(solver.points))
	copy(px, solver.points)
	copy(py, solver.points)

	sort.Slice(px, func(i, j int) bool {
		return px[i].X < px[j].X
	})
	sort.Slice(py, func(i, j int) bool {
		return py[i].Y < py[j].Y
	})

	// Run recursive algorithm.
	dist, pair := solver.ClosestPairRecursive(px, py)
	return dist, pair, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	points := []Point{{2.1, 3.5}, {1.1, 2.9}, {3.6, 4.7}, {4.2, 2.1}, {0.9, 1.5}}
	closestPairSolver := NewClosestPairSolver(points)
	distance, pair, err := closestPairSolver.FindClosestPair()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("The closest pair is %v with a distance of %.2f\n", pair, distance)

	// Edge case: fewer than two points.
	singlePoint := []Point{{2.1, 3.5}}
	closestPairSolver = NewClosestPairSolver(singlePoint)
	distance, pair, err = closestPairSolver.FindClosestPair()
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Performance testing with a large number of points.
	var largePoints []Point
	for i := 0; i < 10000; i++ {
		largePoints = append(largePoints, Point{
			X: (rand.Float64() * 2000) - 1000,
			Y: (rand.Float64() * 2000) - 1000,
		})
	}
	closestPairSolver = NewClosestPairSolver(largePoints)
	distance, pair, err = closestPairSolver.FindClosestPair()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Closest pair in large dataset: %v with a distance of %.2f\n", pair, distance)
}
