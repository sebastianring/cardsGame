package main

import "fmt"

func main() {

	d := newDeck()

	p := player{
		name: "Sebastian",
	}

	d.deal(3, p)

	allCards := d.getCards()
	for _, val := range allCards {
		fmt.Println(val)
	}
}
