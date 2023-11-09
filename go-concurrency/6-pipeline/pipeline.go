package main

import "fmt"

type Cake struct {
	id       int
	baked    bool
	topping  string
	filling  string
	packaged bool
}

func bake(in <-chan Cake) chan Cake {
	out := make(chan Cake)
	go func() {
		for cake := range in {
			cake.baked = true
			fmt.Println("Baking", cake)
			out <- cake
		}
		close(out)
	}()
	return out
}

func filling(in <-chan Cake) chan Cake {
	out := make(chan Cake)
	go func() {
		for cake := range in {
			cake.filling = "cream"
			fmt.Println("Filling:", cake)
			out <- cake
		}
		close(out)
	}()
	return out
}

func topping(in <-chan Cake) chan Cake {
	out := make(chan Cake)
	go func() {
		for cake := range in {
			cake.topping = "strawberry"
			fmt.Println("Topping", cake)
			out <- cake
		}
		close(out)
	}()
	return out
}

func packaged(in <-chan Cake) chan Cake {
	out := make(chan Cake)
	go func() {
		for cake := range in {
			cake.packaged = true
			fmt.Println("Packaging:", cake)
			out <- cake
		}
		close(out)
	}()
	return out
}

func newCakes() chan Cake {
	out := make(chan Cake)
	go func() {
		for i := 0; i < 5; i++ {
			cake := Cake{
				id:      i,
				filling: "none",
				topping: "none",
			}
			fmt.Println("New Cake:", cake)
			out <- cake
		}
		close(out)
	}()
	return out
}

func main() {
	cakes := packaged(topping(filling(bake(newCakes()))))

	for cake := range cakes {
		fmt.Println("Done:", cake)
	}
}
