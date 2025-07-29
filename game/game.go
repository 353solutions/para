package main

import "fmt"

func main() {
	i, err := NewItem(10, 20)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println("i:", i)
	i.Move(200, 300)
	fmt.Println("i (move):", i)

	i2 := Item{30, 40}
	i2.Move(3, 4) // Go will pass i2 as pointer to Move
	// C++: i2->Move(3, 4);
	fmt.Println("i2 (move):", i2)
}

// value -> pointer is OK : inc(&n)
// pointer -> value is not OK: inc(*n)

// i is called "the receiver"
// i is a pointer receiver
func (i *Item) Move(x, y int) {
	i.X = x
	i.Y = y
}

// New/Factory functions
/*
func NewItem(x, y int) Item
func NewItem(x, y int) (Item, error)
func NewItem(x, y int) *Item
*/
func NewItem(x, y int) (*Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("%d/%d out of range for %d/%d", x, y, maxX, maxY)
		// return Item{}, fmt.Errorf("%d/%d out of range for %d/%d", x, y, maxX, maxY)
	}

	i := Item{
		X: x,
		Y: y,
	}
	// Go does escape analysis and will allocate i on the heap
	// go build -gcflags=-m
	return &i, nil
}

const (
	maxX = 400
	maxY = 600
)

type Item struct {
	X int
	Y int
}
