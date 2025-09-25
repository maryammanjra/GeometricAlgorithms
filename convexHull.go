package main

import (
	"fmt"
	"math"
	"sort"
)

// Assuming no collinear points
type Point struct {
	x float64
	y float64
}

type Stack struct {
	stack []Point
}

func (s *Stack) push(p Point) {
	s.stack = append(s.stack, p)
}

func (s *Stack) isEmpty() bool {
	return len(s.stack) == 0
}

func (s *Stack) pop() Point {
	if !s.isEmpty() {
		popped := s.stack[len(s.stack)-1]
		s.stack = s.stack[:len(s.stack)-1]
		return popped
	}
	return Point{-1, -1}
}

func orientedArea(p1 Point, p2 Point, p3 Point) float64 {
	vectorOne := Point{p2.x - p1.x, p2.y - p1.y}
	vectorTwo := Point{p3.x - p2.x, p3.y - p2.y}
	fmt.Println(vectorOne)
	return (vectorOne.x * vectorTwo.y) - (vectorOne.y * vectorTwo.x)
}

// findPolarAngle assumes p2 lies either in quadrant one or quadrant two relative to p1 as origin,
// based on Graham's scan calculating polar angles in reference to the lowest Y-coordinate.
func findPolarAngle(p1 Point, p2 Point) float64 {
	xVector := p2.x - p1.x
	yVector := p1.y - p1.y

	if xVector < 0 {
		return math.Pi + math.Atan(yVector/xVector)
	}

	return math.Atan(yVector / xVector)
}

func findSmallestY(points []Point) Point {
	minPoint := points[0]

	for _, val := range points {
		if val.y < minPoint.y {
			minPoint = val
		}
	}

	return minPoint
}

func sortByPolarAngle(points []Point, smallestY Point) {
	sort.Slice(points, func(i, j int) bool {
		return findPolarAngle(smallestY, points[i]) > findPolarAngle(smallestY, points[j])
	})
}

func grahamsScan(points []Point) []Point {
	startingPoint := findSmallestY(points)
	sortByPolarAngle(points, startingPoint)
}

func main() {

}
