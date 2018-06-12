package main

import(
	"fmt"
	//"io"
	//"os"
	"io/ioutil"
	"regexp"
)

//define the types as ints for theoretically better performance
const OPEN_BRACE = 2
const CLOSE_BRACE = 3
const IDENTIFIER = 4
const KEYWORD = 5
const OPEN_PAREN = 6
const CLOSE_PAREN = 7
const SEMICOL = 8
const INT_LIT = 9

//keyword types
const INT_KEY = 2
const RETURN_KEY = 3

type token struct {
	kind int
	key int
}

var SINGLES = [5]string{"{", "}", "(", ")", ";"}
var KEYWORDS = [2]string{"int", "return"}


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main(){ 

	data, err := ioutil.ReadFile("../return_2.c")
	check(err)

	//compile the special regex expressions
	r_identifier, err  := regexp.Compile("[a-zA-Z]\\w*") 
	r_intliteral, err  := regexp.Compile("\\d+") 
	check(err)



}