package main

import (
	"fmt"
)

type Item struct {
	X int
	Y int
}

type Player struct {
	Name string
	// X int
	Item // embed item
	// T
}

// type T struct {
// 	X int
// }

type mover interface {
	Move(x, y int)
	// Move(int, int)
}

const (
	maxX = 1000
	maxY = 600
)

func main() {
	var i1 Item
	fmt.Println(i1)
	fmt.Printf("i1: %#v\n", i1)

	i2 := Item{1, 2}
	fmt.Printf("i2: %#v\n", i2)

	i3 := Item{
		Y: 10,
		// X: 20,
	}
	fmt.Printf("i3: %#v\n", i3)
	fmt.Println(NewItem(10, 20))
	fmt.Println(NewItem(10, -20))

	i3.Move(100, 200)
	fmt.Printf("i3 (move): %#v\n", i3)

	p1 := Player{
		Name: "Parzival",
		Item: Item{500, 300},
	}
	fmt.Printf("p1: %#v\n", p1)
	fmt.Printf("p1.X: %#v\n", p1.X)
	fmt.Printf("p1.Item.X: %#v\n", p1.Item.X)
	p1.Move(400, 600)
	fmt.Printf("p1 (move): %#v\n", p1)

	ms := []mover{
		&i1,
		&p1,
		&i2,
	}
	moveAll(ms, 0, 0)
	for _, m := range ms {
		fmt.Println(m)
	}
}

/* go >= 1.18
func NewNumber[T int | float64](kind string) T {
	if kind == "int" {
		return 0
	}
	return 0.0
}
*/

// rule of thumb: accept interfaces, return types

func moveAll(ms []mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)
	}
}

// i is called "the receiver"
// if you want to mutate, use pointer receiver
func (i *Item) Move(x, y int) {
	i.X = x
	i.Y = y
}

// func NewItem(x, y int) Item {}
// func NewItem(x, y int) *Item {}
// func NewItem(x, y int) (Item, error) {}
func NewItem(x, y int) (*Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("%d/%d out of bounds %d/%d", x, y, maxX, maxY)
	}

	i := Item{
		X: x,
		Y: y,
	}
	// the go compiler does "escape analysis" and will allocate i on the heap
	return &i, nil
}