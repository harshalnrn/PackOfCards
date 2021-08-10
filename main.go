package main

import "fmt"

// if not return type, then void by default(i.e nothing returned)
func main() {
	//card5 := deck{"deck1", "deck2", "deck3"} //custom type usage inside main method
	//so above card5 is of type 'deck' and hence has access to receiver function of 'deck'

	//card5 = append(card5, "deck4")
	//card5.print()
	cards := newDeck() //we can access any methods part of same package (i.e in this case 'main')
	fmt.Println("items before shuffling")
	cards.print()

	shuffledCards := cards.shufle()
	fmt.Println("items after shuffling")
	shuffledCards.print()
	//hand, remainingCards := deal(cards, 5) // dividing pack into two randomly, into two decks
	//type conversion example in go
	//stringToWrite :=cards.toString())

	//cards.saveToFile("sample")
	//fileContentInString:=string (cards.readFile("sample"))
	//cards := newDeckFromFile("sample")
	//cards.print()

	//fmt.Println([]byte("Hello")) // note: type conversion has rules, and can lead to error on invalid type conversion, if not used appropriatly.
	// hand.print()
	//fmt.Println("next slice")
	//remainingCards.print()

	/* for i, currentEle := range card5 {
		fmt.Println(i, currentEle)
	} */
	/* 	var card string = "pack of cards" //fixed string datatype
	   	//or
	   	card1 := "alternat way of defining variable" //   (var, datatype become optional in this declaration format)
	   	card = "changed from pack of cards to 2 packs of card"
	   	 fmt.Println(card)
	   	fmt.Println(card1)
	   	fmt.Println(newCard())
	   	fmt.Println(newCard1())  */
	//cardArray1()
	//cardArray()

	//	card=1; //so note Go is not dynamic like JS. hence dattype cannot be randomly changed or reassigned. checked out typecasting if its exists.later
}

//note how return type is mentioned.
func newCard() string {
	return "five of spades"
}

//note how return type is mentioned.
func newCard1() int {
	return 5
}

//Arrays datastrucutre

func cardArray() {
	var card = [3]string{"one", "two", "third"} //delcaration +initialization
	card1 := [3]string{"as", "df", "df"}        // doing away with var keyword

	for i, currentEle := range card {
		fmt.Println(i, currentEle)
	}
	for i, currentEle := range card1 {
		fmt.Println(i, currentEle)
	}
}

//arrays
func cardArray1() {
	//to separate declaration and initilization .
	var card [3]int //declarated and auto-initiliased with default 0

	for i, card1 := range card {
		card[i] = i //reintiialise based upon index.
		fmt.Println(card[i], card1)
	}
}

//Slice datastrucue

func cardSlice() {
	card2 := []string{"one", "two", "third"} //variable of slice datatype     // note no size defined
	card2 = append(card2, "fouth")           //dynamic addition //append belongs to which package.
	//itirating the slice
	//i index starting from 0
	//card2 : current element
	//range card2: slice of cards where range used as keyword always
	for i, card2 := range card2 {
		fmt.Println(i, card2)
	}

}
