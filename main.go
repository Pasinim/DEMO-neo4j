package main

import (
	"DEMO-neo4j/funcs"
	"fmt"
	"log"
)

func main() {

	r := funcs.New()

	b := r.InsertItem("ciccio", 123)
	if !b {
		log.Fatal(b)
	}
	//}
	//	log.Fatal(err)
	//if err != nil {
	//it, err := r.GetItemFromSku(11)

	its := r.GetItems()
	fmt.Println(its)
	fmt.Println("Running... Premi un tasto per terminare\n")
	var n int
	fmt.Scan(&n)

}
