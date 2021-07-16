package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected Deck length of 52, but got %d", len(d))
	}
	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected first card to be Ace of Hearts, but got %s", d[0])
	}
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs, but got %s", d[len(d)-1])
	}
}

func TestSaveToDeckAndReadDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 52 {
		t.Errorf("Expected Deck length of 52, but got %d", len(d))
	}

	os.Remove("_decktesting")
}
