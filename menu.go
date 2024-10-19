package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func showMenu() {
	data, err := os.ReadFile("logo.txt")
	if err != nil {
		fmt.Println("Error when reading file:", err)
		return
	}
	asciiArt := string(data)
	fmt.Println(asciiArt)

	fmt.Println("\nMenu:")
	fmt.Println("1. Set (Add/Update)")
	fmt.Println("2. Get (Retrieve)")
	fmt.Println("3. Delete Bucket")
	fmt.Println("4. Save (Save bucket to file)")
	fmt.Println("5. Load (Load bucket from file)")
	fmt.Println("6. List Buckets")
	fmt.Println("7. Exit")
}

func RunMenu(db *NoSQLDB) {
	var option int

	for {
		clearTerminal()
		showMenu()
		fmt.Print("Choose an option: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			handleSet(db)
		case 2:
			handleGet(db)
		case 3:
			handleDelete(db)
		case 4:
			handleSave(db)
		case 5:
			handleLoad(db)
		case 6:
			db.ListBuckets()
		case 7:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option.")
		}

		fmt.Println("Press Enter to continue...")
		fmt.Scanln()
		fmt.Scanln()
	}
}

func handleSet(db *NoSQLDB) {
	var bucket, key string
	fmt.Print("Bucket name (or 'cancel' to return to menu): ")
	fmt.Scan(&bucket)
	if strings.ToLower(bucket) == "cancel" {
		return
	}

	fmt.Print("Key: ")
	fmt.Scan(&key)
	fmt.Print("JSON Value (Press Ctrl+D (macOS) or Ctrl+Z (Windows) to terminate the input):\n")

	scanner := bufio.NewScanner(os.Stdin)
	var value []byte
	for scanner.Scan() {
		value = append(value, scanner.Bytes()...)
		value = append(value, '\n')
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error when reading the input:", err)
		return
	}

	value = bytes.TrimSpace(value)
	if len(value) == 0 {
		fmt.Println("Error: The input value cannot be null")
		return
	}

	var parsedValue interface{}
	if err := json.Unmarshal(value, &parsedValue); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	db.Set(bucket, key, parsedValue)
	fmt.Println("Value inserted successfully!")
}

func handleGet(db *NoSQLDB) {
	var bucket string
	var option int
	fmt.Print("Bucket name (or 'cancel' to return to menu): ")
	fmt.Scan(&bucket)
	if strings.ToLower(bucket) == "cancel" {
		return
	}

	fmt.Println("\nGet Options:")
	fmt.Println("1. Get specific key")
	fmt.Println("2. List all keys")
	fmt.Print("Choose an option: ")
	fmt.Scan(&option)

	switch option {
	case 1:
		var key string
		fmt.Print("Key: ")
		fmt.Scan(&key)
		if value, exists := db.Get(bucket, key); exists {
			jsonValue, _ := json.MarshalIndent(value, "", "  ")
			fmt.Println("Value:", string(jsonValue))
		} else {
			fmt.Println("Data not found.")
		}
	case 2:
		db.GetAll(bucket)
	default:
		fmt.Println("Invalid option.")
	}
}

func handleDelete(db *NoSQLDB) {
	var bucket string
	fmt.Print("Bucket name to delete: ")
	_, err := fmt.Scan(&bucket)
	if strings.ToLower(bucket) == "cancel" {
		return
	}
	if err != nil {
		fmt.Println("Error reading bucket name:", err)
		return
	}

	if _, ok := db.buckets[bucket]; ok {
		db.DeleteBucket(bucket)
		fmt.Println("Bucket", bucket, "Removed successfully!")
	} else {
		fmt.Println("Bucket", bucket, "Does not exist.")
	}
}

func handleSave(db *NoSQLDB) {
	var bucket, filename string
	fmt.Print("Bucket name (or 'cancel' to return to menu): ")
	fmt.Scan(&bucket)
	if strings.ToLower(bucket) == "cancel" {
		return
	}
	fmt.Print("File name: ")
	fmt.Scan(&filename)

	db.SaveBucketToFile(bucket, filename)

	fmt.Println("File created successfully!")
}

func handleLoad(db *NoSQLDB) {
	var bucket, filename string
	fmt.Print("Bucket name (or 'cancel' to return to menu): ")
	fmt.Scan(&bucket)
	if strings.ToLower(bucket) == "cancel" {
		return
	}
	fmt.Print("File name:")
	fmt.Scan(&filename)

	db.LoadBucketFromFile(bucket, filename)

	fmt.Println("Bucket loaded successfully!")
}
