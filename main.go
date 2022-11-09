package main

import (
	"DEMO-neo4j/funcs"
	"fmt"
	"log"
)

func main() {
	r := funcs.New()
	//r.GetItems()
	//fmt.Println(its)
	it, err := r.GetItemFromSku(11)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(it)
	var n int
	fmt.Scan(&n)

}
