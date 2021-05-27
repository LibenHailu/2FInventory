package main

import (
	"log"

	"github.com/LibenHailu/2FInventory/intitiator"
)

func main() {
	// if we crash get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	intitiator.Inventory(true)
}
