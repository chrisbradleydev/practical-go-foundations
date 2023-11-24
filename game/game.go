package main

import (
	"fmt"
	"slices"
	"sort"
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
	Keys []Key
}

// type T struct {
// 	X int
// }

type ByDistance []Player

func (a ByDistance) Len() int {
	return len(a)
}
func (a ByDistance) Swap(i, j int) {
	 a[i], a[j] = a[j], a[i]
}
func (a ByDistance) Less(i, j int) bool {
	return a[i].X < a[j].X
}

type mover interface {
	Move(x, y int)
	// Move(int, int)
}

func (p *Player) FoundKey(k Key) error {
	if k < Jade || k >= invalidKey {
		return fmt.Errorf("invalid key: %#v", k)
	}
	// if !containsKey(p.Keys, k) {
	// 	p.Keys = append(p.Keys, k)
	// }
	if !slices.Contains(p.Keys, k) {
		p.Keys = append(p.Keys, k)
	}
	return nil
}

// func containsKey(keys []Key, k Key) bool {
// 	for _, key := range keys {
// 		if key == k {
// 			return true
// 		}
// 	}
// 	return false
// }

const (
	maxX = 1000
	maxY = 600
)

// go's version of "enum"
const (
	Jade Key = iota + 1
	Copper
	Crystal
	invalidKey // internal (not exported)
)

type Key byte

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

	k := Jade
	fmt.Println("k:", k)
	fmt.Println("key:", Key(17))

	// time.Time implement json.Marshaler interface
	// json.NewEncoder(os.Stdout).Encode(time.Now())

	p1.FoundKey(Jade)
	fmt.Println(p1.Keys)
	p1.FoundKey(Jade)
	fmt.Println(p1.Keys)

	players := []Player{
		{Name: "Alan", Item: Item{40, 0}},
		{Name: "Bobby", Item: Item{30, 0}},
		{Name: "Charles", Item: Item{100, 0}},
		{Name: "Dylan", Item: Item{0, 0}},
	}

	fmt.Println("before sort:", players)
	sort.Sort(ByDistance(players))
	fmt.Println("after sort:", players)

	sort.Slice(players, func(i, j int) bool {
		return players[i].Item.X > players[j].Item.X
	})
	fmt.Println("using slice:", players)
}

// implement fmt.Stringer interface
func (k Key) String() string {
	switch k {
	case Jade:
		return "jade"
	case Copper:
		return "copper"
	case Crystal:
		return "crystal"
	}
	return fmt.Sprintf("<Key %d>", k)
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