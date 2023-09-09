package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)


type createTimeEntryRequest struct {
	// Billable bool `json:"billable"`
	CreatedWith string `json:"created_with"`
	// Description string `json:"description"`
	Duration int `json:"duration"`
	// Duronly bool `json:"duronly"`
	// Pid int `json:"pid"`
	// ProjectID int `json:"project_id"`
	Start string `json:"start"`
	// StartDate string `json:"start_date"`
	// Stop string `json:"stop"`
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

var email = os.Getenv("TOGGL_EMAIL")
var password = os.Getenv("TOGGL_PASSWORD")
var workspaceId, _ = strconv.Atoi(os.Getenv("TOGGL_WORKSPACE_ID"))
var durationMin, _ = strconv.Atoi(os.Getenv("DURATION_MIN"))

func handler(ctx context.Context, event events.CloudWatchEvent) (bool, error) {
	data, err := json.Marshal(createTimeEntryRequest{
		Start: time.Now().UTC().Add(time.Duration(-durationMin) * time.Minute).Format(time.RFC3339),
		Duration: durationMin * 60,
		CreatedWith: "toggl-aws-lambda",
		WorkspaceID: workspaceId,
	})
	
	print(string(data))
	
	if err != nil {
		print(err)
	}
	
	url := fmt.Sprintf("https://api.track.toggl.com/api/v9/workspaces/%d/time_entries", workspaceId)
	
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data));
	if err != nil {
		print(err)
	}
	
	req.Header.Set("Content-Type", "application/json")
	
	req.SetBasicAuth(email, password)
	
	client := &http.Client{}
	resp, err := client.Do(req)
	
	if err != nil {
		panic(err)
	}
	
	defer resp.Body.Close()
		
	if err != nil {
		panic(err)
	}
	
	return true, nil
}

func main() {
	if os.Getenv("ENV") == "development" {
		err := godotenv.Load(".env")
		
		if err != nil {
			panic(err)
		}
	}
	
	lambda.Start(handler)
}