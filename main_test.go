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
		if card.NameSuit != majorCards[card.Number] {
			// If they do not match, fail the test
			t.Errorf("Unexpected card. Expected %s but got %s", majorCards[card.Number], card.NameSuit)
		}
	}
}

func TestMinorArcana(t *testing.T) {
	// Call minorArcana and save the result in a variable
	result := minorArcana()

	// Assert that the length of the result is correct
	expectedLength := len(minorSuits) * len(minorCards)
	if len(result) != expectedLength {
		t.Errorf("Expected length of %d but got %d", expectedLength, len(result))
	}

	// Assert that the elements of the result have the expected values
	for _, fullSuitName := range minorSuits {
		for _, fullNumberName := range minorCards {
			card := tarotDeck{
				Number:   fullNumberName,
				NameSuit: "of " + fullSuitName,
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
		if c.Number == card.Number && c.NameSuit == card.NameSuit {
			return true
		}
	}
	return false
}

func TestShuffle(t *testing.T) {
	// Create a deck of cards using minorArcana
	deck := minorArcana()

	// Make a copy of the deck to ensure the original deck remains unchanged
	originalDeck := make([]tarotDeck, len(deck))
	copy(originalDeck, deck)

	// Call shuffle and save the result in a variable
	shuffledDeck := shuffle(deck)

	// Assert that the length of the result is correct
	if len(shuffledDeck) != len(originalDeck) {
		t.Errorf("Expected length of %d but got %d", len(originalDeck), len(shuffledDeck))
	}

	// Assert that the elements of the result are the same as the original deck
	if !containsAll(shuffledDeck, originalDeck) {
		t.Errorf("Expected shuffled deck to contain the same elements as the original deck")
	}

	// Assert that the order of the shuffled deck is different from the original deck
	if reflect.DeepEqual(originalDeck, shuffledDeck) {
		t.Errorf("Expected shuffled deck to have a different order than the original deck")
	}
}

// containsAll checks whether a slice contains all elements of another slice
func containsAll(slice, subset []tarotDeck) bool {
	for _, elem := range subset {
		if !contains(slice, elem) {
			return false
		}
	}
	return true
}
