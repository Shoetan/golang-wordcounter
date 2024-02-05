package main

import (
	"fmt"
	"os"
	"strings"
	
)

//function to return the count of the words in a text file 

func countWord (s string) int {
	return len(strings.Fields(s))
}

func main() {
		//Args hold the command-line arguments, starting with the program name. This check if the command have an argument
	if len(os.Args) < 2 {
		fmt.Println("📁 File path not added")
		os.Exit(1)
	}

	//store file path in a variable. This comes from the command line arguments which is an array 

	filepath := os.Args[1]
	
	fmt.Println(filepath)



	// now we can read data from the file path
	data, err := os.ReadFile( filepath) 

	if err != nil {
		panic(err)
	}


	count := countWord(string(data))

	fmt.Println(count)
}