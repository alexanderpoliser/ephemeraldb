package menu

import (
	"bufio"
	"bytes"
	"encoding/json"
	"ephemeraldb/internal/db"
	"ephemeraldb/internal/utils"
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

func RunMenu(database *db.NoSQLDB) {
	var option int

	for {
		utils.ClearTerminal()
		showMenu()
		fmt.Print("Choose an option: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			handleSet(database)
		case 2:
			handleGet(database)
		case 3:
			handleDelete(database)
		case 4:
			handleSave(database)
		case 5:
			handleLoad(database)
		case 6:
			database.ListBuckets()
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

func handleSet(database *db.NoSQLDB) {
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

	database.Set(bucket, key, parsedValue)
	fmt.Println("Value inserted successfully!")
}

func handleGet(database *db.NoSQLDB) {
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
		if value, exists := database.Get(bucket, key); exists {
			jsonValue, _ := json.MarshalIndent(value, "", "  ")
			fmt.Println("Value:", string(jsonValue))
		} else {
			fmt.Println("Data not found.")
		}
	case 2:
		database.GetAll(bucket)
	default:
		fmt.Println("Invalid option.")
	}
}

func handleDelete(database *db.NoSQLDB) {
	var bucket string
	fmt.Print("Bucket name to delete (or 'cancel' to return to menu): ")
	fmt.Scan(&bucket) // Removed the _, err assignment, using fmt.Scan directly

	if strings.ToLower(bucket) == "cancel" {
		return
	}

	if database.BucketExists(bucket) {
		database.DeleteBucket(bucket)
		fmt.Println("Bucket", bucket, "removed successfully!")
	} else {
		fmt.Println("Bucket", bucket, "does not exist.")
	}
}

func handleSave(database *db.NoSQLDB) {
	var bucket, filename string
	fmt.Print("Bucket name (or 'cancel' to return to menu): ")
	fmt.Scan(&bucket)
	if strings.ToLower(bucket) == "cancel" {
		return
	}
	fmt.Print("File name: ")
	fmt.Scan(&filename)

	database.SaveBucketToFile(bucket, filename)

	fmt.Println("File created successfully!")
}

func handleLoad(database *db.NoSQLDB) {
	var bucket, filename string
	fmt.Print("Bucket name (or 'cancel' to return to menu): ")
	fmt.Scan(&bucket)
	if strings.ToLower(bucket) == "cancel" {
		return
	}
	fmt.Print("File name:")
	fmt.Scan(&filename)

	database.LoadBucketFromFile(bucket, filename)

	fmt.Println("Bucket loaded successfully!")
}
