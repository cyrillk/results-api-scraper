package main

import "encoding/json"

// JSONLib json lib
type JSONLib struct {
}

// APIResponse APIResponse
type APIResponse struct {
	TestResults []TestResult `json:"test_results"`
	Next        string       `json:"next"`
}

// TestResult TestResult
type TestResult struct {
	TestName   string `json:"test_name"`
	UserID     string `json:"user_id"`
	StartTime  string `json:"test_start_time"`
	FinishTime string `json:"test_finish_time"`
	Status     string `json:"status"`
	Scores     Scores `json:"scores"`
}

// Scores Scores
type Scores struct {
	Level     string    `json:"level"`
	Score     int       `json:"score"`
	Reading   Reading   `json:"reading"`
	Listening Listening `json:"listening"`
}

type Reading struct {
	Score int `json:"score"`
}

type Listening struct {
	Score int `json:"score"`
}

// ParseResponse unmarshals the json
func (jsonLib JSONLib) ParseResponse(body []byte) APIResponse {
	var m APIResponse
	err := json.Unmarshal(body, &m)

	if err != nil {
		panic(err)
	}

	return m
}
