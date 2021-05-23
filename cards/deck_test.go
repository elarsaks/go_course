package main

import (
	"os"
	"testing"
)

func TestNewDec(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d) - 1] != "King of Clubs" {
		t.Errorf("Expected first card of King of Clubs, but got %v", d[len(d) - 1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
	os.Remove("_descktesting")

	deck := newDeck()
	deck.saveToFile("_descktesting")

	loadedDeck := newDeckFromFile("_descktesting")

	if len(loadedDeck) != 52{
		t.Errorf("Expected 52 cards in deck, git %v", len(loadedDeck))
	}

	os.Remove("_descktesting")
}