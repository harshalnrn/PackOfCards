package main

//import required golang packages from installed standard library
import (
	"os"
	"testing"
)

//A module is a collection of "external dependency Go packages" stored in a file tree with a 'go. mod' file at its root
//track all required module dependencies(using go mod),   tools/pluggin for golang : important. note what all!
//question: but will this work only for IDE or also outside IDE in environment??

//testcases:
//case 1: validate no of cards deck : 12
//case2: validate the series (First caed, middle card, last card)

//note method names convention here
// what kind of datatype is *testing.T ?
//look at visual studio clickable features on test methods
/*
func TestNewDeck(t *testing.T) {

	//Act
	d := newDeck()
	//Assert
	//note the control structures in golang
	//check out placeholder way for dynamic values in string. (like java: string.format() %s), instead of concatenation.
	//1
	if len(d) != 12 {
		//test fails
		t.Errorf("expected length of deck as 12, but instead got %v", len(d))
	}
	//2
	if d[0] != "Ace of Spades" { // golang nothing like object oriented..where == , .equals varies
		t.Errorf("expected first element to be Ace of Spades, but instead got %v", d[0])
	}

	//3
	if d[len(d)] != "two of clubs" {
		t.Errorf("expected last element to be two of clubs, but instead got %v", d[len(d)])
	}

	//Test case PASS if all 3 assert conditions pass.

} */

//note take of any cleanup if required, if resource introduced while testing(i.e : example: file)

func TestSaveToFileReadFromFileDeleteFile(t *testing.T) {
	//clean up
	os.Remove("sample_test")
	//CREATE DEC
	d := newDeck()
	//SAVE to file
	d.saveToFile("sample_test")
	//READ file and load deck
	readDeck := newDeckFromFile("sample_test")
	length := len(readDeck)
	//Assert no of cards after reading
	if length != 12 {
		t.Errorf("no of cards, read from file is not 12, but instead %v", length)
	}
	//clean up
	os.Remove("sample_test")

}
