package main

import "fmt"

// import "fmt"

func main() {
	d := newDeck()
	p := player{
		name: "Sebastian",
	}

	d.deal(3, &p)
	// fmt.Println(d.getCardsSlice(", ", true))

	saveString := d.toBytes()
	fmt.Printf("The variable is of type %T and this long: %d \n", saveString, len(saveString))
	// fmt.Println(saveString)
	d.writeCardsToFile("myfile.txt")

	d2 := loadDeckFromFile("myfile.txt")
	fmt.Println(d2.cards)

}
