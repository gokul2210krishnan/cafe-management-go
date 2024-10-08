package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	// "github.com/aws/aws-sdk-go/service/marketplacemetering"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello World..")
	fmt.Println(quote.Go())
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// response, err := marketplacemetering.ResolveCustomer()
	msg, err := greetings.Hello("Rakesh")
	checkErr(err)
	fmt.Println(msg)
	names := []string{"Sample1", "Sample2", "Sample3"}
	msgs, err := greetings.Hellos(names)
	checkErr(err)
	fmt.Println(msgs)
	for _, msg := range msgs {
		fmt.Println(msg)
	}
	msg, err = greetings.Hello("")
	checkErr(err)
	fmt.Println(msg)
	fmt.Println(msg)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
