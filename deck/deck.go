package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// ------   card and its methods
type card struct {
	suit  string
	value uint8
	desc  string
}

// -------  deck and its methods
type deck struct {
	cards []card
}

// ------ init deck, created a deck of 52 cards
func newDeck() *deck {
	cardSuits := []string{"c", "h", "d", "s"} // c = Clubs, h = Hearts, d = Diamonds, s = Spades
	var d deck

	for _, suit := range cardSuits {
		var i uint8 = 2
		for i < 15 {
			d.addCard(suit, i)
			i++
		}
	}

	return &d
}

// ------ init deck from file
func loadDeckFromFile(fileName string) *deck {
	if fileName == "" {
		fmt.Println("You need to enter a filename, exiting.")
	}

	file, err := os.ReadFile("data/" + fileName)

	if err != nil {
		fmt.Println(err)
	}

	var tempDeck deck
	startpos := 0
	for i, data := range file {
		if data == 44 {
			dataslice := file[startpos:i]
			var val int
			// fmt.Printf("Startpos: %d, current i: %d data slice: %d \n", startpos, i, dataslice)

			// Some values are stored in two bytes instead of one (when value of card > 10), so then we need to do some extra convertion
			if i-startpos > 2 {
				val, err = strconv.Atoi(string(dataslice[1:]))
				if err != nil {
					fmt.Println(err)
				}
			} else {
				val = int(dataslice[1])
			}

			// Adding cards to tempDeck, don't want to load the deck on to the original deck until all cards in the file have been read and added correctly
			tempDeck.addCard(string(dataslice[0]), uint8(val))
			startpos = i + 1
		}
	}

	return &tempDeck
}

// ------ deal a number of cards to the player, from deck at random
func (d *deck) deal(numberOfCards int, p *player) {
	length := len(d.cards) - 1
	for i := 0; i < numberOfCards; i++ {
		randomCardInt := rand.Intn(length - i)
		cardToBeDealt := d.cards[randomCardInt]

		p.cards = append(p.cards, cardToBeDealt)
		d.cards = append(d.cards[:randomCardInt], d.cards[randomCardInt+1:]...)
	}
}

// ------ adds a single card to the deck
func (d *deck) addCard(suit string, val uint8) {
	suitMap := map[string]string{
		"c": "Clubs",
		"h": "Hearts",
		"d": "Diamonds",
		"s": "Spades",
	}

	valueMap := map[uint8]string{
		2:  "Two",
		3:  "Three",
		4:  "Four",
		5:  "Five",
		6:  "Six",
		7:  "Seven",
		8:  "Eight",
		9:  "Nine",
		10: "Ten",
		11: "Knight",
		12: "Queen",
		13: "King",
		14: "Ace",
	}

	a := card{
		suit:  suit,
		value: val,
		desc:  valueMap[val] + " of " + suitMap[suit],
	}

	d.cards = append(d.cards, a)
}

// ------ returns all the cards in the deck and the total number of cards
func (d *deck) getCardsSlice(separator string, numbered bool) ([]string, int) {
	var returnString []string
	length := len(d.cards) - 1

	if numbered {
		for i, card := range d.cards {
			if i == length {
				separator = ""
			}
			returnString = append(returnString, strconv.Itoa(i)+". "+card.desc+separator)
		}
	} else {
		for i, card := range d.cards {
			if i == length {
				separator = ""
			}
			returnString = append(returnString, card.desc+separator)
		}
	}

	return returnString, len(returnString)
}

func (d *deck) getCardsShortId(separator string) []string {
	var returnString []string
	length := len(d.cards) - 1

	for i, card := range d.cards {
		if i == length {
			separator = ""
		}

		returnString = append(returnString, card.suit+strconv.Itoa(int(card.value))+separator)
	}

	return returnString
}

// ------ a function which returns all cards in the deck as a string - function to help getting it to bytes.
func (d *deck) toString() string {
	allCards := d.getCardsShortId(",")

	return strings.Join(allCards, "")
}

// ------ a function which returns all cards in the deck as bytes, will most likely be used to be saved in a file
func (d *deck) toBytes() []byte {
	allCards := d.toString()

	fmt.Println(allCards)
	return []byte(allCards)
}

// ----- save all cards to a txt-file
func (d *deck) writeCardsToFile(fileName string) {
	if fileName == "" {
		fmt.Println("You need to enter a filename, exiting.")
		return
	}

	_, err := os.ReadDir("data")
	if err != nil {
		err := os.Mkdir("data", 0755)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	file, err := os.Create("data/" + fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = file.Write(d.toBytes())
	if err != nil {
		fmt.Println(err)
		return
	}
}
