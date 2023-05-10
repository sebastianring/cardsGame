package main

import "fmt"

// import "fmt"

func main() {
	d := newDeck()
	p := player{
		name: "Sebastian",
	}

	d.deal(3, &p)

	saveString := d.toBytes()
	fmt.Printf("The variable is of type %T and this long: %d \n", saveString, len(saveString))
	fmt.Println(saveString)
}
