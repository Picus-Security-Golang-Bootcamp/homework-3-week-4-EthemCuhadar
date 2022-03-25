package csv_utils

import (
	"os"
)

// readCSV reads book information in a CSV file
// and returns a book slice to store the books
// into database.
func readCSV(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}
