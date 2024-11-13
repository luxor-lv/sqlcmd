package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage: sql \"sql statement\"")
		os.Exit(3)
	}

	config, error := loadConfig()
	if error != nil {
		fmt.Println(error)
		os.Exit(3)
	}

	db, error := NewDatabaseConnection(config.Database.Type, config.Database.ConnectionString, args[0] == "-t")
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}

	data, error := db.query(args[len(args)-1])
	if error != nil {
		fmt.Println(error)
		os.Exit(2)
	}

	display(data)
}
