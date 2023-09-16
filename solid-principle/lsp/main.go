package main

import "fmt"

// Liskov Substitution Principle
// In golang, there is no class-based inheritance.
// Instead, Golang proviedes a more powerful approach towrads
// polymorphism via Interface and Struct embeddeing.

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	s.height = height
	s.width = height
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)

	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()

	fmt.Print("Expected an area of ", expectedArea, ", but we got ", actualArea, "\n")
}

type Square2 struct {
	size int // width, height
}

func (s *Square2) Rectangle() Rectangle {
	return Rectangle{s.size, s.size}
}

func main() {
	rc := &Rectangle{2, 3}
	UseIt(rc)

	sq := NewSquare(5)
	UseIt(sq)
}
