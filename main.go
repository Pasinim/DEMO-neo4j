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
	//:= r.ContainsItem("ciccio", 123)
	fmt.Println(r.ContainsItem("x", 12222223))

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
