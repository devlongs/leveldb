package main

import (
	"fmt"
)

func Write() {
	for i := 0; i < 5; i++ {
		// Write data to the database
		key := []byte("key" + fmt.Sprint(i))
		value := []byte("my value is" + fmt.Sprint(i))
		err = db.Put([]byte(key), value, nil)
		if err != nil {
			fmt.Println("Failed to write to database:", err)
			return
		}
	}
}
