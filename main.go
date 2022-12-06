package main

import (
	"crypto/rand"
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
)

type tarotDeck struct {
	number   string
	nameSuit string
	reversed string
}

var majorCards = map[string]string{
	"I":     "The Magician",
	"II":    "The Papess",
	"III":   "The Empress",
	"IV":    "The Emperor",
	"V":     "The Heirophant",
	"VI":    "The Lovers",
	"VII":   "The Chariot",
	"VIII":  "Justice",
	"IX":    "The Hermit",
	"X":     "The Wheel Of Fortune",
	"XI":    "Strength",
	"XII":   "The Hanged Man",
	"XIII":  "Death",
	"XIV":   "Temperance",
	"XV":    "The Devil",
	"XVI":   "The Tower",
	"XVII":  "The Star",
	"XVIII": "The Moon",
	"XIX":   "The Sun",
	"XX":    "The Last Judgment",
	"XXI":   "The World",
	"_":     "The Fool",
}

var minorSuits = []string{
	"Cups",
	"Wands",
	"Swords",
	"Pentacles",
}

var minorCards = []string{
	"Ace",
	"Two",
	"Three",
	"Four",
	"Five",
	"Six",
	"Seven",
	"Eight",
	"Nine",
	"Ten",
	"Page",
	"Knight",
	"Queen",
	"King",
}

// generate a slice of structs major arcana tarot cards from static vars
func majorArcana() []tarotDeck {
	majorArcana := []tarotDeck{}

	for key, value := range majorCards {
		majorArcana = append(majorArcana, tarotDeck{
			number:   key,
			nameSuit: value,
		})
	}
	return majorArcana

}

// generate a slice of structs minor arcana tarot cards from static vars
func minorArcana() []tarotDeck {
	minorArcana := []tarotDeck{}

	for _, suit := range minorSuits {
		for _, number := range minorCards {
			minorArcana = append(minorArcana, tarotDeck{
				number:   number,
				nameSuit: "of " + suit,
			})
		}
	}
	return (minorArcana)
}

// User options function will prompt the user with the specified prompt and a list of options
// It will return the selected option as a string
func userOptions(prompt string, options []string) string {
	// Use the promptui Select function to create a prompt with the specified label and options
	promptOptions := promptui.Select{
		Label: prompt,
		Items: options,
	}

	// Run the prompt and capture the selected option and any error
	_, selectedOption, err := promptOptions.Run()

	// If there was an error, print it and return an empty string
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	// Return the selected option as a string
	return selectedOption
}

// shuffle function will shuffle the specified deck using the crypto/rand package
// It will return the shuffled deck
func shuffle(decks []tarotDeck) []tarotDeck {
	// Iterate over the deck
	for i := range decks {
		// Generate a random byte using crypto/rand
		b := make([]byte, 1)
		rand.Read(b)

		// Use the random byte to calculate a random index
		j := int(b[0]) % (i + 1)

		// Swap the cards at the current index and the random index
		decks[i], decks[j] = decks[j], decks[i]
	}

	// Return the shuffled deck
	return decks
}

// ask the user for a number and suggest a default
func askForNumber(prompt string, defaultValue int) int {
	fmt.Printf("%s %d]: ", prompt, defaultValue)

	// Read the user's input
	var input string
	fmt.Scanln(&input)

	// If the input is empty, return the default value
	if input == "" {
		return defaultValue
	}

	// Otherwise, try to convert the input to an integer
	value, err := strconv.Atoi(input)
	if err != nil {
		// If the conversion fails, return the default value
		return defaultValue
	}

	// If the conversion succeeds, return the input value
	return value
}

func main() {
	// Ask the user which deck to use
	deckSize := userOptions("Which cards would you like to use?", []string{"Full Deck", "Major Arcana only", "Minor Arcana only"})

	// Ask the user if they want to include reversed cards
	deckReverse := userOptions("Would you like to include reversed cards?", []string{"Upright and reversed", "Upright only"})

	// Declare a slice of tarotDeck structs to hold the selected deck
	var decks []tarotDeck

	// Based on the user's selection, generate the appropriate deck
	switch {
	case deckSize == "Major Arcana only":
		// If the user selected the major arcana only, generate the major arcana deck
		decks = majorArcana()
	case deckSize == "Minor Arcana only":
		// If the user selected the minor arcana only, generate the minor arcana deck
		decks = minorArcana()
	case deckSize == "Full Deck":
		// If the user selected the full deck, generate the major and minor arcana decks and combine them
		decks = append(majorArcana(), minorArcana()...)
	}

	// If the user selected to include reversed cards, randomly assign upright or reversed to each card in the deck
	if deckReverse == "Upright and reversed" {
		// Create a new deck to hold the upright and reversed cards
		var newDecks []tarotDeck

		// Iterate over the selected deck
		for i := range decks {
			// Generate a random byte using crypto/rand
			b := make([]byte, 1)
			rand.Read(b)

			// If the random byte is even, add an upright version of the card to the new deck
			// Otherwise, add a reversed version of the card
			if b[0]%2 == 0 {
				newDecks = append(newDecks, tarotDeck{
					number:   decks[i].number,
					nameSuit: decks[i].nameSuit,
				})
			} else {
				newDecks = append(newDecks, tarotDeck{
					number:   decks[i].number,
					nameSuit: decks[i].nameSuit,
					reversed: "Reversed",
				})
			}
		}

		// Set the selected deck to the new deck with upright and reversed cards
		decks = newDecks
	}

	// Shuffle the selected deck
	decks = shuffle(decks)

	// Ask the user how many cards they want to draw
	fmt.Printf("There are %d cards in the deck. ", len(decks))
	// Ask the user for a number and provide a default value of 8
	numCards := askForNumber("How many would you like to draw? [Default:", 8)

	// If the number of cards the user wants to draw is greater than the number of cards in the deck,
	// set the number of cards to draw to the number of cards in the deck
	if numCards > len(decks) {
		numCards = len(decks)
		// Print a message to user
		fmt.Printf("There are only %d cards in the deck!\n", len(decks))
	}
	// print a blank line to separate output
	fmt.Println()

	// Draw and print the specified number of cards from the deck
	for i := 0; i < numCards; i++ {
		// Pop the first card from the deck
		card := decks[0]
		decks = decks[1:]

		// Print the card
		fmt.Printf("Card %d: %s %s (%s)\n", i+1, card.number, card.nameSuit, card.reversed)
	}
}
