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
const OPEN_PAREN = 4
const CLOSE_PAREN = 5
const SEMICOL = 6
const IDENTIFIER = 7 
const KEYWORD = 8
const INT_LIT = 9

//keyword types
const INT_KEY = 2
const RETURN_KEY = 3


type token struct {
	kind int
	key int
	id_string string
	index int
}

var SINGLES = [5]string{"{", "}", "\\(", "\\)", ";"}
var KEYWORDS = [2]string{"int", "return"}


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func insert(list []token, tk token, position int) ([]token){
	//does the actual insertion

	//check if the underlying capacity is insufficient
	if cap(list) <= len(list) + 1 {
		new_list := make([]token, len(list) + 1, len(list) + 1)
		copy(new_list, list)
		list = new_list
	}


	copy(list[position+1:], list[position:len(list) - 1])
	list[position] = tk


	return list

}

func addToTokenList(list []token, tk token)([]token){
	//catch two edge cases where the list is empty or contains one element
	if len(list) == 0 {
		return append(list, tk)
	}
	if len(list) < 2 {
		if tk.index < list[0].index {
			return insert(list, tk, 0)
		} else {
			//the index is larger
			return insert(list, tk, 1)
		}
	}

	for i := 0; i < len(list); i++ {
		if tk.index < list[i].index{
			//insert the token
			return insert(list, tk, i)
		}

	}
	//this should never get reached
	return nil
}

func parseFile(filename string) ([]token){ 

	data, err := ioutil.ReadFile(filename)
	check(err)

	token_list := []token{}

	//check for the singles
	for i := 0; i < len(SINGLES); i++ {
		exp, err := regexp.Compile(SINGLES[i])
		check(err)

		results := exp.FindAllString(string(data), -1)
		positions := exp.FindAllStringIndex(string(data), -1)

		for j := 0; j < len(results); j++ {
			new_token := token{i+2, 0, "", positions[j][0]}
			token_list = addToTokenList(token_list, new_token)
		}
	}

	//compile the special regex expressions
	r_identifier, err  := regexp.Compile("[a-zA-Z]\\w*") 
	//r_intliteral, err  := regexp.Compile("\\d+") 
	check(err)

	//handle the identifiers
	id_results := r_identifier.FindAllString(string(data), -1)
	id_results_pos := r_identifier.FindAllStringIndex(string(data), -1)


	RESULT_LOOP:
	for i := 0; i < len(id_results); i++ {
		new_token := token{}
		for j := 0; j < len(KEYWORDS); j++ {
			if id_results[i] == KEYWORDS[j] {
				new_token.kind = KEYWORD
				new_token.key = j+2
				new_token.index = id_results_pos[i][0]
				token_list = addToTokenList(token_list, new_token)
				continue RESULT_LOOP
			}
		}
		new_token.kind = IDENTIFIER
		new_token.id_string = string(id_results[i])
		new_token.index = id_results_pos[i][0]
		token_list = addToTokenList(token_list, new_token)
	}

	fmt.Println(token_list)
	return token_list

}