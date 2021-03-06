package main

import (
	"encoding/csv"
	"log"
	"os"
)

// ResultsWriter results writer
type ResultsWriter interface {
	Write(data []string)
}

// CSVResultsWriter results writer to CSV
type CSVResultsWriter struct {
	ResultsWriter
}

var writer = csv.NewWriter(os.Stdout)

var hasHeaders = false

var headers = []string{
	"user_id",
	"test_name",
	"test_start_time",
	"test_finish_time",
	"status",
	"level",
	"score",
	"reading.score",
	"listening.score",
}

// Write writes
func (resultsWriter CSVResultsWriter) Write(items []string) {

	if hasHeaders == false {
		writeItems(headers)
		hasHeaders = true
	}

	writeItems(items)
}

func writeItems(items []string) {

	writer.Write(items)
	writer.Flush()

	if err := writer.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}
