package main

import (
	"fmt"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	fmt.Println("Testing to see if a new deck has 52 cards")
	if len(d.cards) != 52 {
		t.Errorf("Expected length: 52 but got: %d \n", len(d.cards))
	}
}

func TestNewDockFromFile(t *testing.T) {
	fmt.Println("Testing to see if a good file can be loaded")
	d := newDeckFromFile("testfile.txt")

	if len(d.cards) == 0 {
		t.Errorf("Expected some cards, got none. Might be an issue with the file. Length: %v ", len(d.cards))
	}

	d2 := newDeckFromFile("")
	if len(d2.cards) != 0 {
		t.Errorf("Did not expect any cards, as the function will not run if no filename is given. Length: %v ", len(d2.cards))
	}

	d3 := newDeckFromFile("testfilecorrupt.txt")
	if len(d3.cards) != 0 {
		fmt.Println(d3)
		t.Errorf("Did not expect any cards as the file is not correctly formatted. Cards: %v ", d3.cards)
	}
}

func TestHappyflowDeckPlayerDeal(t *testing.T) {
	fmt.Println("Testing to create a new deck, a player and dealing 3 cards from the deck to the player")

	d := newDeck()
	var p player

	d.deal(3, &p)
	if len(p.cards) != 3 {
		t.Errorf("Player does not have 3 cards")
	}
}

func TestEdgecaseDeckPlayerDeal(t *testing.T) {
	fmt.Println("Testing to deal 0 cards")
	d := newDeck()
	var p player

	d.deal(0, &p)
	if len(p.cards) > 0 {
		t.Errorf("Player does have cards, expected 0")
	}
}
