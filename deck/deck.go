package main

import (
	"crypto/rand"
	"fmt"
)

type player struct {
	name  string
	cards []card
}

type deck struct {
	cards []card
}

type card struct {
	desc  string
	value uint8
}

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

func (d *deck) deal(numberOfCards int, p *player) {
	length := len(d.cards) - 1
	for i := 0; i < numberOfCards; i++ {
		lol, randomCardInt := rand.Intn(length - i)
		cardToBeDealt := d.cards[randomCardInt]
		p.cards = append(p.cards, cardToBeDealt)
		d.cards = append(d.cards[:randomCardInt], d.cards[randomCardInt+1:])
	}
}

func (d *deck) addCard(desc string, val uint8) {
	a := card{
		desc:  desc,
		value: val,
	}
	d.cards = append(d.cards, a)
	fmt.Printf("Card added: %s, %d \n", a.desc, a.value)
}

func (d *deck) getCards() []string {
	var returnString []string

	for _, card := range d.cards {
		returnString = append(returnString, card.desc)
	}

	return returnString
}
