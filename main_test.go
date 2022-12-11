package main

import (
	"reflect"
	"testing"
)

func TestMajorArcana(t *testing.T) {
	// Generate the major arcana using the majorArcana function
	arcana := majorArcana()

	// Iterate over the arcana
	for _, card := range arcana {
		// Check if the card's number and nameSuit match the expected values
		if card.nameSuit != majorCards[card.number] {
			// If they do not match, fail the test
			t.Errorf("Unexpected card. Expected %s but got %s", majorCards[card.number], card.nameSuit)
		}
	}
}

func TestMinorArcana(t *testing.T) {
	// call minorArcana and save the result in a variable
	result := minorArcana()

	// assert that the length of the result is correct
	if len(result) != len(minorSuits)*len(minorCards) {
		t.Errorf("Expected length of %d but got %d", len(minorSuits)*len(minorCards), len(result))
	}

	// assert that the elements of the result have the expected values
	for _, suit := range minorSuits {
		for _, number := range minorCards {
			card := tarotDeck{
				number:   number,
				nameSuit: "of " + suit,
			}
			if !contains(result, card) {
				t.Errorf("Expected to find card %v in result but did not", card)
			}
		}
	}
}

// contains checks whether a slice of tarotDeck structs contains a given tarotDeck struct
func contains(cards []tarotDeck, card tarotDeck) bool {
	for _, c := range cards {
		if c == card {
			return true
		}
	}
	return false
}

func TestShuffle(t *testing.T) {
	// create a deck of cards using minorArcana
	deck := minorArcana()

	// call shuffle and save the result in a variable
	result := shuffle(deck)

	// assert that the length of the result is correct
	if len(result) != len(deck) {
		t.Errorf("Expected length of %d but got %d", len(deck), len(result))
	}

	// assert that the elements of the result are the same as the original deck, but in a different order
	if !reflect.DeepEqual(deck, result) {
		t.Errorf("Expected result to contain the same elements as the original deck, but in a different order")
	}
}
