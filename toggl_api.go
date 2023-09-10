package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type createTimeEntryRequest struct {
	// Billable bool `json:"billable,omitempty"`
	CreatedWith string `json:"created_with,omitempty"`
	Description string `json:"description,omitempty"`
	Duration    int    `json:"duration"`
	// Duronly bool `json:"duronly"`
	// Pid int `json:"pid"`
	// ProjectID int `json:"project_id"`
	Start string `json:"start"`
	// StartDate string `json:"start_date"`
	Stop string `json:"stop,omitempty"`
	// TagAction string `json:"tag_action"`
	// TagIDs []int `json:"tag_ids"`
	// Tags []string `json:"tags"`
	// TaskID int `json:"task_id"`
	// TID int `json:"tid"`
	// UID int `json:"uid"`
	// UserID int `json:"user_id"`
	// WID int `json:"wid"`
	WorkspaceID int `json:"workspace_id"`
}

type getTimeEntriesRequest struct {
	// Since int `json:"since" url:"since"`
	// Before string `json:"before" url:"before"`
	StartDate string `json:"start_date" url:"start_date"`
	EndDate   string `json:"end_date" url:"end_date"`
}

type timeEntry struct {
	At              string   `json:"at"`
	Billable        bool     `json:"billable"`
	Description     string   `json:"description"`
	Duration        int      `json:"duration"`
	Duronly         bool     `json:"duronly"`
	ID              int      `json:"id"`
	PID             int      `json:"pid"`
	ProjectID       int      `json:"project_id"`
	ServerDeletedAt string   `json:"server_deleted_at"`
	Start           string   `json:"start"`
	Stop            string   `json:"stop"`
	TagIDs          []int    `json:"tag_ids"`
	Tags            []string `json:"tags"`
	TaskID          int      `json:"task_id"`
	TID             int      `json:"tid"`
	UID             int      `json:"uid"`
	UserID          int      `json:"user_id"`
	WID             int      `json:"wid"`
	WorkspaceID     int      `json:"workspace_id"`
}

func createTimeEntry(entry timeEntry) error {
	log.Print("Creating time entry", entry)

	url := fmt.Sprintf("https://api.track.toggl.com/api/v9/workspaces/%d/time_entries", workspaceId)
	reqBody := &createTimeEntryRequest{
		Start:       entry.Start,
		Stop:        entry.Stop,
		Duration:    entry.Duration,
		Description: entry.Description,
		CreatedWith: "toggl-aws-lambda",
		WorkspaceID: entry.WorkspaceID,
	}

	return request(url, http.MethodPost, reqBody, nil)
}

func deleteTimeEntry(entry timeEntry) error {
	log.Print("Deleting time entry", entry)
	url := fmt.Sprintf("https://api.track.toggl.com/api/v9/workspaces/%d/time_entries/%d", entry.WorkspaceID, entry.ID)
	return request(url, http.MethodDelete, nil, nil)
}

func getTimeEntriesInPast24Hours() (*[]timeEntry, error) {
	log.Print("Getting time entries in past 24 hours")
	url := "https://api.track.toggl.com/api/v9/me/time_entries"
	reqBody := &getTimeEntriesRequest{
		StartDate: time.Now().UTC().Add(time.Duration(-24) * time.Hour).Format(time.RFC3339),
		EndDate:   time.Now().UTC().Format(time.RFC3339),
	}
	timeEntries := &[]timeEntry{}

	err := request(url, http.MethodGet, reqBody, timeEntries)

	if err != nil {
		return nil, err
	}

	return timeEntries, nil
}
