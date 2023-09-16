package main

import "fmt"

type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getPower() int {
	return g.power
}

type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "Ak47 gun",
			power: 5,
		},
	}
}

func getGun(gunType string) (IGun, error) {
	if gunType == "Ak47" {
		return newAk47(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}

func main() {
	ak47, err := getGun("Ak47")
	if err != nil {
		fmt.Println(err)
	}

	printDetails(ak47)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s\n", g.getName())
	fmt.Printf("Power: %d\n", g.getPower())
}
