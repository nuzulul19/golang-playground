package main

func main() {
	// card := "Ace of Spades" // same with: var card string = "Ace of Spades"
	// card := newCard()
	// cards := []string{newCard(), "Ace of Diamonds"}
	// cards := newDeck()
	// cards.saveToFile("my_cards")
	cards := newDeckFromFile("my_cards")

	_, remainingDeck := deal(cards, 5)

	// hand.print()
	remainingDeck.shuffle()
	remainingDeck.print()
}
