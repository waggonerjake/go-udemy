package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

//Kind of makes the deck function as a class
type deck []string

func newDeck() deck {
	cards := deck{}
	cardsSuits := []string{"Hearts", "Spades", "Diamonds", "Clubs"}
	cardsValues := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	for _, suit := range cardsSuits {
		for _, value := range cardsValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

//Receiver; whoever calls this function is what gets passed into the d variable (kind of like "this"),
//but needs to be of type deck. Essentially makes it a member function without OOP
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return deck(strings.Split(string(data), ","))
}

//Will change the deck by reference
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	for i := range d {
		newPosition := rng.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
