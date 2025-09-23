package main

type Point struct {
	x float32
	y float32
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

func orientedArea(p1 Point, p2 Point, p3 Point) float32 {
	vectorOne := Point{p2.x - p1.x, p2.y - p1.y}
	vectorTwo := Point{p3.x - p2.x, p3.y - p2.y}
	return (vectorOne.x * vectorTwo.y) - (vectorOne.y * vectorTwo.x)
}

func main() {

}
