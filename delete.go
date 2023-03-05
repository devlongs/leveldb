package main

import (
	"fmt"
)

// Delete data from the database
func delete() {
	err = db.Delete([]byte("key"), nil)
	if err != nil {
		fmt.Println("Failed to delete from database:", err)
		return
	}
}
