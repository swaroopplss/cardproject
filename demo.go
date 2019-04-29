package main

func main() {
	cards := newDeck()
	cards.saveToFile("card_list")
}
