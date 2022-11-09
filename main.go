package main

import (
	"DEMO-neo4j/utility"
	"fmt"
	_ "github.com/golang-migrate/migrate/source/file"
	"time"
)

func main() {

	drv := utility.InitDb()

	fmt.Print(drv)
	time.Sleep(5 * time.Minute)

}
