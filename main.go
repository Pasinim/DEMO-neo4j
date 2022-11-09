package main

import (
	"DEMO-neo4j/funcs"
	"fmt"
)

func main() {
	r := funcs.New()
	r.GetItems()

	var n int
	fmt.Scan(&n)

}
