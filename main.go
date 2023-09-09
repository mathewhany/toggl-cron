package main

import (
	"context"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

var email string
var password string
var workspaceId int
var durationMin int

func createEmptyTimeEntry() (string, error) {
	return createTimeEntry(timeEntry {
		Start: time.Now().UTC().Format(time.RFC3339),
		Duration: durationMin * 60,
		WorkspaceID: workspaceId,
	})
}

func breakTimeEntry(entry timeEntry) ([]timeEntry, error) {
	entries := make([]timeEntry, 0)

	for _, description := range strings.Split(entry.Description, "&") {
		newEntry := entry
		newEntry.Description = strings.Trim(description, " ")
		entries = append(entries, newEntry)
	}
	
	return entries, nil
}

func handler(ctx context.Context, event interface {}) (string, error) {	
	_, err := createEmptyTimeEntry()
	
	if err != nil {
		return "", err
	}
	
	timeEntries, err := getTimeEntriesInPast24Hours()
	
	if err != nil {
		return "", err
	}
	
	for _, timeEntry := range timeEntries {
		print(timeEntry.Description + "\n")
		deleteTimeEntry(timeEntry)
		entryParts, err := breakTimeEntry(timeEntry)
		if err != nil {
			return "", err
		}
		
		for _, entryPart := range entryParts {
			_, err := createTimeEntry(entryPart)
			
			if err != nil {
				return "", err
			}
		}	
	}
	
	return "", nil
}

func main() {
	if os.Getenv("ENV") == "development" {
		err := godotenv.Load(".env")
		
		if err != nil {
			panic(err)
		}
	}
	
	email = os.Getenv("TOGGL_EMAIL")
	password = os.Getenv("TOGGL_PASSWORD")
	workspaceId, _ = strconv.Atoi(os.Getenv("TOGGL_WORKSPACE_ID"))
	durationMin, _ = strconv.Atoi(os.Getenv("DURATION_MIN"))
 
	lambda.Start(handler)
}