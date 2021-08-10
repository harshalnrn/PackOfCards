package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Create a new custom type that extends an existing type and add a custom method called 'receiver function
//note the naming conventions
type deck []string // this is like saying in oop world, that new custom type 'deck' extends slice of string, where go 'custom type' is like a java class

//type keyword

//identify receiver and non receiver functions below

func newDeck() deck {
	//creates a deck from available pack of cards(ie. slice)

	cards := deck{} // cards of type deck
	//below two variables describe the pack of cards
	cardSuits := []string{"Spades", "Hearts", "Diamond", "clubs"}
	cardValuesOfEachSuit := []string{"Ace", "one", "two"}
	// now for every combination of(suite, value) : we create a new card. (n2 loop)

	//replace variable with _ if do not want to use it explicitly. example: index
	for _, suitEle := range cardSuits {
		for _, valueEle := range cardValuesOfEachSuit {
			cards = append(cards, valueEle+" of "+suitEle) //apending element to slice
		}
	}

	//return a value of type deck
	return cards
}

//so below custom print() is a recevier function: any variable of type deck gets access print() from anywhere at this moment.
//so here d is reference to the variable of type deck that would call this function.
func (d deck) print() {
	for i, currentEle := range d {
		fmt.Println(i, currentEle)
	}
}
func deal(d deck, handSize int) (deck, deck) { //returning multiple values

	//understand splitting slice further in structs concept

	//divide deck into 2 separate slices
	//headsize being the split index
	return d[:handSize], d[handSize:]

}

//converting data of type deck to string
//receiver function
func (d deck) toString() string {

	stringArr := []string(d) //cast to []string  //can be casted if types part of the hierarchy chain
	//lets create a string by appending each character part of array separated by ,
	/* finalString := ""
	for _, currentEle := range stringArr {
		//check if finalString is not empty/blank before doing
		finalString = finalString + "," + currentEle
	} */

	//using strings package.
	return strings.Join(stringArr, ",")

}
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

//below need not be receiver, since we dont really have a deck yet.

func newDeckFromFile(filename string) deck {
	//input file --[]byte ---string ---deck

	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("file could not be read. terminating newDeckCreation")
		os.Exit(1) //terminates complete program.
	}
	//type conversion of []byte to string
	s := strings.Split(string(byteSlice), ",") //[] string
	return deck(s)                             //type conversion to deck

}
func readFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func (d deck) shufle() deck {

	//understand how : the random  generation logic for index,  to shuffle cards
	//                 : each shuffle also to be random
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	//for each index of deck
	// generate random number n between 0 to length-1 by using available golang defined method
	//swap d[i] with d[n]
	for i := range d {
		randomIndex := r.Intn(len(d) - 1)
		temp := d[i]
		d[i] = d[randomIndex]
		d[randomIndex] = temp
	}
	return d
}

//setting up context:
/* The building blocks of Golang are types, functions, and packages—in contrast to classes-objects in object-oriented languages like Java, for modular design patterns.
However, basic concepts of OOP (encapsulation, abstraction, and polymorphism) are also available */
