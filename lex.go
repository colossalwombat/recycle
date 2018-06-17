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
	list = append(list, token{})
	copy(list[position+1:], list[position:])
	list[position] = tk

	//fixes a slice issue, will rework later
	/*if len(list) > 2 {
		list = list[:len(list) -1]
	}*/

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


	list = append(list, token{})
	for i := 0; i < len(list); i++ {
		//fmt.Printf("%d < %d, %d > %d\n", tk.index, list[i].index, tk.index, list[i - 1].index)
		if tk.index < list[i].index{
			//insert the token
			return insert(list, tk, i)
		}

	}
	//this should never get reached
	return nil
}

func main(){ 

	data, err := ioutil.ReadFile("/Users/Jack/Desktop/return_2.c")
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
			fmt.Println(token_list)
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
				addToTokenList(token_list, new_token)
				continue RESULT_LOOP
			}
		}
		new_token.kind = IDENTIFIER
		new_token.id_string = string(id_results[i])
		new_token.index = id_results_pos[i][0]
		addToTokenList(token_list, new_token)
	}

	fmt.Println(token_list)

}