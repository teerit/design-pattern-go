package main

import (
	"fmt"
	"strings"
)

// Recursive relationship ??

type GraphicalObject struct {
	Name, Color string
	Children    []GraphicalObject
}

func (g *GraphicalObject) String() string {
	sb := strings.Builder{}
	g.print(&sb, 0)
	return sb.String()
}

func (g *GraphicalObject) print(sb *strings.Builder, depth int) {
	sb.WriteString(strings.Repeat("*", depth))
	if len(g.Color) > 0 {
		sb.WriteString(g.Color)
		sb.WriteRune(' ')
	}
	sb.WriteString(g.Name)
	sb.WriteRune('\n')

	for _, child := range g.Children {
		child.print(sb, depth+1)
	}
}

func NewCircle(color string) *GraphicalObject {
	return &GraphicalObject{"Circle", color, nil}
}

func NewSquare(color string) *GraphicalObject {
	return &GraphicalObject{"Square", color, nil}
}

func main() {
	drawing := GraphicalObject{"My Drawing", "", nil}
	drawing.Children = append(drawing.Children, *NewCircle("Red"))
	drawing.Children = append(drawing.Children, *NewCircle("Yellow"))

	group := GraphicalObject{"Group 1", "", nil}
	group.Children = append(group.Children, *NewCircle("Blue"))
	group.Children = append(group.Children, *NewCircle("Blue"))
	drawing.Children = append(drawing.Children, group)

	fmt.Println(drawing.String())
}
