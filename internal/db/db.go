package db

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// Structure of the NoSQL database
type NoSQLDB struct {
	buckets map[string]map[string]interface{}
	mu      sync.RWMutex
}

// Initializes a new database
func NewNoSQLDB() *NoSQLDB {
	return &NoSQLDB{buckets: make(map[string]map[string]interface{})}
}

// Creates or updates data in a bucket
func (db *NoSQLDB) Set(bucket, key string, value interface{}) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, exists := db.buckets[bucket]; !exists {
		db.buckets[bucket] = make(map[string]interface{})
	}
	db.buckets[bucket][key] = value
}

// Retrieves data from a bucket
func (db *NoSQLDB) Get(bucket, key string) (interface{}, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	if b, exists := db.buckets[bucket]; exists {
		value, exists := b[key]
		return value, exists
	}
	return nil, false
}

// Gets all data from a bucket
func (db *NoSQLDB) GetAll(bucket string) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	if data, ok := db.buckets[bucket]; ok {
		if len(data) == 0 {
			fmt.Println("Bucket is empty.")
			return
		}
		fmt.Println("Data in bucket", bucket+":")
		for key, value := range data {
			jsonValue, _ := json.MarshalIndent(value, "", "  ")
			fmt.Println("Key:", key)
			fmt.Println("Value:", string(jsonValue))
			fmt.Println("---")
		}
	} else {
		fmt.Println("Bucket", bucket, "not found.")
	}
}

// BucketExists checks if a bucket exists.
func (db *NoSQLDB) BucketExists(bucket string) bool {
	_, ok := db.buckets[bucket]
	return ok
}

// Removes the selected bucket
func (db *NoSQLDB) DeleteBucket(bucket string) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	delete(db.buckets, bucket)
}

// Creates the "data" folder if it doesn't exist
func ensureDataDir() error {
	path := "data"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModePerm)
	}
	return nil
}

func (db *NoSQLDB) SaveBucketToFile(bucket, filename string) (int, error) {
	if err := ensureDataDir(); err != nil {
		return 0, err
	}

	db.mu.RLock()
	defer db.mu.RUnlock()

	data, exists := db.buckets[bucket]
	if !exists {
		return 0, fmt.Errorf("bucket '%s' not found", bucket)
	}

	filePath := filepath.Join("data", filename)
	file, err := os.Create(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return 0, err
	}

	n, err := file.Write(jsonData)
	return n, err
}

// Loads a bucket from a file in the "data" folder
func (db *NoSQLDB) LoadBucketFromFile(bucket, filename string) error {
	filePath := filepath.Join("data", filename)

	file, err := os.Open(filePath)

	if err != nil {
		return fmt.Errorf("error opening file: %w", err) // More informative error
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("error getting file info: %w", err)
	}
	if fileInfo.Size() == 0 {
		return fmt.Errorf("file is empty: %s", filePath)
	}

	var data map[string]interface{}
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return fmt.Errorf("error decoding JSON: %s", err)
	}

	db.mu.Lock()
	defer db.mu.Unlock()
	db.buckets[bucket] = data
	return nil
}

// Lists all buckets with approximate memory usage
func (db *NoSQLDB) ListBuckets() {
	db.mu.RLock()
	defer db.mu.RUnlock()

	if len(db.buckets) == 0 {
		fmt.Println("No buckets available.")
		return
	}
	fmt.Println("Available buckets:")
	for bucket, data := range db.buckets {
		var size int64
		jsonData, _ := json.Marshal(data)
		size = int64(len(jsonData))
		fmt.Printf("- %s (approx. %.2f KB)\n", bucket, float64(size)/1024)

	}
}
