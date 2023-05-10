package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

// ------   card and its methods
type card struct {
	desc  string
	value uint8
}

// -------  deck and its methods
type deck struct {
	cards []card
}

// ------ init deck, created a deck of 52 cards
func newDeck() deck {
	cardSuits := []string{"Hearts", "Diamonds", "Spades", "Clubs"}
	cardValuesStrings := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Knight", "Queen", "King", "Ace"}
	var d deck

	for _, suit := range cardSuits {
		for val, valueString := range cardValuesStrings {
			d.addCard((valueString + " of " + suit), uint8(val+2))
		}
	}

	return d
}

// ------ deal a number of cards to the player, from a deck at random
func (d *deck) deal(numberOfCards int, p *player) {
	length := len(d.cards) - 1
	for i := 0; i < numberOfCards; i++ {
		randomCardInt := rand.Intn(length - i)
		fmt.Println(randomCardInt, d.cards[randomCardInt])
		cardToBeDealt := d.cards[randomCardInt]
		p.cards = append(p.cards, cardToBeDealt)
		d.cards = append(d.cards[:randomCardInt], d.cards[randomCardInt+1:]...)
	}
}

// ------ adds a single card to the deck
func (d *deck) addCard(desc string, val uint8) {
	a := card{
		desc:  desc,
		value: val,
	}
	d.cards = append(d.cards, a)
}

// ------ returns all the cards in the deck and the total number of cards
func (d *deck) getCardsSlice(separator string, numbered bool) ([]string, int) {
	var returnString []string
	length := len(d.cards)

	if numbered {
		for i, card := range d.cards {
			if i == length {
				separator = ""
			}
			returnString = append(returnString, strconv.Itoa(i)+separator+card.desc)
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

// ------ a function which returns all cards in the deck as a string - function to help getting it to bytes.
func (d *deck) toString() string {
	allCards, _ := d.getCardsSlice("", false)

	return strings.Join(allCards, ", ")
}

// ------ a function which returns all cards in the deck as bytes, will most likely be used to be saved in a file
func (d *deck) toBytes() []byte {
	allCards := d.toString()

	return []byte(allCards)
}
