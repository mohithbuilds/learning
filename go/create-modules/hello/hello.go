package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// set properties of the predefined Logger:
	// the log entry prefix + a flag to disable printing time, source file, and
	// line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// SUCCESSFUL FLOW
	// Get a greeting message and print it
	message, _ := greetings.Hello("Mohith")
	// had err in place of _
	fmt.Println(message)
	// if err != nil {
	// log.Fatal(err)
	// }

	// FAIL FLOW
	mess, e := greetings.Hello("")
	// if an error is returned we should handle it here
	// we will print and exit the program on error
	if e != nil {
		log.Fatal(e)
	}
	// print below never gets executed since error exists and
	// log.Fatal causes program to exit
	fmt.Println(mess)
}
