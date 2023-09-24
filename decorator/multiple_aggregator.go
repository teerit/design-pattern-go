package main

import "fmt"

type Aged interface {
	Age() int
	SetAge(age int)
}

type Bird struct {
	age int
}

// implement Aged interface
func (b *Bird) Age() int {
	return b.age
}

func (b *Bird) SetAge(age int) {
	b.age = age
}

func (b *Bird) Fly() {
	if b.age >= 10 {
		fmt.Println("Flying!")
	}
}

type Lizard struct {
	age int
}

// implement Aged interface
func (l *Lizard) Age() int {
	return l.age
}

func (l *Lizard) SetAge(age int) {
	l.age = age
}

func (l *Lizard) Crawl() {
	if l.age < 10 {
		fmt.Println("Crawling!")
	}
}

type Dragon struct {
	bird   Bird
	lizard Lizard
}

// implement Aged interface
func (d *Dragon) Age() int {
	return d.bird.age
}

func (d *Dragon) SetAge(age int) {
	d.bird.age = age
	d.lizard.age = age
}

func (d *Dragon) Fly() {
	d.bird.Fly()
}

func (d *Dragon) Crawl() {
	d.lizard.Crawl()
}

func NewDragon() *Dragon {
	return &Dragon{Bird{}, Lizard{}}
}

func main() {
	d := Dragon{}
	d.SetAge(5)
	d.Crawl()
	d.Fly()

	// Shape
	circle := Circle{2}
	circle.Resize(2)
	fmt.Println(circle.Render())

	redCircle := ColoredShape{&circle, "Red"}
	fmt.Println(redCircle.Render())

	rhsCircle := TransparentShape{&redCircle, 0.5}
	fmt.Println(rhsCircle.Render())
}
