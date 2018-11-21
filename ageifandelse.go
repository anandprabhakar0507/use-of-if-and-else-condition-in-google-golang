package main

import "fmt"

func main() {

	age := 33

	if age < 13 {
		fmt.Println("You are young")
	} else if age < 20 {
		fmt.Println("you're a teenager")
	} else if age < 30 {
		fmt.Println("you are in your twenties")
	} else if age < 40 {
		fmt.Println("you are in you thirties")
	} else if age < 50 {
		fmt.Println("Now you are going older")
	} else if age < 100 {
		fmt.Println("you are old now")
	}
}
