package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB
var err error

func main() {
	// Open the database
	db, err = leveldb.OpenFile("database/db", nil)
	if err != nil {
		fmt.Println("Failed to open database:", err)
		return
	}
	defer db.Close()

	Write()
	// get()
	//iterate()
	//delete()

}
