package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 56 {
		t.Errorf("Expected length  of 56, but got %v", len(d))
	}

	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected the first card to be 'Ace of Hearts', but got %v", d[0])
	}

	if d[len(d)-1] != "King of Diamonds" {
		t.Errorf("Expected the last card to be 'King of Diamond', but got %v", d[0])
	}
}

func TestWriteFileSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("tmp/_decktesting")

	deck := newDeck()
	deck.saveToFile("tmp/_decktesting")

	loadedDeck := newDeckFromFile("tmp/_decktesting")

	if len(loadedDeck) != 56 {
		t.Errorf("Expected to load 56 cards from deck file, but got %v", len(loadedDeck))
	}

	os.Remove("tmp/_decktesting")
}
