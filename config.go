package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	// Environment variables
	SUBSCRIPTION_ID = "TICKPROMPT_SUBSCRIPTION_ID"
	API_TOKEN       = "TICKPROMPT_API_TOKEN"
	USER_ID         = "TICKPROMPT_USER_ID"
	CACHE_TTL       = "TICKPROMPT_CACHE_TTL"

	// Defaults
	DEFAULT_CACHE_TTL = 900
)

type Config struct {
	SubscriptionId int
	ApiToken       string
	UserId         int
	CacheTTL       time.Duration
}

func LoadConfig() (*Config, error) {
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

	cacheTTL, err := getEnvInt(CACHE_TTL)

	if err != nil {
		cacheTTL = DEFAULT_CACHE_TTL
	}

	return &Config{
		SubscriptionId: subscriptionId,
		ApiToken:       apiToken,
		UserId:         userId,
		CacheTTL:       time.Duration(cacheTTL) * time.Second,
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
