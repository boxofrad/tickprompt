package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	// Environment variables
	SUBSCRIPTION_ID      = "TICKPROMPT_SUBSCRIPTION_ID"
	API_TOKEN            = "TICKPROMPT_API_TOKEN"
	USER_ID              = "TICKPROMPT_USER_ID"
	EXPECTED_DAILY_HOURS = "TICKPROMPT_EXPECTED_DAILY_HOURS"

	// Defaults
	DEFAULT_EXPECTED_DAILY_HOURS = 7.5
)

type Config struct {
	SubscriptionId     int
	ApiToken           string
	UserId             int
	ExpectedDailyHours float32
}

func loadConfig() (*Config, error) {
	subscriptionId, err := getEnvInt(SUBSCRIPTION_ID)

	if err != nil {
		return nil, err
	}

	apiToken, err := getEnvStr(API_TOKEN)

	if err != nil {
		return nil, err
	}

	userId, err := getEnvInt(USER_ID)

	if err != nil {
		return nil, err
	}

	expectedStr := os.Getenv(EXPECTED_DAILY_HOURS)

	var expectedDailyHours float64

	if expectedStr == "" {
		expectedDailyHours = DEFAULT_EXPECTED_DAILY_HOURS
	} else {
		expectedDailyHours, err = strconv.ParseFloat(expectedStr, 32)

		if err != nil {
			return nil, err
		}
	}

	return &Config{
		SubscriptionId:     subscriptionId,
		ApiToken:           apiToken,
		UserId:             userId,
		ExpectedDailyHours: float32(expectedDailyHours),
	}, nil
}

func getEnvInt(key string) (int, error) {
	val, err := getEnvStr(key)

	if err != nil {
		return -1, err
	}

	return strconv.Atoi(val)
}

func getEnvStr(key string) (string, error) {
	var err error

	val := os.Getenv(key)

	if val == "" {
		err = fmt.Errorf("missing env var `%s`", key)
	}

	return val, err
}
