package main

import (
	"DEMO-neo4j/funcs"
	"fmt"
	"log"
	"time"
)

func main() {

	start := time.Now()

	r := funcs.New()
	its := r.GetItems()
	fmt.Println(its)

	it, err := r.GetItemFromSku(11)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(it)
	fmt.Println("Running... Premi un tasto per terminare\n")
	//var n int
	//fmt.Scan(&n)
	duration := time.Since(start)
	fmt.Println(duration.Seconds())

}
