package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Open("../iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)

	reader.FieldsPerRecord = -1

	// Read in all of the CSV records.
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rawCSVData)
}
