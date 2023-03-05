package main

import (
	"fmt"
)

// Iterate over the database
func iterate() {
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("Key: %s, Value: %s\n", string(key), string(value))
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		fmt.Println("Failed to iterate over database:", err)
		return
	}
}
