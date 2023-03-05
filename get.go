package main

import (
	"fmt"
)

// Read data from the database
func get() {
	data, err := db.Get([]byte("key3"), nil)
	if err != nil {
		fmt.Println("Failed to read from database:", err)
		return
	}
	fmt.Println("Value:", string(data))
}
