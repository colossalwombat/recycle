package main

import(
	"os"
	)

//Special Constants for the tree
const TREE_ROOT = -1
const FUNCTION  = -2

type treeNode struct {
	node token
	children []token
	parent token
}


func parseProgram(list []token)(treeNode){
	if list[0].kind != KEYWORD && list[0].key != INT_KEY {
		fmt.Println("Program does not begin with the return specifier. Aborting")
		os.exit(-1)
	}

	//otherwise, check the function
	tree := treeNode{token{TREE_ROOT}, parseFunction(list, tree), nil}
}

func parseFunction(list []token, tree treeNode)(treeNode){
	//this digusting conditional checks for the standard syntax of a function with no arguments
	if list[1].kind == IDENTIFIER && list[2].kind == OPEN_PAREN && list[3].kind == CLOSE_PAREN && list[4].kind == OPEN_BRACE {
		leaf := treeNode{token{FUNCTION, 0, list[1].id_string, nil}, []token{parseStatement(list, 5, leaf)}, tree}
		return leaf
	}
}

func parseStatment(list []token, position int, tree treeNode){
	cur := list[position]

	for cur.kind != CLOSE_PAREN {
		//for each statment

		var amt_to_advance = 0

		//error if the statement doesn't begin with an identifier
		if cur.kind != IDENTIFIER || cur.kind != KEYWORD {
			fmt.Println("Program does not begin with the return specifier. Aborting")
			os.exit(-1)
		}

		//otherwise statment is probably correct 
		children_to_be, amt_to_advance = parseExpression(list, position)

		position += amt_to_advance
		cur = list[position]
	}
}

func parseExpression(list []token, position int) (treeNode, int){

	cur := list[position]

	exp = treeNode{}

	for cur.kind != SEMICOLON {
		cur_exp := exp

		cur_exp







