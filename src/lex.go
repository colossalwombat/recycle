package main

import(
	"fmt"
	//"io"
	//"os"
	"io/ioutil"
	"regexp"
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main(){ 

	data, err := ioutil.ReadFile("../return_2.c")
	check(err)

	r_identifier, err  := regexp.Compile("[a-zA-Z]\\w*") 
	r_intliteral, err  := regexp.Compile("\\d+") 
	check(err)

	//list all relevant keywords
	keywords := [2]string{"int", "return"}

	id_results := r_identifier.FindAllString(string(data), -1)

	for result := range id_results {
		if result
	}


}