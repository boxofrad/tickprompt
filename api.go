package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Tick API requires a user agent with an email address
const USER_AGENT = "Tickprompt (daniel+open-source@floppy.co)"

type Entry struct {
	Hours float32 `json:"hours"`
}

func GetEntries(config *Config) ([]Entry, error) {
	today := time.Now()
	tommorow := today.Add(24 * time.Hour)

	url := fmt.Sprintf(
		"https://www.tickspot.com/%d/api/v2/users/%d/entries.json?start_date=%s&end_date=%s",
		config.SubscriptionId,
		config.UserId,
		formatDate(today),
		formatDate(tommorow),
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	authHeader := fmt.Sprintf("Token token=%s", config.ApiToken)
	req.Header.Set("Authorization", authHeader)
	req.Header.Set("User-Agent", USER_AGENT)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	buf := bytes.Buffer{}
	buf.ReadFrom(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected response from Tick API (%d) %s", resp.StatusCode, buf.String())
	}

	entries := []Entry{}
	json.Unmarshal(buf.Bytes(), &entries)

	return entries, nil
}

func formatDate(t time.Time) string {
	return t.Format("2006-01-02")
}
