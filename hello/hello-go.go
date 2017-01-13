
package main

import (
	"github.com/tiwaana/golang/anand"
	"fmt"
	"time"
)

func anand_print() {
	fmt.Printf("Anand says Hello, world ##### local .\n")
	fmt.Println("Anand says Hello world with println ##### local ");
	fmt.Print("Andy says Hello work with print ##### local \n");
}


func main() {

	go anand_print()
	go anand.Anand_print()
	anon_print := func () {
		fmt.Println("Anon function \n")
	}
	anon_print()
	time.Sleep(1000000 * time.Millisecond)
	fmt.Printf("Hello, world.\n")
	fmt.Println("Hello world with println");
	fmt.Print("Hello work with print\n");


}

